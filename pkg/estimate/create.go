package estimate

import (
	"context"
	"roofix/ent"
	"roofix/ent/schema"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/homeOwner"
	"roofix/pkg/msg"
	pkgPartner "roofix/pkg/partner"
	"roofix/pkg/postal"
	price "roofix/pkg/pricing"
	"roofix/pkg/util/log"
)

type Input struct {
	ID                *string
	ExternalCompanyID *string
	HomeOwner         *homeOwner.Input
	SalesRep          *SalesRep
	MaterialDetail    *MaterialDetail
	MeasurementType   enum.Measure
	ExtraCharge       *price.ExtraCharge
	PartnerID         *string
	CtxUserID         *string
	CtxApiUserID      *string
}

type Estimation struct {
	Estimator   string
	SQ          float64
	Pitch       float64
	Price       float64
	Bounds      []schema.Point
	Summary     string
	RawResponse map[string]interface{}
}

type SalesRep struct {
	ID          *string `json:"id"` // salesRepID
	CompanyName string  `json:"companyName" validate:"required"`
	FirstName   string  `json:"firstName" validate:"required"`
	LastName    string  `json:"lastName" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Mobile      string  `json:"mobile"`
}

// create estimate request with price estimate details
func create(ctx context.Context, db *ent.Client, inp *Input, homeOwnerHashCheck bool) (id string, err error) {
	// validate input
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	// get sales rep user id
	salesRepID := SalesRepID(ctx, inp.SalesRep.Email)
	if salesRepID == nil && inp.SalesRep.Email != "" && inp.PartnerID != nil {
		// * create sales rep user and link it to partner
		if uID := pkgPartner.QuickAddSalesRepTX(
			ctx, tx, *inp.PartnerID, inp.SalesRep.Email, inp.SalesRep.FirstName, inp.SalesRep.LastName, inp.SalesRep.Mobile,
		); uID != "" {
			salesRepID = &uID
		}
	}

	// save HomeOwner
	var hoID string
	if hoID, err = homeOwner.CreateTx(ctx, tx, inp.PartnerID, inp.HomeOwner, homeOwnerHashCheck); err != nil {
		return "", err
	}

	md := inp.MaterialDetail
	qry := tx.Estimate.Create()
	if inp.ID != nil {
		qry.SetID(*inp.ID)
	}

	qry.
		SetStatus(enum.EstimateStatusNew).
		SetRegionID(postal.GetStateRegion(inp.HomeOwner.State)).
		SetCurrentMaterial(md.CurrentMaterial.String()).
		SetNewRoofingMaterial(md.NewRoofingMaterial.String()).
		SetLowSlope(md.LowSlope).
		SetCurrentMaterialLowSlope(md.CurrentMaterialLowSlope.String()).
		SetNewRoofingMaterialLowSlope(md.NewRoofingMaterialLowSlope.String()).
		SetRedeck(md.Redeck).
		SetLayers(uint8(md.Layers)).
		SetLayer2Material(md.Layer2Material.String()).
		SetLayer3Material(md.Layer3Material.String()).
		SetMeasureType(inp.MeasurementType).
		SetPartialPercentage(md.PartialPercent).
		SetNillableMaterialMappingNote(inp.MaterialDetail.OverrideSummary).
		SetCompanyRefName(inp.SalesRep.CompanyName).
		SetNillableCompanyRefID(inp.ExternalCompanyID).
		SetNillablePartnerID(inp.PartnerID).
		SetNillableSalesRepID(salesRepID).
		SetHomeOwnerID(hoID).
		SetNillableCreatorID(inp.CtxUserID).
		SetNillableCreatorAPIID(inp.CtxApiUserID)

	if inp.ExtraCharge != nil {
		qry.SetExtraChargeType(inp.ExtraCharge.Type).
			SetExtraCharges(inp.ExtraCharge.Val).
			SetExtraChargeNote(inp.ExtraCharge.Note)
		// TODO: save condition & its val
	}

	// save, job
	j, err := qry.Save(ctx)
	if err != nil {
		log.Error(err)
		return "", msg.AsError(msg.FailedToSave, "Estimate")
	}

	return j.ID, tx.Commit()
}

// SalesRepID returns sales rep user id by email
func SalesRepID(ctx context.Context, email string) *string {
	if email != "" {
		if u, err := ent.GetClient().User.Query().Where(user.Email(email)).Select(user.FieldID).Only(ctx); err != nil {
			log.Warn("failed to find SalesRep by email: %s, error: %s", email, err.Error())
		} else {
			return &u.ID
		}
	}

	return nil
}
