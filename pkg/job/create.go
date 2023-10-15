package job

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
	"time"
)

const PartnerPriceMultiplyFactor = 0.9 // less by 10%

type Input struct {
	EstimateID       string
	RegionID         uint8
	Price            float64
	HomeOwnerID      *string
	SalesRepID       *string
	PartnerID        *string // requesting partner id
	EpcPartnerID     *string // requesting partner id
	CompanyRefName   *string
	EstimatePdfDocID *string
	CtxUserID        *string
	CtxAPIUserID     *string
}

func Create(ctx context.Context, tx *ent.Tx, input *Input) (string, error) {
	var wop float64
	if input.Price > 0 {
		wop = input.Price * PartnerPriceMultiplyFactor
	}
	qry := tx.Job.Create().
		SetEstimateID(input.EstimateID).
		SetProgress(enum.JobProgressNew).
		SetProgressAt(time.Now().UTC()).
		SetRegionID(input.RegionID).
		SetPrice(input.Price).
		SetContractPrice(input.Price).
		SetWorkOrderPrice(wop).
		SetCompanyName(str.Val(input.CompanyRefName)).
		SetNillableHomeOwnerID(input.HomeOwnerID).
		SetNillableSalesRepID(input.SalesRepID).
		SetNillableRequesterID(input.PartnerID).
		SetNillableEpcPartnerID(input.EpcPartnerID).
		SetNillableEstimatePdfID(input.EstimatePdfDocID).
		SetNillableCreatorID(input.CtxUserID).
		SetNillableCreatorAPIID(input.CtxAPIUserID)

	// save, job
	if j, err := qry.Save(ctx); err != nil {
		return "", err
	} else {
		return j.ID, err
	}
}
