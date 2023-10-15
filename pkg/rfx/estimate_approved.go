package rfx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"roofix/ent"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/estimate"
	"roofix/pkg/nearmap"
	"roofix/pkg/util/log"
	"roofix/pkg/util/validate"
	"roofix/server/model"
	"roofix/server/router/response"
	"time"
)

type EstimateApprovedInput struct {
	ExternalID             string   `json:"externalID" validate:"required"`
	ApprovalStatusID       *uint8   `json:"approvalStatusID"`       // Approval Status (Job) EPC
	ChangeOrderPrice       *float64 `json:"changeOrderPrice"`       // Change Order Price
	ContractPrice          *float64 `json:"contractPrice"`          // Contract Price
	ContractorID           *string  `json:"contractorID"`           // Contractor
	EPCCompanyID           *string  `json:"epcCompanyID"`           // EPC Company
	FinalInspectionDate    *int64   `json:"finalInspectionDate"`    // Final Inspection Date Unix datetime
	FinalInspectionDoc     *string  `json:"finalInspectionDoc"`     // Final Inspection Doc Unrestricted URL
	InstallDate            *int64   `json:"installDate"`            // Install Date Unix datetime
	IntegrationPartnerID   *string  `json:"integrationPartnerID"`   // Integration Partner
	JobNotes               *string  `json:"jobNotes"`               // Job Notes -> note
	StatusID               *uint8   `json:"statusID"`               // Job Status -> progress
	MaterialDeliveryDate   *int64   `json:"materialDeliveryDate"`   // Material Delivery Date	Unix datetime
	PONumber               *string  `json:"poNumber"`               // PO Number
	ProductionImages       *string  `json:"productionImages"`       // Production Images Unrestricted URL
	ProgressInspectionDate *int64   `json:"progressInspectionDate"` // Progress Inspection Date	Unix datetime
	RequesterID            *string  `json:"requesterID"`            // Requester -> creator
	SalesRepID             *string  `json:"salesRepID"`             // Sales Rep -> sales rep
	ShingleColor           *string  `json:"shingleColor"`           // Shingle Color V2
	SolarCompanyID         *string  `json:"solarCompanyID"`         // Solar Company -> Requester
	WorkOrderPrice         *float64 `json:"workOrderPrice"`         // Work Order Price -> Partner Price
}

// EstimateApprovedHandler change from roofix bubble app
func EstimateApprovedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.EstApproved
	apiUID := account.CtxApiUserID(ctx)
	failed := func(err error) {
		audit.ApiOpFailed(ctx, action, apiUID, err)
		// return back with 200 with FailureStatus
		response.Ok(w, &nearmap.EstimateDetail{
			FailureStatus: err.Error(),
		})
	}

	var inp EstimateApprovedInput
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		failed(err)
		return
	}

	if err := estimateApproved(ctx, apiUID, &inp); err != nil {
		failed(err)
		return
	}

	audit.OpSuccess(ctx, action, fmt.Sprintf("approved estimate: %s", inp.ExternalID))
	response.Ok(w, nil)
}

func estimateApproved(ctx context.Context, apiUID string, inp *EstimateApprovedInput) error {
	if err := validate.Struct(inp); err != nil {
		return err
	}

	// approve estimate
	jobID, err := estimate.Approve(ctx, &model.ApproveEstimateInput{
		ID:  inp.ExternalID,
		Epc: inp.EPCCompanyID,
	}, nil, &apiUID)
	if err != nil {
		log.Error(err)
		return err
	}

	// InspectionDocURL:       inp.FinalInspectionDoc,             // Final Inspection Doc
	// ProdImagesURL:          inp.ProductionImages,               // Production Images
	// EpcID:

	// update job ==>
	update := ent.UpdateJobInput{
		ChangeOrderPrice:       inp.ChangeOrderPrice,               // Change Order Price
		ContractPrice:          inp.ContractPrice,                  // Contract Price
		RoofingPartnerID:       inp.ContractorID,                   // Contractor
		InspectionDate:         toTime(inp.FinalInspectionDate),    // Final Inspection Date
		InstallDate:            toTime(inp.InstallDate),            // Install Date
		IntegrationPartnerID:   inp.IntegrationPartnerID,           // Integration Partner
		MaterialDeliveryDate:   toTime(inp.MaterialDeliveryDate),   // Material Delivery Date
		PoNumber:               inp.PONumber,                       // PO Number
		ProgressInspectionDate: toTime(inp.ProgressInspectionDate), // Progress Inspection Date
		ShingleColor:           inp.ShingleColor,                   // Shingle Color V2
		WorkOrderPrice:         inp.WorkOrderPrice,                 // Work Order Price
	}

	db := ent.GetClient()
	defer db.CloseClient()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Job Status
	progress := toProgress(inp.StatusID)
	if progress != nil {
		now := time.Now().UTC()
		update.Progress = progress
		update.ProgressAt = &now

		err = tx.JobProgressHistory.Create().
			SetJobID(jobID).
			SetStatus(*progress).
			SetComplete(true).
			SetCreatorAPIUserID(apiUID).
			Exec(ctx)

		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	// save
	if err = tx.Job.UpdateOneID(jobID).SetInput(update).Exec(ctx); err != nil {
		return err
	}
	err = tx.Commit()

	return err
}

func toTime(v *int64) *time.Time {
	if v == nil {
		return nil
	}

	t := time.Unix(*v, 0).UTC()
	return &t
}

func toProgress(v *uint8) *enum.JobProgress {
	if v == nil {
		return nil
	}
	var p enum.JobProgress
	switch *v {
	case 2:
		p = enum.JobProgressNew
	case 4:
		p = enum.JobProgressPermitting
	case 5:
		p = enum.JobProgressScheduled
	case 6:
		p = enum.JobProgressInProgress
	case 7:
		p = enum.JobProgressDelayed
	case 8:
		p = enum.JobProgressInstalled
	default:
		return nil
	}

	return &p
}
