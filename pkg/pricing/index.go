/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package pricing

import "roofix/pkg/material"

type Product string

type Material string

const (
	UnknownMaterial Material = "Unknown"

	Asphalt   Material = "Asphalt"
	Metal     Material = "Metal"
	Tile      Material = "Tile"
	Woodshake Material = "Woodshake"
	LowSlope  Material = "Low Slope"

	// BestValueArch a.k.a "Best Value BestValueArch"
	BestValueArch     Material = "Best Value Architectural"
	MoreExpensiveArch Material = "More Expensive Premium Architectural"
	StandingSeam      Material = "Standing Seam"
	ConcreteTile      Material = "Concrete Tile"
	ClayTile          Material = "Clay Tile"
	ModBit            Material = "Mod Bit"
	Coating           Material = "Coating"
	Repaper           Material = "Repaper"

	UnknownProduct   Product = "Unknown"
	RRLaminateAR     Product = "R & R Laminate AR"
	RRStandingSeam   Product = "R & R Standing Seam"
	RRConcreteTile   Product = "R & R Concrete Tile"
	RRClayTile       Product = "R & R Clay Tile(Barrel)"
	RRModBit         Product = "R & R Mod Bit"
	CoatingProduct   Product = "Coating"
	WoodShakeTearOff Product = "Wood Shake:Tear off, Haul, and Dispose"
	AddLayerFelt     Product = "Add. Layer Felt:Tear off, Haul, and Dispose"
	LayerOfShingles  Product = "Layer of Shingles & Felt:Tear off, Haul, and Dispose"
	LayerOfMetal     Product = "Layer of Metal: No haul off"
	Sheathing1by2    Product = "Sheathing 1/2\" OSB+Clips"
	Steep7by12       Product = "STEEP 7/12-9/12"
	Steep10by12      Product = "STEEP 10/12-12/12"
	TwoStory         Product = "TWO STORY"
	BasicPerSqAvg    Product = "BASIC PER SQ AVE"
	DisposalFees     Product = "Additional Disposal Fees - 30YD Dumpster"
	PermitFees       Product = "Additional Permit Fees"
)

func (m Product) UInt8() uint8 {
	return uint8(m.Int())
}

func (m Product) Int() int {
	switch m {
	case RRLaminateAR:
		return 1
	case RRStandingSeam:
		return 2
	case RRConcreteTile:
		return 3
	case RRClayTile:
		return 4
	case RRModBit:
		return 5
	case CoatingProduct:
		return 6
	case WoodShakeTearOff:
		return 7
	case AddLayerFelt:
		return 8
	case LayerOfShingles:
		return 9
	case LayerOfMetal:
		return 10
	case Sheathing1by2:
		return 11
	case Steep7by12:
		return 12
	case Steep10by12:
		return 13
	case TwoStory:
		return 14
	case BasicPerSqAvg:
		return 15
	case DisposalFees:
		return 16
	case PermitFees:
		return 17
	default:
		return 0
	}
}

func (m Product) IsValid() bool {
	switch m {
	case RRLaminateAR, RRStandingSeam, RRConcreteTile, RRClayTile, RRModBit, CoatingProduct, WoodShakeTearOff, AddLayerFelt,
		LayerOfShingles, LayerOfMetal, Sheathing1by2, Steep7by12, Steep10by12, TwoStory, BasicPerSqAvg,
		DisposalFees, PermitFees:
		return true
	default:
		return false
	}
}

type ProductInfo struct {
	ID   int
	Name string
}

type Input struct {
	PostalCountry string  `json:"postalCountry" validate:"required"`
	PostalCode    string  `json:"postalCode"  validate:"required"`
	Product       Product `json:"material"  validate:"valid"`
	Description   *string `json:"description"`
	Price         float64 `json:"price"`
}

type ProductPrice struct {
	ID        string  `json:"id"`
	Country   string  `json:"postal_country"`
	State     string  `json:"state"`
	StateAbr  string  `json:"state_abr"`
	Zip       string  `json:"postal_code"`
	City      string  `json:"city"`
	ProductID uint8   `json:"product_id"`
	Price     float64 `json:"price"`
	PricePer  string  `json:"price_per"`
}

func CurrentMaterial(val string) Material {
	switch val {
	case material.ThreeTabShingles, material.Architectural, material.Shingles:
		return Asphalt
	case material.ConcreteTile, material.ClayTile:
		return Tile
	case material.MetalShakes, material.MetalTitle, material.StandingSeamMetal, material.MetalRPanelExpFastener:
		return Metal
	case material.WoodShakes:
		return Woodshake
	case material.LowSlopeOnly:
		return LowSlope
	default:
		return UnknownMaterial
	}
}

func NewMaterial(val string) Material {
	switch val {
	case material.BestValue:
		return BestValueArch
	case material.MoreExpensive:
		return MoreExpensiveArch
	case material.StandingSeamMetal:
		return StandingSeam
	case material.ConcreteTile:
		return ConcreteTile
	case material.ClayTile, material.ClayTileBarrel:
		return ClayTile
	case material.ModBit:
		return ModBit
	case material.Coating:
		return Coating
	case material.Repaper:
		return Repaper
	default:
		return UnknownMaterial
	}
}
