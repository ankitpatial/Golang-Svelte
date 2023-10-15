/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import (
	"context"
	"errors"
	"fmt"
	"math"
	"roofix/pkg/enum"
	"roofix/pkg/material"
	"roofix/pkg/model"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/slice"
	"strings"
)

const (
	estimateFee             = 750
	laborSurchargeThreshold = 7250
	laborSurcharge          = 1500
	commission              = 1 // % of material cost
	m50                     = .50
	m80                     = .80
	m85                     = .85
	m90                     = .90
	m105                    = 1.05
	none                    = 0
	add15                   = 15
	add20                   = 20
	add25                   = 25
	add30                   = 30
	add40                   = 40
	add90                   = 90
	sub5                    = -5
	sub20                   = -20
	sub30                   = -30
)

type (
	Summary struct {
		Total       float64
		Items       []*Entry
		Note        string
		Summary     string
		SQ          float64
		Pitch       float64
		NewMaterial string
		Redeck      bool
	}

	Entry struct {
		Existing string
		New      string
		Logic    string
		LogicVal string
		Amount   float64
	}

	Data struct {
		State                      string
		Zip                        string
		Lat                        float64
		Lng                        float64
		CurrentMaterial            string
		NewRoofingMaterial         string
		LowSlope                   bool
		CurrentMaterialLowSlope    string
		NewRoofingMaterialLowSlope string
		Redeck                     bool
		Layers                     uint8
		Layer2Material             string
		Layer3Material             string
		PartialPercentage          float64
		ExtraCharge                *ExtraCharge
	}

	ExtraCharge struct {
		Type       enum.ExtraCharge
		Val        float64
		Note       string
		Conditions []*model.ExtraChargeCondition
	}

	Measurement struct {
		CurrentMaterial Material
		Squares         float64 `json:"squares"`
		Pitch           float64 `json:"pitch"`
		// WasteFactor in %
		WasteFactor uint
		IsFlat      bool
	}

	Logic struct {
		existing         Material
		new              Material
		product          Product
		multiplier       float64
		extra            float64
		extraWithPartial float64
	}
)

var (
	materialLogic = []Logic{
		{Asphalt, BestValueArch, RRLaminateAR, m90, none, 0},
		{Asphalt, MoreExpensiveArch, RRLaminateAR, m105, none, 0},
		{Asphalt, StandingSeam, RRStandingSeam, m90, sub20, 0},
		{Asphalt, ConcreteTile, RRConcreteTile, m90, sub30, 0},
		{Asphalt, ClayTile, RRClayTile, m90, sub30, 0},
		{Asphalt, ModBit, RRModBit, m90, sub5, 0},

		{Metal, BestValueArch, RRLaminateAR, m90, add30, 0},
		{Metal, MoreExpensiveArch, RRLaminateAR, m90, add30, 0},
		{Metal, StandingSeam, RRStandingSeam, m90, none, 0},
		{Metal, ConcreteTile, RRConcreteTile, m90, sub20, 0},
		{Metal, ClayTile, RRClayTile, m90, sub20, 0},
		{Metal, ModBit, RRModBit, m90, add15, 0},

		{Tile, BestValueArch, RRLaminateAR, m90, add40, 200},
		{Tile, MoreExpensiveArch, RRLaminateAR, m90, add90, 0},
		{Tile, StandingSeam, RRStandingSeam, m90, add20, 0},
		{Tile, ConcreteTile, RRConcreteTile, m90, none, 0},
		{Tile, ClayTile, RRClayTile, m90, none, 0},
		{Tile, ModBit, RRModBit, m90, add30, 0},

		{Woodshake, BestValueArch, RRLaminateAR, m90, add30, 0},
		{Woodshake, MoreExpensiveArch, RRLaminateAR, m90, add30, 0},
		{Woodshake, ConcreteTile, RRConcreteTile, m90, sub20, 0},
		{Woodshake, ClayTile, RRClayTile, m90, sub20, 0},
		{Woodshake, ModBit, RRModBit, m90, add25, 0},

		{LowSlope, ModBit, RRModBit, m90, none, 0},
		{LowSlope, Coating, CoatingProduct, m90, none, 0},
	}

	layerLogic = []Logic{
		{UnknownMaterial, Asphalt, LayerOfShingles, m90, 0, 0},
		{UnknownMaterial, Metal, LayerOfMetal, m90, 0, 0},
		{UnknownMaterial, Woodshake, WoodShakeTearOff, m90, 0, 0},
	}
)

// Calculate job material pricing
func Calculate(ctx context.Context, inp *Data, list []*Measurement) (*Summary, error) {
	// check measurements length
	if len(list) == 0 {
		return nil, errors.New("no measurement found")
	}

	// get products with prices for job zip code
	prices, n, err := priceList(ctx, inp.State, inp.Zip, inp.Lat, inp.Lng)
	if err != nil {
		return nil, err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\n**%s**  \n", n))
	summary := &Summary{Note: n}
	// iterate measurement and do pricing
	var sq, p float64
	var nm string
	var redeck bool
	for _, m := range list {
		sq, p, nm, redeck, err = calc(inp, summary, prices, m, &sb)
		if err != nil {
			return nil, err
		}

		summary.SQ += sq
		summary.NewMaterial += fmt.Sprintf("%s. ", nm)
		if !m.IsFlat {
			summary.Pitch = p
			summary.Redeck = redeck
		}
	}

	total(inp, summary, &sb)
	return summary, nil
}

// calc for job material price calculation
func calc(
	inp *Data, summary *Summary, prices []*ProductPrice, m *Measurement, sb *strings.Builder,
) (sq, pitch float64, newMaterial string, redeck bool, err error) {
	// current material
	if m.CurrentMaterial == UnknownMaterial {
		err = errors.New(fmt.Sprintf("Current Material: '%s' is unknown.", inp.CurrentMaterial))
		return
	}

	// Steep Slope current material & pitch is 0 use default pitch 5/12
	if m.CurrentMaterial != LowSlope && m.Pitch == 0 {
		sb.WriteString(note("Pitch is 0, setting default pitch value i.e. 5/12"))
		m.Pitch = 5
	}

	// Low Slope current material & pitch >=  MinSteepSlope Pitch
	if m.CurrentMaterial == LowSlope && m.Pitch >= material.SteepSlopeMinPitch { // ignored
		sb.WriteString(note(
			fmt.Sprintf(
				"Setting pitch = 0, because current material is **Low Slope** and input pitch is %.2f which is >= %.2f",
				m.Pitch,
				material.SteepSlopeMinPitch,
			),
		))
		m.Pitch = 0
	}

	// If currentMaterial is Wood Shakes set redeck = true
	if m.CurrentMaterial == Woodshake && !inp.Redeck {
		sb.WriteString(note("Setting redeck = true (its required with Wood Shakes)"))
		inp.Redeck = true
	}

	pitch = m.Pitch
	redeck = inp.Redeck
	switch {
	// STEEP SLOP, anything below 2 is "Low Slope"
	case pitch >= material.SteepSlopeMinPitch:
		newMaterial = inp.NewRoofingMaterial
		nm := NewMaterial(newMaterial)
		if nm == UnknownMaterial {
			err = errors.New(fmt.Sprintf("new material '%s' is unknown", inp.NewRoofingMaterial))
			return
		}

		// waster factor, add default if missing
		if m.WasteFactor <= 0 {
			log.Info("* waste factor is missing, using default 15% waste factor")
			m.WasteFactor = material.WasteFactorSteep
		}

		// SQ with WasteFactor
		sq = m.Squares * wasteFactor(m.WasteFactor)
		msg := fmt.Sprintf("Area(%.2f SQ) * Waste Factor(%d%%)", m.Squares, m.WasteFactor)
		// SQ with Partial Percentage
		var isPartial bool
		if inp.PartialPercentage > 0 {
			sq *= inp.PartialPercentage / 100
			isPartial = true
			msg += fmt.Sprintf(" * Partial(%.0f%%)", inp.PartialPercentage)
		}

		sb.WriteString(fmt.Sprintf("\n### Steep Slope :: Area = %.2f SQ, Pitch = %.2f  \n", sq, pitch))
		sb.WriteString(fmt.Sprintf("%s = **%.2f SQ**  \n", msg, sq))

		if nm == material.Repaper { // REPAPER - Reuse Existing Tile
			cm := material.CurrentSteepSlope(inp.CurrentMaterial)
			err = repaperPricing(cm, prices, summary, sq, sb)
			if err != nil {
				return
			}
		} else { // STEEP SLOPE with layers
			err = steepSlopeMaterialPricing(inp, prices, summary, m.CurrentMaterial, nm, sq, isPartial, sb)
			if err != nil {
				return
			}
		}
	// LOW SLOPE
	default:
		newMaterial = inp.NewRoofingMaterialLowSlope
		nm := NewMaterial(newMaterial)
		// waste factor, add default if missing
		if m.WasteFactor <= 0 {
			log.Info("* waster factor is missing, using default 10% waste factor")
			m.WasteFactor = material.WasteFactorDefault
		}

		// SQ with WasteFactor( do not add partial % to low slope)
		sq = m.Squares * wasteFactor(m.WasteFactor)

		sb.WriteString(fmt.Sprintf("\n### Low Slope  :: Area = %.2f SQ  \n", sq))
		sb.WriteString(fmt.Sprintf("Area(%.2f SQ) * Waste Factor(%d%%) = **%.2f SQ**  \n", m.Squares, m.WasteFactor, sq))
		sb.WriteString("\n")

		var e *Entry
		e, err = lowSlopeMaterialPricing(prices, nm, sq)
		if err != nil {
			return
		}

		summary.Total += e.Amount
		summary.Items = append(summary.Items, e)
		reportEntry(sb, e)
	}

	// steep charges will apply if pitch >= 7
	if pitch >= 7.0 {
		steepCharges(sq, pitch, prices, summary, sb)
	}

	// re-deck
	if inp.Redeck {
		redeckPrice(prices, summary, sq, sb)
	}

	return
}

func total(d *Data, p *Summary, sb *strings.Builder) {
	if p.Total <= 0 {
		p.Summary = sb.String()
		return
	}

	format := func(a, b string) string {
		return fmt.Sprintf("| %s | %s |  \n", a, b)
	}

	sb.WriteString("\n| Description | Amount |  \n")
	sb.WriteString("|---|--:|  \n")
	sb.WriteString(format("Sub Total", num.FormatMoney(p.Total)))

	// extra charges
	amtE := extraCharge(d, p, sb, d.ExtraCharge, format)

	// commission
	amtC := num.RoundFloat((p.Total/100)*commission, 2)
	sb.WriteString(format(fmt.Sprintf("Commission(%d%%)", commission), num.FormatMoney(amtC)))

	// estimate feed
	sb.WriteString(format("Estimate Fee", num.FormatMoney(estimateFee)))

	// grand total
	p.Total = p.Total + amtE + amtC + estimateFee

	// minimum order threshold check
	if p.Total < laborSurchargeThreshold {
		sb.WriteString(format("Minimum Labor Surcharge", num.FormatMoney(laborSurcharge)))
		p.Total += laborSurcharge
	}

	sb.WriteString(fmt.Sprintf("| %s | <mark>%s</mark> |  \n", "Grand Total", num.FormatMoney(p.Total)))
	p.Summary = sb.String()
}

// extraCharge applied to the given input
// returns the total amount of extra charge
func extraCharge(d *Data, p *Summary, sb *strings.Builder, ec *ExtraCharge, format func(a, b string) string) float64 {
	amtE := 0.0
	if ec == nil {
		return amtE
	}

	switch ec.Type {
	case enum.ExtraChargeAmount:
		amtE = ec.Val
		sb.WriteString(format("Additional", num.FormatMoney(amtE)))
	case enum.ExtraChargePercent:
		ep := int8(ec.Val)
		amtE = num.RoundFloat((p.Total/100)*float64(ep), 2)
		sb.WriteString(format(fmt.Sprintf("Additional(%d%%)", ep), num.FormatMoney(amtE)))
	case enum.ExtraChargePerSQ:
		charges := ec.Val
		amtE = num.RoundFloat(p.SQ*charges, 2)
		sb.WriteString(format(fmt.Sprintf("Additional(%s * %.2fSQ)", num.FormatMoney(charges), p.SQ), num.FormatMoney(amtE)))
	case enum.ExtraChargeConditional:
		// return if there is no condition
		if len(ec.Conditions) == 0 {
			return amtE
		}

		// loop through each condition
		for _, c := range ec.Conditions {
			switch c.Condition {
			case enum.ConditionAddIfStateIn:
				// expected state value be in []string
				states := c.ConditionVal.([]string)
				if slice.HasFold(d.State, states) {
					amtE = c.Val
					sb.WriteString(format(fmt.Sprintf("Add for state %s", d.State), num.FormatMoney(amtE)))
				}

			case enum.ConditionSubtractIfStateIn:
				// expected state value be in []string
				states := c.ConditionVal.([]string)
				if slice.HasFold(d.State, states) {
					amtE = c.Val // must be a negative value
					// let's make sure it's negative
					if amtE > 0 {
						amtE = -amtE
					}
					sb.WriteString(format(fmt.Sprintf("Subtract for state %s", d.State), num.FormatMoney(amtE)))
				}
			}
		}

	}

	return amtE
}

// steepSlopeMaterialPricing applies for pitch >= 3/12
func steepSlopeMaterialPricing(
	jd *Data, prices []*ProductPrice, summary *Summary, cm, nm Material, sq float64, isPartial bool, sb *strings.Builder,
) error {
	if e, err := materialPricing(cm, nm, sq, isPartial, prices); err != nil {
		return err
	} else {
		summary.Total += e.Amount
		summary.Items = append(summary.Items, e)
		reportEntry(sb, e)
	}

	// Layers = 2 +
	if jd.Layers > 1 {
		if e, err := layerPrice("Layers 2", jd.Layer2Material, prices, sq); err != nil {
			return err
		} else {
			summary.Total += e.Amount
			summary.Items = append(summary.Items, e)
			reportEntry(sb, e)
		}
	}

	// Layers = 3 + add selection for layer 2 +
	if jd.Layers > 2 {
		if e, err := layerPrice("Layers 3", jd.Layer3Material, prices, sq); err != nil {
			return err
		} else {
			summary.Total += e.Amount
			summary.Items = append(summary.Items, e)
			reportEntry(sb, e)
		}
	}

	return nil
}

func repaperPricing(
	cm material.CurrentSteepSlope, prices []*ProductPrice, summary *Summary, sq float64, sb *strings.Builder,
) error {
	if cm != material.ConcreteTile && cm != material.ClayTile {
		return errors.New(fmt.Sprintf("'%s' -> 'Repaper' is not defiled", cm))
	}

	p := UnknownProduct
	switch cm {
	case material.ConcreteTile:
		p = RRConcreteTile
	case material.ClayTile:
		p = RRClayTile
	}

	pp := productPrice(p, prices)
	amt := pp.Price * m50 * sq
	e := &Entry{
		Existing: string(cm),
		New:      "Repaper",
		Logic:    fmt.Sprintf("%s * %.2f * #sq", p, m50),
		LogicVal: fmt.Sprintf("%s * %.2f * %.2f = %s", num.RoundAndFormatFloat(pp.Price), m50, sq, num.FormatMoney(amt)),
		Amount:   amt,
	}

	summary.Total += e.Amount
	summary.Items = append(summary.Items, e)
	reportEntry(sb, e)

	return nil
}

// lowSlopeMaterialPricing applies for pitch < 3/12
// partial percentage will not apply to low slope
func lowSlopeMaterialPricing(prices []*ProductPrice, nm Material, sq float64) (*Entry, error) {
	// check New Low Slope Material
	if nm != ModBit && nm != Coating {
		return nil, errors.New(fmt.Sprintf("'New Low Slope Material' %s expected to be %s or %s", nm, ModBit, Coating))
	}

	// partial % will not be applied to low slope
	return materialPricing(LowSlope, nm, sq, false, prices)
}

func materialPricing(current Material, new Material, sq float64, isPartial bool, prices []*ProductPrice) (*Entry, error) {
	var pricing *ProductPrice
	var mul, add float64
	product := UnknownProduct
	for _, l := range materialLogic {
		if l.existing == current && l.new == new {
			product = l.product
			pricing = productPrice(l.product, prices)
			mul = l.multiplier
			// if there is partial and material logic has extra value for it then use that
			if isPartial && l.extraWithPartial != 0 {
				add = l.extraWithPartial
			} else {
				add = l.extra
			}
			break
		}
	}

	if product == UnknownProduct {
		return nil, errors.New(fmt.Sprintf("'%s' -> '%s' product maping not found", current, new))
	}

	if pricing == nil {
		return nil, errors.New(fmt.Sprintf("'%s' -> '%s' pricing not found", current, new))
	}

	var op, opVal string
	addT := add * sq
	switch {
	case add < 0:
		op = fmt.Sprintf("- %s/sq", num.FormatMoney(math.Abs(add)))
		opVal = fmt.Sprintf("- %s", num.RoundAndFormatFloat(math.Abs(addT)))
	case add > 0:
		op = fmt.Sprintf("+ %s/sq", num.FormatMoney(math.Abs(add)))
		opVal = fmt.Sprintf("+ %s", num.RoundAndFormatFloat(math.Abs(addT)))
	default:
		op, opVal = "", ""
	}

	l := fmt.Sprintf("%s * %.2f * #sq", product, mul)
	lv := fmt.Sprintf("%s * %.2f * %.2f", num.RoundAndFormatFloat(pricing.Price), mul, sq)
	amt := pricing.Price * mul * sq
	amt = amt + (addT)
	return &Entry{
		Existing: string(current),
		New:      string(new),
		Logic:    fmt.Sprintf("%s %s", l, op),
		LogicVal: fmt.Sprintf("%s %s = %s", lv, opVal, num.FormatMoney(amt)),
		Amount:   amt,
	}, nil
}

// steepCharges will apply if pitch value >= 7/12.
func steepCharges(sq, pitch float64, prices []*ProductPrice, summary *Summary, sb *strings.Builder) {
	var product Product
	var existing string
	mul := float64(1)
	switch {
	case pitch >= 13:
		existing = "Pitch = 13/12"
		product = Steep10by12

	case pitch >= 12:
		existing = "Pitch = 12/12"
		product = Steep10by12
		mul = m90

	case pitch >= 11:
		existing = "Pitch = 11/12"
		product = Steep10by12
		mul = m85

	case pitch >= 10:
		existing = "Pitch = 10/12"
		product = Steep10by12
		mul = m80

	case pitch >= 9:
		existing = "Pitch = 9/12"
		product = Steep7by12
		mul = m90
	case pitch >= 8:
		existing = "Pitch = 8/12"
		product = Steep7by12
		mul = m85

	case pitch >= 7:
		existing = "Pitch = 7/12"
		product = Steep7by12
		mul = m80
	}

	p := productPrice(product, prices)
	amt := p.Price * mul * sq
	e := &Entry{
		Existing: existing,
		New:      "#sq *",
		Logic:    fmt.Sprintf("#sq * %s * %.2f", product, mul),
		LogicVal: fmt.Sprintf("%.2f * %s * %.2f = %s", sq, num.RoundAndFormatFloat(p.Price), mul, num.FormatMoney(p.Price*mul*sq)),
		Amount:   amt,
	}

	summary.Total += e.Amount
	summary.Items = append(summary.Items, e)
	reportEntry(sb, e)
}

func layerPrice(layer, material string, prices []*ProductPrice, sq float64) (*Entry, error) {
	var l, lv string
	var amt float64
	nm := layerMaterial(material)
	if nm == UnknownMaterial {
		l = "Add $50.00 * #sq"
		lv = fmt.Sprintf("50.00 * %.2f", sq)
		amt = 50 * sq
	} else {
		var prod Product
		var pricing *ProductPrice
		var mul float64
		for _, l := range layerLogic {
			if l.new == nm {
				prod = l.product
				pricing = productPrice(l.product, prices)
				mul = l.multiplier
				break
			}
		}

		if pricing == nil {
			return nil, errors.New(fmt.Sprintf("'%s' -> '%s' pricing not found", layer, prod))
		}

		l = fmt.Sprintf("%s * %.2f * #sq", prod, mul)
		lv = fmt.Sprintf("%s * %.2f * %.2f", num.RoundAndFormatFloat(pricing.Price), mul, sq)
		amt = pricing.Price * mul * sq
	}

	return &Entry{
		Existing: layer,
		New:      string(nm),
		Logic:    l,
		LogicVal: fmt.Sprintf("%s = %s", lv, num.FormatMoney(amt)),
		Amount:   amt,
	}, nil
}

func redeckPrice(prices []*ProductPrice, summary *Summary, sq float64, sb *strings.Builder) {
	p := productPrice(Sheathing1by2, prices)
	amt := p.Price * m90 * sq
	e := &Entry{
		Existing: "Redeck",
		New:      "Yes",
		Logic:    fmt.Sprintf("Add %s * %.2f * #sq", Sheathing1by2, m90),
		LogicVal: fmt.Sprintf("%s * %.2f * %.2f = %s", num.RoundAndFormatFloat(p.Price), m90, sq, num.FormatMoney(amt)),
		Amount:   amt,
	}
	summary.Total += e.Amount
	summary.Items = append(summary.Items, e)
	reportEntry(sb, e)
}

func layerMaterial(val string) Material {
	m := CurrentMaterial(val)
	switch m {
	case Asphalt, Woodshake:
		return m
	default:
		return UnknownMaterial
	}
}

func productPrice(p Product, prices []*ProductPrice) *ProductPrice {
	for _, v := range prices {
		if v.ProductID == p.UInt8() {
			return v
		}
	}

	return nil
}

func note(str string) string {
	return fmt.Sprintf("🔍 %s  \n", str)
}

func reportEntry(sb *strings.Builder, e *Entry) {
	sb.WriteString(fmt.Sprintf("##### %s → %s  \n", e.Existing, e.New))
	sb.WriteString(fmt.Sprintf("\t%s  \n", e.Logic))
	sb.WriteString(fmt.Sprintf("\t%s  \n", e.LogicVal))
}

func wasteFactor(p uint) float64 {
	return (float64(p) / 100) + 1
}
