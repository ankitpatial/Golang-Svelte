/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"errors"
	"roofix/ent"
	entEst "roofix/ent/estimate"
	pkgEst "roofix/pkg/estimate"
	"roofix/pkg/pricing"
	"roofix/pkg/task"
	"roofix/pkg/util/log"
	num2 "roofix/pkg/util/num"
	"roofix/pkg/util/parse"
	"roofix/pkg/util/str"
	"strconv"
	"strings"
)

type TaskGetReport struct {
	EstID string `json:"estID"`
}

func (t TaskGetReport) Name() task.Name {
	return task.EagleViewGetReport
}

// Process task, pull report from EagleView and save
func (t TaskGetReport) Process(ctx context.Context, progress chan string, done chan error) {
	progress <- "processing EagleView GetReport request"
	db := ent.GetClient()
	defer db.CloseClient()

	// 1. pull estimate from DB
	est, err := db.Estimate.Query().WithHomeOwner().Where(entEst.ID(t.EstID)).Only(ctx)
	if err != nil {
		done <- err
	}

	// 2. call API GetReport
	progress <- "calling EagleView API"
	rep, err := GetReport(ctx, est.EstimatorReportID)
	if err != nil {
		done <- err
		return
	}

	// 3. save Estimate Response
	progress <- "retrieved report"
	raw := parse.StructToMap(rep)
	//desc := fmt.Sprintf("Get Report: Status: %s, SubStatus: %s", rep.Status.String(), rep.SubStatus.String())
	/*
		err = db.EstimateResponse.Create().
			SetEstimateID(t.EstID).
			SetDescription(desc).
			SetRaw(raw).
			Exec(ctx)
	*/
	if rep.Measurement == nil {
		progress <- "report has no Measurements yet"
		done <- nil
		return
	}

	// 4. calculate price
	progress <- "calculating pricing"
	owner := est.Edges.HomeOwner
	data := &pricing.Data{
		State:                      owner.State,
		Zip:                        owner.Zip,
		Lat:                        owner.Latitude,
		Lng:                        owner.Longitude,
		CurrentMaterial:            est.CurrentMaterial,
		NewRoofingMaterial:         est.NewRoofingMaterial,
		LowSlope:                   est.LowSlope,
		CurrentMaterialLowSlope:    est.CurrentMaterialLowSlope,
		NewRoofingMaterialLowSlope: est.NewRoofingMaterialLowSlope,
		Redeck:                     est.Redeck,
		Layers:                     est.Layers,
		Layer2Material:             est.Layer2Material,
		Layer3Material:             est.Layer3Material,
		PartialPercentage:          est.PartialPercentage,
		ExtraCharge: &pricing.ExtraCharge{
			Type:       est.ExtraChargeType,
			Val:        est.ExtraCharges,
			Note:       str.Val(est.ExtraChargeNote),
			Conditions: est.ExtraChargeCond,
		},
	}
	p, err := pricing.Calculate(ctx, data, measurements(rep))
	if err != nil {
		log.Error(err)
		done <- errors.New("failed on calculating pricing")
		return
	}

	// 5. update measurement order
	inp := &pkgEst.Estimation{
		Estimator:   ProviderName,
		SQ:          rep.Measurement.AreaValue / 100, // sq.ft -> squares,
		Pitch:       rep.Measurement.PitchValue,
		Price:       p.Total,
		Bounds:      nil,
		Summary:     p.Summary,
		RawResponse: raw,
	}
	if err = pkgEst.UpdateEstimation(ctx, db, t.EstID, inp); err != nil {
		log.Error(err)
		done <- errors.New("failed to save Estimate Pricing detail")
		return
	}

	// 6. set job status to ESTIMATED
	//job.Estimated(ctx, est.Edges.Job.ID, "")

	// done!
	done <- nil
}

func measurements(res *ReportDetail) []*pricing.Measurement {
	// use AreaValue & PitchValue if PitchTable is empty
	// PitchTable is coming empty for staging api, this may not be the case for production api but still lets handle it.
	var tbl []*pricing.Measurement
	if len(res.Measurement.PitchTable) == 0 {
		tbl = append(tbl, &pricing.Measurement{
			Squares: res.Measurement.AreaValue / 100,
			Pitch:   res.Measurement.PitchValue,
		})
		return tbl
	}

	// use pitch table
	for _, p := range res.Measurement.PitchTable {
		tbl = append(tbl, &pricing.Measurement{
			Squares: toSQ(p.Area),
			Pitch:   toPitch(p.Pitch),
		})
	}

	return tbl
}

func toSQ(str string) float64 {
	v := strings.Replace(strings.ToLower(str), " sq. ft", "", 1) // e.g.  1234 sq. ft
	num, err := strconv.ParseFloat(v, 32)
	if err != nil {
		log.Error(err)
		return 0
	}

	// convert to SQ.
	// 1 sq is 100 sq. ft
	return num2.RoundFloat(num/100, 2)
}

func toPitch(str string) float64 {
	v := strings.Replace(str, "/12", "", 1) // e.g. 10/12
	num, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Error(err)
		return 0
	}
	return num
}

// SubmitGetReport request to EagleView
func SubmitGetReport(ctx context.Context, estID string) error {
	return task.Submit(ctx, &TaskGetReport{EstID: estID})
}
