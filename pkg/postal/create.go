/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"strings"
)

func CreateBulk(ctx context.Context, list []*Detail) error {
	log.Info("CreateBulk")
	db := ent.GetClient()
	defer db.CloseClient()

	for _, p := range list {
		if p.State == "" {
			continue
		}

		country := strings.ToUpper(p.Country)
		err := db.PostalCode.
			Create().
			SetID(ID(country, p.Zip)).
			SetCountry(country).
			SetCode(p.Zip).
			SetCity(p.City).
			SetState(p.State).
			SetStateAbr(p.StateAbr).
			SetRegionID(GetStateRegion(p.State)).
			SetLatitude(p.Latitude).
			SetLongitude(p.Longitude).
			SetAccuracy(p.Accuracy).
			Exec(ctx)

		if err != nil {
			log.Warn("failed to save %s", p.Zip)
			return err
		}
	}
	return nil
}

// ID is MD5Hash of country and zip code
func ID(country, zip string) string {
	return crypt.MD5Hash(fmt.Sprintf("%s-%s", country, zip))
}
