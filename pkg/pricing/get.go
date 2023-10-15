/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/ent/postalcode"
	"roofix/ent/pricing"
	"strings"
)

// ForAll from database
// we do need to add country filter if needed in future
func ForAll(ctx context.Context) ([]*ProductPrice, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var res []*ProductPrice
	err := pricingQry(db).Select().Scan(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// priceList for a given Zip location.
// If exact zip code price are found in pricing table, return them.
// Else find the nearest zip and return its pricing
//
// Returns:
//
//	pricing, note, error
func priceList(ctx context.Context, state, zip string, lat, lng float64) ([]*ProductPrice, string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	country := "US"
	// check we have an exact match
	hasPricing, err := db.Pricing.Query().Where(pricing.PostalCountry(country), pricing.PostalCode(zip)).Exist(ctx)
	if err != nil {
		return nil, "", err
	}

	// pricing for zip exists
	var res []*ProductPrice
	if hasPricing {
		err = pricingQry(db).
			Where(pricing.PostalCountry(country), pricing.PostalCode(zip)).
			Select().
			Scan(ctx, &res)
		if err != nil {
			return nil, "", err
		} else if len(res) == 0 {
			return nil, "", errors.New(fmt.Sprintf("pricing for %s not found", zip))
		}

		return res, fmt.Sprintf("Pricing by zip: %s", zip), nil
	}

	// need to find the nearest city in the state
	city, nearestZip, miles, err := nearestPricingZip(ctx, country, state, lat, lng)
	if err != nil {
		return nil, "", err
	} else if nearestZip == "" {
		msg := fmt.Sprintf("no direct(zip: %s) or nearby(state: %s, lat: %f, lng: %f) pricing found", zip, state, lat, lng)
		return nil, "", errors.New(msg)
	}

	err = pricingQry(db).
		Where(pricing.PostalCountry(country), pricing.PostalCode(nearestZip)).
		Select().
		Scan(ctx, &res)

	if err != nil {
		return nil, "", err
	} else if len(res) == 0 {
		return nil, "", errors.New(fmt.Sprintf("pricing for %s not found", zip))
	}

	return res, fmt.Sprintf("Pricing by nearest zip: %s, %s - %s (%.2f miles)", state, city, nearestZip, miles), nil
}

// nearestPricingZip zip will use the country & state and will
func nearestPricingZip(ctx context.Context, country, state string, lat, lng float64) (city, zip string, miles float32, err error) {
	db := ent.SqlDB()
	qry := `select distinct t2.city, t2.code as zip,
	(0.621371 * 6371 * ACOS(
		COS(RADIANS(t2.latitude)) * COS(RADIANS(?)) *
		COS(RADIANS(t2.longitude) - RADIANS(?)) +
		SIN(RADIANS(t2.latitude)) * SIN(RADIANS(?))
	)) AS miles
from pricing as t1 left join postal_codes t2 on t1.postal_id = t2.id
where t2.country = ? and (t2.state = ? or t2.state_abr = ?)
order by miles LIMIT 1`
	rows, err := db.QueryContext(ctx, qry, lat, lng, lat, country, state, strings.ToUpper(state))
	if err != nil {
		return city, zip, miles, err
	}

	if !rows.Next() {
		return city, zip, miles, err
	}

	err = rows.Scan(&city, &zip, &miles)
	return city, zip, miles, err
}

func AllProducts() []*ProductInfo {
	return []*ProductInfo{
		{ID: RRLaminateAR.Int(), Name: string(RRLaminateAR)},
		{ID: RRStandingSeam.Int(), Name: string(RRStandingSeam)},
		{ID: RRConcreteTile.Int(), Name: string(RRConcreteTile)},
		{ID: RRClayTile.Int(), Name: string(RRClayTile)},
		{ID: RRModBit.Int(), Name: string(RRModBit)},
		{ID: CoatingProduct.Int(), Name: string(Coating)},
		{ID: WoodShakeTearOff.Int(), Name: string(WoodShakeTearOff)},
		{ID: AddLayerFelt.Int(), Name: string(AddLayerFelt)},
		{ID: LayerOfShingles.Int(), Name: string(LayerOfShingles)},
		{ID: LayerOfMetal.Int(), Name: string(LayerOfMetal)},
		{ID: Sheathing1by2.Int(), Name: string(Sheathing1by2)},
		{ID: Steep7by12.Int(), Name: string(Steep7by12)},
		{ID: Steep10by12.Int(), Name: string(Steep10by12)},
		{ID: TwoStory.Int(), Name: string(TwoStory)},
		{ID: BasicPerSqAvg.Int(), Name: string(BasicPerSqAvg)},
		{ID: DisposalFees.Int(), Name: string(DisposalFees)},
		{ID: PermitFees.Int(), Name: string(PermitFees)},
	}
}

func pricingQry(db *ent.Client) *ent.PricingQuery {
	return db.Pricing.
		Query().
		Where(func(t1 *sql.Selector) {
			t2 := sql.Table(postalcode.Table)
			t1.Join(t2).On(t1.C(pricing.PostalColumn), t2.C(postalcode.FieldID))
			t1.OrderBy(t2.C(postalcode.FieldState), t2.C(postalcode.FieldCity), t1.C(pricing.FieldProductID))
			t1.Select(
				t1.C(pricing.FieldID),
				t1.C(pricing.FieldPostalCountry),
				t2.C(postalcode.FieldState),
				t2.C(postalcode.FieldStateAbr),
				t2.C(postalcode.FieldCity),
				t1.C(pricing.FieldPostalCode),
				t1.C(pricing.FieldPrice),
				t1.C(pricing.FieldPricePer),
				t1.C(pricing.FieldProductID),
			)
		})
}
