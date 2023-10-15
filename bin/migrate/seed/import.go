/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package seed

import (
	"context"
	_ "embed"
	"encoding/json"
	"roofix/ent"
	"roofix/pkg/postal"
	price "roofix/pkg/pricing"
	"roofix/pkg/util/log"
	"strings"
)

//go:embed us-postal.json
var f []byte

//go:embed "Price-List-June-2023.csv"
var usPricingCsv []byte

func USPostal(ctx context.Context) {
	db := ent.GetClient()
	defer db.CloseClient()

	//
	// US postal count is 40972
	//

	log.Info("reading postal json codes")
	var data []postal.Detail
	if err := json.Unmarshal(f, &data); err != nil {
		log.Error(err)
	}

	log.Info("saving %d postal codes...", len(data))
	for _, p := range data {
		if p.State == "" {
			continue
		}

		country := strings.ToUpper(p.Country)
		err := db.PostalCode.
			Create().
			SetID(postal.ID(country, p.Zip)).
			SetCountry(country).
			SetCode(p.Zip).
			SetCity(p.City).
			SetState(p.State).
			SetStateAbr(p.StateAbr).
			SetRegionID(postal.GetStateRegion(p.State)).
			SetLatitude(p.Latitude).
			SetLongitude(p.Longitude).
			SetAccuracy(p.Accuracy).
			Exec(ctx)

		if err != nil && strings.Contains(err.Error(), "Duplicate entry ") {
			log.Info("already exist %s", p.Zip)
		} else if err != nil {
			log.Warn("FAILED to save %s, error: %s", p.Zip, err.Error())
		}
	}

	log.Info("Done!")
}

func USPricing(ctx context.Context) {
	db := ent.GetClient()
	if _, err := db.Pricing.Delete().Exec(ctx); err != nil {
		log.Error(err)
		return
	}

	log.Info("seeding db with US pricing")
	data := price.ReadUSPricingCsv(usPricingCsv)
	if err := price.CreateBulk(ctx, data); err != nil {
		log.Error(err)
	}
}
