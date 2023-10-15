package estimate

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/nearmap"
	"roofix/pkg/task"
	"roofix/pkg/util/log"
	"roofix/pkg/util/parse"
	"roofix/pkg/util/str"
)

func RequestRealtime(ctx context.Context, inp *Input) (*nearmap.EstimateDetail, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// create estimate
	id, err := create(ctx, db, inp, false)
	if err != nil {
		return nil, err
	}

	// do area estimation & material pricing
	est, err := materialPricing(ctx, inp)
	if err != nil {
		markAsFailed(ctx, db, id, err.Error())
		notifyFailed(ctx, id)
		return nil, err
	}

	// update estimate
	p := est.Price
	err = UpdateEstimation(ctx, db, id, &Estimation{
		Estimator:   nearmap.ProviderName,
		SQ:          p.TotalSq,
		Pitch:       p.Pitch,
		Price:       p.Amt,
		Summary:     p.Summary,
		RawResponse: parse.StructToMap(est.Response),
		Bounds:      est.Boundary,
	})
	if err != nil { // log error, but continue
		log.Error(err)
	}

	isPrimary := inp.MeasurementType == enum.MeasurePrimaryStructureOnly
	resp := nearmap.FillEstimateDetail(est.Rollups, isPrimary)
	resp.ExternalID = str.Val(inp.ID)
	resp.Price = p.Amt
	resp.TotalSquares = p.TotalSq

	// raise notification
	notifyCompleted(ctx, id)

	// print PDF
	genPdf(ctx, id, inp, p)

	// DONE, return
	return resp, err
}

func Request(ctx context.Context, inp *Input) (string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	id, err := create(ctx, db, inp, true)
	if err != nil {
		return "", err
	}
	return id, task.Submit(ctx, BGTask{
		ID:    id,
		Input: *inp,
	})
}

func materialPricing(ctx context.Context, inp *Input) (*nearmap.EstimateInfo, error) {
	owner := inp.HomeOwner
	detail := inp.MaterialDetail
	return nearmap.Estimate(ctx, &nearmap.Data{
		StreetAddress:              fmt.Sprintf("%s %s", owner.StreetNumber, owner.StreetName),
		City:                       owner.City,
		State:                      owner.State,
		Zip:                        owner.Zip,
		Latitude:                   owner.Latitude,
		Longitude:                  owner.Longitude,
		CurrentMaterial:            detail.CurrentMaterial.String(),
		NewRoofingMaterial:         detail.NewRoofingMaterial.String(),
		LowSlope:                   detail.LowSlope,
		CurrentMaterialLowSlope:    detail.CurrentMaterialLowSlope.String(),
		NewRoofingMaterialLowSlope: detail.NewRoofingMaterialLowSlope.String(),
		Redeck:                     detail.Redeck,
		Layers:                     uint8(detail.Layers),
		Layer2Material:             detail.Layer2Material.String(),
		Layer3Material:             detail.Layer3Material.String(),
		PartialPercentage:          detail.PartialPercent,
		EstimateType:               inp.MeasurementType,
		ExtraCharge:                inp.ExtraCharge,
	})
}
