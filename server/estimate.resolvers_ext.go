package server

import (
	"roofix/pkg/material"
	"roofix/server/model"
)

func estimateCurrentMaterial(v model.CurrentMaterial) material.CurrentSteepSlope {
	switch v {
	case model.CurrentMaterialThreeTabAsphaltShingles:
		return material.ThreeTabShingles
	case model.CurrentMaterialArchAsphaltShingles:
		return material.Architectural
	case model.CurrentMaterialTileClay:
		return material.ClayTile
	case model.CurrentMaterialTileConcrete:
		return material.ConcreteTile
	case model.CurrentMaterialWoodShakes:
		return material.WoodShakes
	case model.CurrentMaterialMetalShakes:
		return material.MetalShakes
	case model.CurrentMaterialMetalTitle:
		return material.MetalTitle
	case model.CurrentMaterialStandingSeamMetal:
		return material.StandingSeamMetal
	case model.CurrentMaterialSlate:
		return material.Slate
	case model.CurrentMaterialMetalRPanelExpFastener:
		return material.MetalRPanelExpFastener
	case model.CurrentMaterialLowSlopeOnly:
		return material.LowSlopeOnly
	default:
		return ""
	}
}

func estimateNewMaterial(v model.NewMaterial) material.NewSteepSlope {
	switch v {
	case model.NewMaterialArchBestValue:
		return material.BestValue
	case model.NewMaterialArchMoreExpensive:
		return material.MoreExpensive
	case model.NewMaterialStandingSeamMetal:
		return material.StandingSeamMetal
	case model.NewMaterialTileConcrete:
		return material.ConcreteTile
	case model.NewMaterialTileClayBarrel:
		return material.ClayTileBarrel
	case model.NewMaterialRepaper:
		return material.Repaper
	case model.NewMaterialModBit:
		return material.ModBit
	default:
		return ""
	}
}
