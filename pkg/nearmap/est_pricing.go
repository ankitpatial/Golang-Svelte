package nearmap

import (
	"context"
	"fmt"
	"math"
	"roofix/pkg/enum"
	"roofix/pkg/material"
	"roofix/pkg/pricing"
	"roofix/pkg/util/num"
	"strings"
)

type Input struct {
	MeasurementType enum.JobMeasurement
	CurrentMaterial string
	Zip             string
}

type Price struct {
	TotalSq     float64
	Pitch       float64
	PitchDegree float64
	Amt         float64
	Summary     string
	Material    string
	IsRedeck    bool
}

type Detail struct {
	// Area is total SQ
	Area     float64
	Pitch    float64
	Material string
	IsRedeck bool
}

func getPrice(ctx context.Context, mt enum.Measure, inp *pricing.Data, rollups map[string]*Rollup) (*Price, error) {
	// materials
	var ts float64
	var pm []*pricing.Measurement
	var sb strings.Builder
	sb.WriteString("### NearMap Response  \n")
	isPrimaryStructure := mt == enum.MeasurePrimaryStructureOnly
	d := FillEstimateDetail(rollups, isPrimaryStructure)
	for _, v := range d.RoofArea {
		tsf := v.SqFt
		sq := tsf / 100
		ts += sq
		if sq == 0 { // ignore the 0 sq entries
			continue
		}

		if v.IsFlat { // flat roof area returned from nearmap
			sb.WriteString(fmt.Sprintf(
				"%s = **%s SF**  \n", v.Description, num.RoundAndFormatFloat(v.SqFt),
			))
			pm = append(pm, &pricing.Measurement{
				CurrentMaterial: pricing.CurrentMaterial(material.LowSlopeOnly),
				Squares:         sq,
				Pitch:           0,
				WasteFactor:     d.FlatWasteFactor,
				IsFlat:          true,
			})
		} else {
			sb.WriteString(fmt.Sprintf(
				"%s = **%s SF** primary pitch = %0.2f  \n", v.Description, num.RoundAndFormatFloat(v.SqFt), d.PrimaryPitch,
			))
			pm = append(pm, &pricing.Measurement{
				CurrentMaterial: pricing.CurrentMaterial(inp.CurrentMaterial),
				Squares:         sq,
				Pitch:           d.PrimaryPitch,
				WasteFactor:     d.SteepWasteFactor,
			})
		}
	}

	// == price calc ==
	summary, err := pricing.Calculate(ctx, inp, pm)
	if err != nil {
		return nil, err
	}

	sb.WriteString(summary.Summary)

	return &Price{
		TotalSq:  summary.SQ,
		Pitch:    summary.Pitch,
		Amt:      num.RoundFloat(summary.Total, 2),
		Summary:  sb.String(),
		Material: summary.NewMaterial,
		IsRedeck: summary.Redeck,
	}, nil
}

// roofingPitch is X/Y where Y = 12 and X = tan(P) * Y
//
// Variables:
//
//	P = Roof Pitch (degrees)
//	X = Rise (2 decimal places)
//	Y = Run = 12
func roofingPitch(p float64) float64 {
	x := math.Tan(p*math.Pi/180) * 12
	return num.RoundFloat(x, 2)
}

// totalSqFt is actual roof area from flat sq. ft with available roofing pitch i.e. SqFt * Multiplier
//
// Formulas:
//
//	T = (F/Y) * (√(x^2+y^2))
//	T = 3580 ft^2
//	M = 1 / cos(P)
//	M = 1.155
//
// Variables:
//
//	F = Flat SqFt
//	T = Total SqFt
//	M = Multiplier (3 decimal places)
//	P = Roof Pitch (degrees)
//	Y = Run = 12
func totalSqFt(x float64, flatSqFt float64) float64 {
	if flatSqFt == 0 {
		return 0
	}

	// h = √(x^2 + y^2)
	h := math.Sqrt(math.Pow(x, 2) + 144.0)
	// w = f/y
	w := flatSqFt / 12
	// T = (F/Y)*(√(x^2+y^2))
	t := w * h
	return num.RoundFloat(t, 2)
}
