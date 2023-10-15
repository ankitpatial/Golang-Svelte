package estimate

import (
	"errors"
	"fmt"
	"roofix/pkg/material"
	"roofix/pkg/util/num"
	"roofix/pkg/util/ptr"
	"strings"
)

type MaterialDetail struct {
	CurrentMaterial            material.CurrentSteepSlope `json:"currentMaterial"`
	NewRoofingMaterial         material.NewSteepSlope     `json:"newRoofingMaterial"`
	CurrentMaterialLowSlope    material.CurrentLowSlope   `json:"currentMaterialLowSlope"`
	NewRoofingMaterialLowSlope material.NewLowSlope       `json:"newRoofingMaterialLowSlope"`
	LowSlope                   bool                       `json:"lowSlope"`
	Redeck                     bool                       `json:"redeck"`
	Layers                     int                        `json:"layers"`
	Layer2Material             material.Layer             `json:"layer2Material"`
	Layer3Material             material.Layer             `json:"layer3Material"`
	PartialPercent             float64                    `json:"partialPercent"`
	OverrideSummary            *string
}

func MaterialDetailCheck(
	cm material.CurrentSteepSlope,
	nm material.NewSteepSlope,
	nmLowSlope material.NewLowSlope,
	lowSlope,
	redeck bool,
	layers int,
	layer2 material.Layer,
	layer3 material.Layer,
	partial *int,
) (*MaterialDetail, error) {
	// job input
	detail := &MaterialDetail{
		CurrentMaterial:            cm,
		NewRoofingMaterial:         nm,
		CurrentMaterialLowSlope:    "",
		NewRoofingMaterialLowSlope: nmLowSlope,
		LowSlope:                   lowSlope,
		Redeck:                     redeck,
		Layers:                     layers,
		PartialPercent:             float64(num.ReadIntPtr(partial)),
	}

	// required,  current roof material
	if detail.CurrentMaterial == "" {
		return nil, errors.New(fmt.Sprintf("'%s' is not a valid 'Current Roof Material'", detail.CurrentMaterial))
	}
	// validate, current roof material
	if !detail.CurrentMaterial.IsValid() {
		return nil, errors.New(fmt.Sprintf("'%s' is not a valid 'Current Roof Material'", detail.CurrentMaterial))
	}

	var sb strings.Builder
	// material, Low Slope Only
	if detail.CurrentMaterial == material.LowSlopeOnly { // low slope only
		// validate Low Slope new roof material
		if !detail.CurrentMaterialLowSlope.IsValid() {
			return nil, errors.New(fmt.Sprintf("%s is not a valid low slope 'Current Roof Material'", detail.CurrentMaterialLowSlope))
		}

		// validate Low Slope new roof material
		if !detail.NewRoofingMaterialLowSlope.IsValid() {
			return nil, errors.New(fmt.Sprintf("%s is not a valid low slope 'New Roof Material'", detail.NewRoofingMaterialLowSlope))
		}

		if detail.Layers > 1 {
			detail.Layers = 1
			sb.WriteString(fmt.Sprintf("- overriding 'Layers' value %d → 1 (current material is Low Slope)  \n", detail.Layers))
		}

	} else {
		// validate, New Roof Material
		if !detail.NewRoofingMaterial.IsValid() {
			return nil, errors.New(fmt.Sprintf("%s is not a valid 'New Roof Material'", detail.NewRoofingMaterial))
		}

		// lowSlope = true & now material specified, set default Mod Bit
		if lowSlope && nmLowSlope == "" {
			detail.NewRoofingMaterialLowSlope = material.ModBit
			sb.WriteString(fmt.Sprintf("- overriding 'New Low Slope Material' → %s (no option was provided)", material.ModBit))
		}

		// re-decking is required with existing material = Wood Shakes
		if detail.CurrentMaterial == material.WoodShakes && !detail.Redeck {
			detail.Redeck = true
			sb.WriteString("- overriding 'Redeck' value → Yes (required with 'Wood Shakes')  \n")
		}

		if layers > 1 {
			if layer2 == "" {
				detail.Layer2Material = material.Layer(detail.CurrentMaterial)
				sb.WriteString(fmt.Sprintf("- overriding 'Layer 2 Material' value → %s (no option was provided)  \n", detail.Layer2Material))
			} else {
				detail.Layer2Material = layer2
			}
		}

		if layers > 2 {
			if layer3 == "" {
				detail.Layer3Material = material.Layer(detail.CurrentMaterial)
				sb.WriteString(fmt.Sprintf("- overriding 'Layer 3 Material' value → %s (no option was provided)  \n", detail.Layer3Material))
			} else {
				detail.Layer3Material = layer2
			}
		}
	}

	// set partial to 0 if too low or too high
	if detail.PartialPercent < 0 || detail.PartialPercent > 99 {
		sb.WriteString(fmt.Sprintf("- overriding 'Partial Roof Replacement' value %.0f %% → 0 (invalid range must be 0-99)  \n", detail.PartialPercent))
		detail.PartialPercent = 0
	}

	if sb.Len() > 0 {
		detail.OverrideSummary = ptr.Str(fmt.Sprintf("Value Override Summary:  \n" + sb.String()))
	}

	return detail, nil
}
