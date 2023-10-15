/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"bytes"
	"encoding/csv"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"strconv"
)

func ReadUSPricingCsv(data []byte) []*Input {
	var result []*Input
	r := csv.NewReader(bytes.NewReader(data))
	list, err := r.ReadAll()
	if err != nil {
		log.Error(err)
		return nil
	}

	for _, rec := range list {
		co := "US"
		zip := rec[3]
		if _, err := strconv.Atoi(zip); err != nil {
			log.Warn("skip zip: %s", zip)
			continue
		}

		// LAMINATE AR
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       RRLaminateAR,
			Price:         num.ParseFloat(rec[4]),
		})

		// STANDING SEAM
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       RRStandingSeam,
			Price:         num.ParseFloat(rec[5]),
		})

		// CONCRETE TILE
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       RRConcreteTile,
			Price:         num.ParseFloat(rec[6]),
		})

		// CLAY TILE(Barrel)
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       RRClayTile,
			Price:         num.ParseFloat(rec[7]),
		})

		// MOD BIT
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       RRModBit,
			Price:         num.ParseFloat(rec[8]),
		})

		// COATING
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       CoatingProduct,
			Price:         num.ParseFloat(rec[9]),
		})

		// Wood Shake:Tear off, Haul, and Dispose
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       WoodShakeTearOff,
			Price:         num.ParseFloat(rec[11]),
		})

		// Add. Layer Felt:Tear off, Haul, and Dispose
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       AddLayerFelt,
			Price:         num.ParseFloat(rec[12]),
		})

		// Layer of Shingles & Felt:Tear off, Haul, and Dispose
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       LayerOfShingles,
			Price:         num.ParseFloat(rec[13]),
		})

		// Layer of Metal: No haul off
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       LayerOfMetal,
			Price:         num.ParseFloat(rec[14]),
		})

		// Sheathing 1/2" OSB+Clips
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       Sheathing1by2,
			Price:         num.ParseFloat(rec[15]),
		})

		// STEEP 7/12-9/12
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       Steep7by12,
			Price:         num.ParseFloat(rec[16]),
		})

		// STEEP 10/12-12/12
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       Steep10by12,
			Price:         num.ParseFloat(rec[17]),
		})

		// TWO STORY
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       TwoStory,
			Price:         num.ParseFloat(rec[18]),
		})

		// BASIC PER SQ AVE
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       BasicPerSqAvg,
			Price:         num.ParseFloat(rec[19]),
		})

		// Additional Disposal Fees - 30YD Dumpster
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       DisposalFees,
			Price:         num.ParseFloat(rec[21]),
		})

		// Additional Permit Fees
		result = append(result, &Input{
			PostalCountry: co,
			PostalCode:    zip,
			Product:       PermitFees,
			Price:         num.ParseFloat(rec[22]),
		})
	}

	return result
}
