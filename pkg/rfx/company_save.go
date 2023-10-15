package rfx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"roofix/ent"
	entPartner "roofix/ent/partner"
	"roofix/ent/partnercontact"
	entUser "roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/uid"
	"roofix/pkg/util/validate"
	"roofix/server/router/response"
)

type CompanyInput struct {
	ExternalID           string   `json:"id" validate:"required"`
	Name                 *string  `json:"name"`
	Website              *string  `json:"website"`
	Phone                *string  `json:"phone"`
	Address              *string  `json:"address"`
	Type                 *uint8   `json:"type"`
	Status               *uint8   `json:"status"`
	PrimaryContactID     *string  `json:"primaryContactID"`
	AccountingContactID  *string  `json:"accountingContactID"`
	ApprovalContactIDs   []string `json:"approvalContactIDs"`
	CustServiceContactID *string  `json:"custServiceContactID"`
	OpsContactID         *string  `json:"operationsContactID"`
	InvoicingContactID   *string  `json:"invoicingContactID"`
	EPCStatus            *uint8   `json:"epcStatus"`
	EPCs                 []string `json:"epcs"`
}

type SaveResponse struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	FailureStatus string `json:"failureStatus"`
}

// SaveCompanyHandler save company data
func SaveCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var id string
	ctx := r.Context()
	action := audit.PartnerSave
	apiUID := account.CtxApiUserID(ctx)
	failed := func(err error) {
		if id != "" {
			audit.ApiOpFailed(ctx, action, apiUID, fmt.Errorf("save company with refID: %s failed with error, %v", id, err))
		} else {
			audit.ApiOpFailed(ctx, action, apiUID, err)
		}
		// return back with 200 with failureStatus
		response.Ok(w, SaveResponse{
			FailureStatus: err.Error(),
		})
	}

	// bind modal
	var inp CompanyInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		audit.ApiOpFailed(ctx, action, apiUID, err)
		failed(err)
		return
	}

	if err := validate.Struct(inp); err != nil {
		failed(err)
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	var exist bool
	// check if partner exists with given externalID
	// get partner by externalID
	if p, err := ent.GetClient().Partner.Query().Where(entPartner.ExternalID(inp.ExternalID)).Select(entPartner.FieldID).First(ctx); err != nil {
		if !ent.IsNotFound(err) {
			failed(log.Wrap(err, "unexpected error"))
			return
		}
	} else {
		exist = true
		id = p.ID
	}

	var message string
	if exist { // update partner
		if saveErr := updatePartner(ctx, db, id, &inp); saveErr != nil {
			failed(saveErr)
			return
		}
		message = "successfully updated company"
		logPartnerActivity(ctx, db, "updated", id, apiUID)
	} else { // create partner
		id = uid.ULID()
		if saveErr := createPartner(ctx, db, id, &inp); saveErr != nil {
			failed(saveErr)
			return
		}
		message = "successfully created company"
		logPartnerActivity(ctx, db, "created", id, apiUID)
	}

	response.Ok(w, SaveResponse{
		Message: message,
		Success: true,
	})
}

// createPartner will have 2 scenarios
// 1. partner without primary user (creat company from admin panel)
// 2. partner with primary user (signup as company on roofix bubble app)
func createPartner(ctx context.Context, db *ent.Client, id string, inp *CompanyInput) error {
	// name is required for both scenarios
	if inp.Name == nil {
		return msg.AsError(msg.ParamMissing, "name")
	}
	// type is required for both scenarios
	if inp.Type == nil {
		return msg.AsError(msg.ParamMissing, "type")
	}

	// status is required for both scenarios
	if inp.Status == nil {
		return msg.AsError(msg.ParamMissing, "status")
	}

	// let's create company first
	statusEPC := epcStatus(num.ReadUInt8Ptr(inp.EPCStatus))
	err := db.Partner.Create().
		SetID(id).
		SetExternalID(inp.ExternalID).
		SetName(*inp.Name).
		SetNillableWebsite(inp.Website).
		SetNillablePhone(inp.Phone).
		SetNillableAddress(inp.Address).
		SetType(companyType(num.ReadUInt8Ptr(inp.Type), statusEPC)).
		SetNillableStatus(companyStatus(inp.Status)).
		SetEpcStatus(statusEPC).
		Exec(ctx)
	if err != nil {
		return log.Wrap(err, "create company")
	}

	return savePartnerContacts(ctx, db, id, inp)
}

// updatePartner information, please note that:
// - type of partner cannot be changed, that can cause lot of issues
func updatePartner(ctx context.Context, db *ent.Client, id string, inp *CompanyInput) error {
	qry := db.Partner.UpdateOneID(id).
		SetNillableWebsite(inp.Website).
		SetNillablePhone(inp.Phone).
		SetNillableAddress(inp.Address).
		SetNillableStatus(companyStatus(inp.Status))

	if inp.Name != nil {
		qry.SetName(*inp.Name)
	}

	if inp.EPCStatus != nil {
		qry.SetEpcStatus(epcStatus(num.ReadUInt8Ptr(inp.EPCStatus)))
	}

	if err := qry.Exec(ctx); err != nil {
		return log.Wrap(err, "update company")
	}

	return savePartnerContacts(ctx, db, id, inp)
}

func savePartnerContacts(ctx context.Context, db *ent.Client, pid string, inp *CompanyInput) error {
	// save primary contact
	if inp.PrimaryContactID != nil {
		if e := saveContact(ctx, db, pid, *inp.PrimaryContactID, enum.PartnerContactPrimary, enum.PartnerContactRoleAdmin); e != nil {
			return log.Wrap(e, "save primary contact")
		}
	}

	// save accounts contact
	if inp.AccountingContactID != nil {
		if e := saveContact(ctx, db, pid, *inp.AccountingContactID, enum.PartnerContactAccounting, enum.PartnerContactRoleAdmin); e != nil {
			return log.Wrap(e, "save accounting contact")
		}
	}

	// approval contacts
	if len(inp.ApprovalContactIDs) > 0 {
		for _, c := range inp.ApprovalContactIDs {
			if e := saveContact(ctx, db, pid, c, enum.PartnerContactApproval, enum.PartnerContactRoleAdmin); e != nil {
				return log.Wrap(e, "save approval contact: "+c)
			}
		}
	}

	// save customer service contact
	if inp.CustServiceContactID != nil {
		if e := saveContact(ctx, db, pid, *inp.CustServiceContactID, enum.PartnerContactCustomerService, enum.PartnerContactRoleAdmin); e != nil {
			return log.Wrap(e, "save customer service contact")
		}
	}

	if inp.InvoicingContactID != nil {
		if e := saveContact(ctx, db, pid, *inp.InvoicingContactID, enum.PartnerContactInvoicing, enum.PartnerContactRoleAdmin); e != nil {
			return log.Wrap(e, "save invoicing contact")
		}
	}

	if inp.OpsContactID != nil {
		if e := saveContact(ctx, db, pid, *inp.OpsContactID, enum.PartnerContactOperations, enum.PartnerContactRoleAdmin); e != nil {
			return log.Wrap(e, "save operations contact")
		}
	}

	return nil
}

// saveContact will save partner:contact relationship. please note that:
// - it will fail if user is already link to another partner
func saveContact(ctx context.Context, db *ent.Client, pID, externalUID string, t enum.PartnerContact, r enum.PartnerContactRole) error {
	var uID string
	// check user exists with given external id
	if u, err := db.User.Query().Where(entUser.ExternalID(externalUID)).Select(entUser.FieldID).First(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fmt.Errorf("failed to find primaryContact by its id")
		} else {
			return log.Wrap(err, "unexpected")
		}
	} else {
		uID = u.ID
	}

	// check if user is already linked to another partner
	chkQry := db.PartnerContact.Query().Where(partnercontact.PartnerIDNEQ(pID), partnercontact.UserID(uID))
	if exist, err := chkQry.Exist(ctx); err != nil {
		return log.Wrap(err, "unexpected")
	} else if exist {
		return fmt.Errorf("userID: %s is already linked to another company", uID)
	}

	// check partner:contact relationship exists
	pcQry := db.PartnerContact.Query().Where(partnercontact.PartnerID(pID), partnercontact.UserID(uID))
	if exist, err2 := pcQry.Exist(ctx); err2 != nil {
		return log.Wrap(err2, "unexpected")
	} else if exist {
		// update partner:contact relationship
		return db.PartnerContact.Update().Where(partnercontact.PartnerID(pID), partnercontact.UserID(uID)).
			SetType(t).
			SetRole(r).
			Exec(ctx)
	} else {
		// create partner:contact relationship
		return db.PartnerContact.Create().
			SetPartnerID(pID).
			SetUserID(uID).
			SetType(t).
			SetRole(r).
			Exec(ctx)
	}
}

// log partner changes activity
func logPartnerActivity(ctx context.Context, db *ent.Client, action, pID string, apiUID string) {
	creator := account.CtxApiUsername(ctx)
	desc := fmt.Sprintf("%s by APIUser: %s", action, creator)
	// log activity
	if e := db.PartnerActivity.Create().SetPartnerID(pID).SetDescription(desc).SetCreatorAPIID(apiUID).Exec(ctx); e != nil {
		log.Error(e)
	}
}

func companyType(v uint8, status enum.EPCStatus) enum.Partner {
	t := enum.PartnerNone
	switch v {
	case 1: // Solar
		t = enum.PartnerSolar
		if status == enum.EPCStatusMultipleDealers || status == enum.EPCStatusVerticallyIntegrated {
			t = enum.PartnerEpc
		}
	case 2: // Integration Partner
		t = enum.PartnerIntegration
	case 3: // Contractor
		t = enum.PartnerRoofing
	case 4: // RFX
		t = enum.PartnerRFX
	}

	return t
}

func epcStatus(v uint8) enum.EPCStatus {
	switch v {
	case 1:
		return enum.EPCStatusDealer
	case 2:
		return enum.EPCStatusMultipleDealers
	case 3:
		return enum.EPCStatusVerticallyIntegrated
	default:
		return enum.EPCStatusNone
	}
}

func companyStatus(v *uint8) *enum.PartnerStatus {
	if v == nil {
		return nil
	}

	var s enum.PartnerStatus
	switch *v {
	case 1: // Active
		s = enum.PartnerStatusActive
	case 2: // Deactivated
		s = enum.PartnerStatusInActive
	}

	return &s
}
