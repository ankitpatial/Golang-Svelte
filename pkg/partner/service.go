/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package partner

import "roofix/server/model"

type Service string

const (
	NoService                       Service = ""                                     // 0
	AsphaltShingleTearOff           Service = "ASPHALT_SHINGLE_TEAR_OFF"             // 1
	AsphaltShingleOverlay           Service = "ASPHALT_SHINGLE_OVERLAY"              // 2
	AsphaltShingleSmallRepairs      Service = "ASPHALT_SHINGLE_SMALL_REPAIRS"        // 3
	RolledAsphaltTearOff            Service = "ROLLED_ASPHALT_TEAR_OFF"              // 4
	TarGravelRepairs                Service = "TAR_GRAVEL_REPAIRS"                   // 5
	ConcreteTileRemoveReplace       Service = "CONCRETE_TILE_REMOVER_REPLACE"        // 6
	ConcreteTileUnderArrayRoofSwaps Service = "CONCRETE_TILE_UNDER_ARRAY_ROOF_SWAPS" // 7
	ConcreteTileSmallRepairs        Service = "CONCRETE_TILE_SMALL_REPAIRS"          // 8
	SlateTileRemoveReplace          Service = "SLATE_TILE_REMOVE_REPLACE"            // 9
	SlateTileUnderArrayRoofSwaps    Service = "SLATE_TILE_UNDER_ARRAY_ROOF_SWAPS"    // 10
	SlateTileSmallRepairs           Service = "SLATE_TILE_SMALL_REPAIRS"             // 11
	MetalRoofRepairs                Service = "METAL_ROOF_REPAIRS"                   // 12
	FlatRoofRepairs                 Service = "FLAT_ROOF_REPAIRS"                    // 13
	RoofTrussRetrofits              Service = "ROOF_TRUSS_RETROFITS"                 // 14
)

var Services = []*model.Entity{
	{ID: AsphaltShingleTearOff.String(), Name: AsphaltShingleTearOff.Description()},
	{ID: AsphaltShingleOverlay.String(), Name: AsphaltShingleOverlay.Description()},
	{ID: AsphaltShingleSmallRepairs.String(), Name: AsphaltShingleSmallRepairs.Description()},
	{ID: RolledAsphaltTearOff.String(), Name: RolledAsphaltTearOff.Description()},
	{ID: TarGravelRepairs.String(), Name: TarGravelRepairs.Description()},
	{ID: ConcreteTileRemoveReplace.String(), Name: ConcreteTileRemoveReplace.Description()},
	{ID: ConcreteTileUnderArrayRoofSwaps.String(), Name: ConcreteTileUnderArrayRoofSwaps.Description()},
	{ID: ConcreteTileSmallRepairs.String(), Name: ConcreteTileSmallRepairs.Description()},
	{ID: SlateTileRemoveReplace.String(), Name: SlateTileRemoveReplace.Description()},
	{ID: SlateTileUnderArrayRoofSwaps.String(), Name: SlateTileUnderArrayRoofSwaps.Description()},
	{ID: SlateTileSmallRepairs.String(), Name: SlateTileSmallRepairs.Description()},
	{ID: MetalRoofRepairs.String(), Name: MetalRoofRepairs.Description()},
	{ID: FlatRoofRepairs.String(), Name: FlatRoofRepairs.Description()},
	{ID: RoofTrussRetrofits.String(), Name: RoofTrussRetrofits.Description()},
}

var AllServices = []Service{
	NoService,
	AsphaltShingleTearOff,
	AsphaltShingleOverlay,
	AsphaltShingleSmallRepairs,
	RolledAsphaltTearOff,
	TarGravelRepairs,
	ConcreteTileRemoveReplace,
	ConcreteTileUnderArrayRoofSwaps,
	ConcreteTileSmallRepairs,
	SlateTileRemoveReplace,
	SlateTileUnderArrayRoofSwaps,
	SlateTileSmallRepairs,
	MetalRoofRepairs,
	FlatRoofRepairs,
	RoofTrussRetrofits,
}

func (e Service) ToUInt8() uint8 {
	for idx, i := range AllServices {
		if e == i {
			return uint8(idx)
		}
	}
	return 0
}

func (e *Service) FromUInt8(v uint8) {
	for idx, i := range AllServices {
		if uint8(idx) == v {
			*e = i
			break
		}
	}
}

func (e Service) String() string {
	return string(e)
}

func (e Service) Description() string {
	switch e {
	case AsphaltShingleTearOff:
		return "Asphalt Shingle - Tear-off"
	case AsphaltShingleOverlay:
		return "Asphalt Shingle - Overlay"
	case AsphaltShingleSmallRepairs:
		return "Asphalt Shingle - Small Repairs"
	case RolledAsphaltTearOff:
		return "Rolled Asphalt - Tear-off"
	case TarGravelRepairs:
		return "Tar & Gravel - Repairs"
	case ConcreteTileRemoveReplace:
		return "Concrete Tile - Remove & Replace(Remove Tiles - Replace Underlayment - Replace Tiles)"
	case ConcreteTileUnderArrayRoofSwaps:
		return "Concrete Tile - Under Array Roof Swaps"
	case ConcreteTileSmallRepairs:
		return "Concrete Tile - Small Repairs"
	case SlateTileRemoveReplace:
		return "Slate Tile - Remove & Replace(Remove Tiles - Replace Underlayment - Replace Tiles)"
	case SlateTileUnderArrayRoofSwaps:
		return "Slate Tile - Under Array Roof Swaps"
	case SlateTileSmallRepairs:
		return "Slate Tile - Small Repairs"
	case MetalRoofRepairs:
		return "Metal Roof - Repairs"
	case FlatRoofRepairs:
		return "Flat Roof - Repairs"
	case RoofTrussRetrofits:
		return "Roof Truss Retrofits"
	default:
		return ""
	}
}
