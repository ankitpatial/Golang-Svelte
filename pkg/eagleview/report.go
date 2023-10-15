/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"encoding/json"
	"fmt"
	"roofix/pkg/util/http"
	"roofix/pkg/util/log"
)

type ReportDetail struct {
	ReportId           int          `json:"ReportId"`
	Street             string       `json:"Street"`
	DatePlaced         string       `json:"DatePlaced"`
	ReportDownloadLink string       `json:"ReportDownloadLink"`
	EligibleForUpgrade bool         `json:"EligibleForUpgrade"`
	Status             Status       `json:"StatusId"`
	SubStatus          SubStatus    `json:"SubStatusId"`
	ProductPrimary     string       `json:"ProductPrimary"`
	ProductDelivery    string       `json:"ProductDelivery"`
	Measurement        *Measurement `json:"TotalMeasurements"`
}

type Measurement struct {
	BuildingName     string           `json:"BuildingName"`
	Area             string           `json:"Area"`
	PrimaryPitch     string           `json:"PrimaryPitch"`
	LengthRidge      string           `json:"LengthRidge"`
	LengthValley     string           `json:"LengthValley"`
	LengthEave       string           `json:"LengthEave"`
	LengthRake       string           `json:"LengthRake"`
	LengthHip        string           `json:"LengthHip"`
	WallsSidingArea  string           `json:"WallsSidingArea"`
	WallsMasonryArea string           `json:"WallsMasonryArea"`
	AreaValue        float64          `json:"AreaValue"`
	PitchValue       float64          `json:"PitchValue"`
	TotalRoofFacets  string           `json:"TotalRoofFacets"`
	PitchTable       []*AreaWithPitch `json:"PitchTable"`
}

type AreaWithPitch struct {
	Area           string `json:"RoofArea"`
	Pitch          string `json:"Pitch"`
	AreaPercentage string `json:"PercentageRoofArea"`
}

// GetReport from EagleView
// doc ref: https://restdoc.eagleview.com/#V3GetReport
func GetReport(ctx context.Context, reportID uint) (*ReportDetail, error) {
	// get api token detail
	host, token, err := apiToken(ctx)
	if err != nil {
		return nil, err
	}

	b, err := http.Get(fmt.Sprintf("%s/v3/Report/GetReport?reportId=%d", host, reportID), http.WithAuthBearer(token), http.WithContentTypeJson())
	if err != nil {
		log.Error(err)
		if len(b) > 0 {
			var er ErrResp
			_ = json.Unmarshal(b, &er)
			return nil, er
		}
		return nil, err
	}

	// parse response
	var res ReportDetail
	if err := json.Unmarshal(b, &res); err != nil {
		log.Error(err)
		log.Warn("failed to parse order response")
		return nil, nil
	}

	// done
	return &res, nil
}
