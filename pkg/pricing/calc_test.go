/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"context"
	"roofix/pkg/enum"
	"roofix/pkg/material"
	"roofix/pkg/model"
	"testing"
)

func TestCalc(t *testing.T) {
	ctx := context.Background()
	t.Run("nil job data check", func(t *testing.T) {
		if _, err := Calculate(ctx, nil, []*Measurement{}); err == nil {
			t.Error("error was expected")
		}
	})

	t.Run("zip is required", func(t *testing.T) {
		if _, err := Calculate(ctx, &Data{}, []*Measurement{}); err == nil {
			t.Error("error was expected")
		}
	})

	t.Run("sq is required", func(t *testing.T) {
		if _, err := Calculate(ctx, &Data{Zip: "12345"}, []*Measurement{}); err == nil {
			t.Error("error was expected")
		}
	})

	t.Run("exact match", func(t *testing.T) {
		j := &Data{
			State:              "Alabama",
			Zip:                "35233",
			Lat:                33.5118197,
			Lng:                -86.8011326,
			CurrentMaterial:    material.ThreeTabShingles,
			NewRoofingMaterial: material.ConcreteTile,
			LowSlope:           false,
			Redeck:             true,
			Layers:             1,
		}
		if p, err := Calculate(ctx, j, []*Measurement{
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         3,
				Pitch:           7,
			},
		}); err != nil {
			t.Error("not expecting err", err)
		} else {
			t.Log(p)
		}
	})

	t.Run("nearby zip match", func(t *testing.T) {
		j := &Data{
			State:                      "New York",
			Zip:                        "10018",
			Lat:                        40.7524632,
			Lng:                        -73.9847395,
			CurrentMaterial:            material.Shingles,
			NewRoofingMaterial:         material.MoreExpensive,
			LowSlope:                   true,
			CurrentMaterialLowSlope:    material.RolledRoofing,
			NewRoofingMaterialLowSlope: material.Coating,
			Redeck:                     false,
			Layers:                     3,
			Layer2Material:             material.ClayTile,
			Layer3Material:             material.Shingles,
			PartialPercentage:          60,
		}
		if p, err := Calculate(ctx, j, []*Measurement{
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         17,
				Pitch:           8,
			},
		}); err != nil {
			t.Error("not expecting err", err)
		} else {
			t.Log(p)
		}
	})

	t.Run("repaper", func(t *testing.T) {
		j := &Data{
			State:              "New York",
			Zip:                "10018",
			Lat:                40.7524632,
			Lng:                -73.9847395,
			CurrentMaterial:    material.ConcreteTile,
			NewRoofingMaterial: material.Repaper,
			Layers:             1,
			PartialPercentage:  30,
		}
		if p, err := Calculate(ctx, j, []*Measurement{
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         17,
				Pitch:           10,
			},
		}); err != nil {
			t.Error("not expecting err", err)
		} else {
			t.Log(p)
		}
	})

	t.Run("Extra Charges", func(t *testing.T) {
		j := &Data{
			State:                      "New York",
			Zip:                        "10018",
			Lat:                        40.7524632,
			Lng:                        -73.9847395,
			CurrentMaterial:            material.Shingles,
			NewRoofingMaterial:         material.MoreExpensive,
			LowSlope:                   true,
			CurrentMaterialLowSlope:    material.RolledRoofing,
			NewRoofingMaterialLowSlope: material.Coating,
			Redeck:                     false,
			Layers:                     3,
			Layer2Material:             material.ClayTile,
			Layer3Material:             material.Shingles,
			ExtraCharge: &ExtraCharge{
				Type: enum.ExtraChargeConditional,
				Conditions: []*model.ExtraChargeCondition{
					{
						Condition:    enum.ConditionSubtractIfStateIn,
						ConditionVal: []string{"Texas"},
						Val:          250,
					},
					{
						Condition:    enum.ConditionAddIfStateIn,
						ConditionVal: []string{"Oklahoma", "Arkansas", "new york"},
						Val:          1000,
					},
				},
			},
		}
		if p, err := Calculate(ctx, j, []*Measurement{
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         17,
				Pitch:           6.5,
			},
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         17,
				Pitch:           1.5,
			},
		}); err != nil {
			t.Error("not expecting err", err)
		} else {
			t.Log(p)
		}
	})

	t.Run("random", func(t *testing.T) {
		j := &Data{
			State:                      "New York",
			Zip:                        "10011",
			Lat:                        40.747812,
			Lng:                        -74.007014,
			CurrentMaterial:            material.ThreeTabShingles,
			NewRoofingMaterial:         material.ModBit,
			LowSlope:                   true,
			NewRoofingMaterialLowSlope: material.ModBit,
			Redeck:                     true,
			Layers:                     3,
			Layer2Material:             material.ThreeTabShingles,
			Layer3Material:             material.ThreeTabShingles,
		}
		if p, err := Calculate(ctx, j, []*Measurement{
			{
				CurrentMaterial: CurrentMaterial(j.CurrentMaterial),
				Squares:         20.77,
				Pitch:           7.39,
			},
		}); err != nil {
			t.Error("not expecting err", err)
		} else {
			t.Log(p)
		}
	})
}
