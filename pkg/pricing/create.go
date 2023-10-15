/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/pkg/msg"
	"roofix/pkg/postal"
	"roofix/pkg/util/crypt"
)

func CreateBulk(ctx context.Context, list []*Input) error {
	if list == nil {
		return msg.AsError(msg.ParamMissing, "list")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// iterate
	var productID uint8
	var id, postalID string
	for _, p := range list {
		productID = p.Product.UInt8()
		id = ID(p.PostalCountry, p.PostalCode, productID)
		postalID = postal.ID(p.PostalCountry, p.PostalCode)
		_, err := db.Pricing.Create().
			SetID(id).
			SetPostalID(postalID).
			SetPostalCountry(p.PostalCountry).
			SetPostalCode(p.PostalCode).
			SetPrice(p.Price).
			SetPricePer("SQ").
			SetNillableDescription(p.Description).
			SetProductID(productID).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// ID is MD5Hash of country, zip code & materialID
func ID(country, zip string, material uint8) string {
	return crypt.MD5Hash(fmt.Sprintf("%s-%s-%d", country, zip, material))
}
