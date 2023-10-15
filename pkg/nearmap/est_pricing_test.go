package nearmap

import (
	"context"
	"encoding/json"
	"os"
	"roofix/pkg/enum"
	"roofix/pkg/material"
	"testing"
)

func TestTotalSqFt(t *testing.T) {
	p := roofingPitch(30.0)
	f := 0.0 //3100.0 // flat sq. ft

	expected := 3579.8
	actual := totalSqFt(p, f)
	if actual != expected {
		t.Error("expected was ", expected, "but got", actual)
	}
}

func TestPrintPitch(t *testing.T) {
	//p := roofingPitch(16.9)
	//t.Log("roofing pitch:", p)

	//var res Response
	//b, _ := os.ReadFile("./sample/test.json")
	//_ = json.Unmarshal(b, &res)
	//flag := isNil(PitchDegreePrimaryRoof, res.Rollups)
	//t.Log("PitchDegree is null:", flag)

	// set until a month old
	//if tm, err := time.Parse(time.DateOnly, "2022-03-17"); err != nil {
	//	t.Error(err)
	//} else {
	//	tm = util.LastMonth(tm)
	//	t.Log("date:", tm.Format(time.DateOnly))
	//}
}

func TestGetPrice(t *testing.T) {
	j := &Data{
		StreetAddress:              "548 West 22nd Street",
		City:                       "New York",
		State:                      "New York",
		Zip:                        "10011",
		Latitude:                   40.747812,
		Longitude:                  -74.007014,
		CurrentMaterial:            material.ThreeTabShingles,
		NewRoofingMaterial:         material.BestValue,
		LowSlope:                   true,
		CurrentMaterialLowSlope:    material.BlackMembrane,
		NewRoofingMaterialLowSlope: material.ModBit,
		Redeck:                     true,
		Layers:                     1,
		Layer2Material:             material.ThreeTabShingles,
		Layer3Material:             material.ThreeTabShingles,
		EstimateType:               enum.MeasurePrimaryStructureOnly,
		PartialPercentage:          30,
	}

	var res Response
	b, _ := os.ReadFile("./sample/02.json")
	_ = json.Unmarshal(b, &res)

	ctx := context.Background()
	p, err := getPrice(ctx, j.EstimateType, pricingData(j), res.Rollups)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(p.Summary)
}
