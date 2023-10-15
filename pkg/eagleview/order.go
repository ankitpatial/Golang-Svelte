/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"encoding/json"
	"fmt"
	"roofix/ent"
	"roofix/pkg/util/http"
	"roofix/pkg/util/log"
)

type Order struct {
	List []*Report `json:"OrderReports"`
}

type Report struct {
	ReferenceId                string        `json:"ReferenceId"`
	Addresses                  []*Address    `json:"ReportAddresses"`
	PrimaryProductId           int           `json:"PrimaryProductId"`
	DeliveryProductId          int           `json:"DeliveryProductId"`
	MeasurementInstructionType int           `json:"MeasurementInstructionType"`
	ChangesInLast4Years        bool          `json:"ChangesInLast4Years"`
	ReportAttributes           []*ReportAttr `json:"ReportAttibutes,omitempty"`
}

type Address struct {
	Address     string  `json:"Address"`
	City        string  `json:"City"`
	State       string  `json:"State"`
	Zip         string  `json:"Zip"`
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
	AddressType int     `json:"AddressType"`
	//Country
}

type ReportAttr struct {
	Attribute int
	Value     string
}

type OrderResponse struct {
	ID        uint   `json:"OrderId"`
	ReportIDs []uint `json:"ReportIds"`
}

// PlaceOrder will submit a measurement order to EagleView
// doc ref: https://restdoc.eagleview.com/#PlaceOrder
func PlaceOrder(ctx context.Context, job *ent.Estimate, refID string) (*OrderResponse, error) {
	// get api token detail
	host, token, err := apiToken(ctx)
	if err != nil {
		return nil, err
	}

	// order data\
	order := fromJob(job, refID)
	// post
	b, err := http.Post(fmt.Sprintf("%s/v2/Order/PlaceOrder", host), order, http.WithAuthBearer(token), http.WithContentTypeJson())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("order is submitted successfully")

	// parse response
	var res OrderResponse
	if err := json.Unmarshal(b, &res); err != nil {
		log.Error(err)
		log.Warn("failed to parse order response")
		return nil, nil
	}

	// done
	return &res, nil
}

func fromJob(j *ent.Estimate, redID string) *Order {
	ho := j.Edges.HomeOwner
	return &Order{
		List: []*Report{
			{
				ReferenceId: redID,
				Addresses: []*Address{
					{
						Address:     fmt.Sprintf("%s %s", ho.StreetNumber, ho.StreetName),
						City:        ho.City,
						State:       ho.State,
						Zip:         ho.Zip,
						Latitude:    ho.Latitude,
						Longitude:   ho.Longitude,
						AddressType: 4, // 4, Use this value if the user specifies a latitude and longitude.
					},
				},
				PrimaryProductId:  44, // 44, QuickSquares
				DeliveryProductId: 45, // 45, QuickDelivery
				//MeasurementInstructionType: int(j.MeasurementType),
				ChangesInLast4Years: false,
			},
		},
	}
}
