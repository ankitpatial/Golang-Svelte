package estimate

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/document"
	"roofix/ent/estimate"
	"roofix/ent/homeowner"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/job"
	"roofix/pkg/util/log"
	"roofix/server/model"
)

// Approve estimate & create job
// returns jobID
func Approve(ctx context.Context, input *model.ApproveEstimateInput, ctxUID, ctxApiUID *string) (string, error) {
	log.Info("approve estimate")
	db := ent.GetClient()
	defer db.CloseClient()

	// pull estimate
	est, err := db.Estimate.Query().
		WithHomeOwner(func(ho *ent.HomeOwnerQuery) { ho.Select(homeowner.FieldID) }).
		WithPartner(func(p *ent.PartnerQuery) { p.Select(partner.FieldID) }).
		WithSalesRep(func(rep *ent.UserQuery) { rep.Select(user.FieldID) }).
		WithCreator(func(c *ent.UserQuery) { c.Select(user.FieldID) }).
		WithPdf(func(d *ent.DocumentQuery) { d.Select(document.FieldID) }).
		Where(estimate.ID(input.ID)).
		Only(ctx)
	if err != nil {
		return "", err
	} else if est.Status != enum.EstimateStatusPending { // only pending status cna be approved
		return "", fmt.Errorf("estimate status is not Pending")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	// 💾 update estimate status
	estStatus := enum.EstimateStatusApproved
	err = tx.Estimate.UpdateOneID(est.ID).SetStatus(estStatus).Exec(ctx)
	if err != nil {
		return "", err
	} else {
		log.Info("saved estimate status")
	}

	// 💾 create estimate activity
	err = tx.EstimateActivity.Create().
		SetEstimateID(est.ID).
		SetDescription(fmt.Sprintf("changed status to %s", estStatus)).
		SetNillableCreatorID(ctxUID).
		SetNillableCreatorAPIID(ctxApiUID).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	} else {
		log.Info("saved estimate activity")
	}

	// 💾 save owner contact info
	if est.Edges.HomeOwner != nil && (input.OwnerEmail != nil || input.OwnerPhone != nil) {
		err = tx.HomeOwner.UpdateOneID(est.Edges.HomeOwner.ID).
			SetNillableEmail(input.OwnerEmail).
			SetNillablePhone(input.OwnerPhone).
			Exec(ctx)
		if err != nil {
			_ = tx.Rollback()
			return "", err
		}
		log.Info("saved owner contact info")
	}

	// new job ==>
	jInp := &job.Input{
		EstimateID:     est.ID,
		RegionID:       est.RegionID,
		Price:          est.Price,
		CompanyRefName: &est.CompanyRefName,
		CtxUserID:      ctxUID,
		CtxAPIUserID:   ctxApiUID,
	}
	// set home-owner
	if est.Edges.HomeOwner != nil {
		jInp.HomeOwnerID = &est.Edges.HomeOwner.ID
	}
	// set sales rep
	if est.Edges.SalesRep != nil {
		jInp.SalesRepID = &est.Edges.SalesRep.ID
	}
	// request or solar partner
	if est.Edges.Partner != nil {
		jInp.PartnerID = &est.Edges.Partner.ID
	}
	// Estimate PDF docID
	if est.Edges.Pdf != nil {
		jInp.EstimatePdfDocID = &est.Edges.Pdf.ID
	}

	// TODO: need to fix it
	//if input.Epc != nil {
	//    jInp.EpcPartnerID = &est.Edges.Partner.ID
	//}

	// 💾 create job
	if jobID, err2 := job.Create(ctx, tx, jInp); err2 != nil {
		_ = tx.Rollback()
		return "", err2
	} else {
		log.Info("created job")

		return jobID, tx.Commit()
	}

}
