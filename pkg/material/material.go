package material

const (
	SteepSlopeMinPitch = 2.0
	// WasteFactorDefault 10%
	WasteFactorDefault = uint(10)
	// WasteFactorSteep 15%
	WasteFactorSteep = uint(15) //

	Shingles               = "Asphalt Shingles"
	ThreeTabShingles       = "3-Tab Asphalt Shingles"
	Architectural          = "Architectural Asphalt Shingles"
	BestValue              = "Best Value - Architectural"
	MoreExpensive          = "More Expensive - Premium Architectural"
	ModBit                 = "Modified Bitumen"
	Coating                = "Coating"
	GravelBur              = "Gravel BUR"
	LowSlopeOnly           = "Low Slope Only"
	MetalShakes            = "Metal Shakes"
	MetalTitle             = "Metal Tile"
	MetalRPanelExpFastener = "Metal R-Panel Exposed Fastener"
	RolledRoofing          = "Rolled Roofing"
	StandingSeamMetal      = "Standing Seam Metal"
	ClayTile               = "Clay Tile"
	ClayTileBarrel         = "Clay Tile (Barrel)"
	ConcreteTile           = "Concrete Tile"
	BlackMembrane          = "Black Membrane"
	WhiteMembrane          = "White Membrane"
	WoodShakes             = "Wood Shakes"
	// Repaper - Reuse Existing Tile
	Repaper = "Repaper"
	Slate   = "Slate"
)

// CurrentSteepSlope a.k.a Current Material - old Steep Slope Material
type CurrentSteepSlope string

func (e CurrentSteepSlope) String() string {
	return string(e)
}

func (e CurrentSteepSlope) IsValid() bool {
	switch e {
	case "",
		ThreeTabShingles,
		Architectural,
		ClayTile,
		ConcreteTile,
		WoodShakes,
		MetalShakes,
		MetalTitle,
		StandingSeamMetal,
		MetalRPanelExpFastener,
		Slate,
		LowSlopeOnly:
		return true
	}
	return false
}

// NewSteepSlope a.k.a New Roofing Material - Steep Slope Products
type NewSteepSlope string

func (e NewSteepSlope) String() string {
	return string(e)
}

func (e NewSteepSlope) IsValid() bool {
	switch e {
	case "",
		BestValue,
		MoreExpensive,
		StandingSeamMetal,
		ConcreteTile,
		ClayTileBarrel,
		Repaper:
		return true
	}

	return false
}

// CurrentLowSlope a.k.a Current Material - Low Slope
type CurrentLowSlope string

func (e CurrentLowSlope) String() string {
	return string(e)
}

func (e CurrentLowSlope) IsValid() bool {
	switch e {
	case "",
		RolledRoofing,
		Shingles,
		BlackMembrane,
		WhiteMembrane,
		GravelBur:
		return true
	}

	return false
}

// NewLowSlope a.k.a New Low Slope Material - Low Slope Products
type NewLowSlope string

func (e NewLowSlope) String() string {
	return string(e)
}

func (e NewLowSlope) IsValid() bool {
	switch e {
	case "",
		ModBit,
		Coating:
		return true
	}

	return false
}

// Layer a.k.a Current Material - old Steep Slope Material
type Layer string

func (e Layer) String() string {
	return string(e)
}

func (e Layer) IsValid() bool {
	switch e {
	case "",
		ThreeTabShingles,
		Architectural,
		Shingles,
		WoodShakes,
		RolledRoofing,
		BlackMembrane,
		WhiteMembrane,
		LowSlopeOnly,

		// RFX estimate req can send following
		ClayTile,
		ConcreteTile,
		MetalShakes,
		MetalTitle,
		StandingSeamMetal,
		MetalRPanelExpFastener,
		Slate:
		return true
	}
	return false
}
