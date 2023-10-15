/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package nearmap

import (
	"context"
	"roofix/ent"
	"roofix/ent/schema"
	"roofix/pkg/enum"
	"roofix/pkg/material"
	"roofix/pkg/pricing"
	"roofix/pkg/util"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"time"
)

const parcelCountry = "US"

type Data struct {
	StreetAddress              string
	City                       string
	State                      string
	Zip                        string
	Latitude                   float64
	Longitude                  float64
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
	EstimateType               enum.Measure
	ExtraCharge                *pricing.ExtraCharge
}

type EstimateInfo struct {
	Response *Response
	Price    *Price
	Rollups  map[string]*Rollup
	Boundary []schema.Point
}

type EstimateDetail struct {
	ID                                string      `json:"-"`
	RespID                            string      `json:"-"`
	ExternalID                        string      `json:"externalID"`
	Price                             float64     `json:"contractPrice"`
	PrimaryPitchInDegrees             float64     `json:"primaryPitchInDegrees"`
	PrimaryPitch                      float64     `json:"primaryPitch"`
	TileArea                          float64     `json:"tileArea"`
	TileRatio                         float64     `json:"tileRatio"`
	ShingleArea                       float64     `json:"shingleArea"`
	ShingleRatio                      float64     `json:"shingleRatio"`
	MetalArea                         float64     `json:"metalArea"`
	MetalRatio                        float64     `json:"metalRatio"`
	FlatArea                          float64     `json:"flatArea"`
	FlatRatio                         float64     `json:"flatRatio"`
	GableArea                         float64     `json:"gableArea"`
	GableRatio                        float64     `json:"gableRatio"`
	HipArea                           float64     `json:"hipArea"`
	HipRatio                          float64     `json:"hipRatio"`
	DutchGableArea                    float64     `json:"dutchGableArea"`
	DutchGableRatio                   float64     `json:"dutchGableRatio"`
	TurretArea                        float64     `json:"turretArea"`
	TurretRatio                       float64     `json:"turretRatio"`
	DominantRoofMaterial              string      `json:"dominantRoofMaterial"`
	DominantRoofMaterialID            int         `json:"dominantRoofMaterialID"`
	RoofCount                         int         `json:"roofCount"`
	TotalUnclippedArea                float64     `json:"totalUnclippedArea"`
	TotalSquares                      float64     `json:"totalSquares"`
	RoofMaterialRatioTotal            float64     `json:"roofMaterialRatioTotal"`
	RoofTypeRatioTotal                float64     `json:"roofTypeRatioTotal"`
	RoofMaterialSurfaceAreaTotal      float64     `json:"roofMaterialSurfaceAreaTotal"`
	RoofTypeSurfaceAreaTotal          float64     `json:"roofTypeSurfaceAreaTotal"`
	TreeOverhangRatioPrimaryRoof      float64     `json:"treeOverhangRatioPrimaryRoof"`
	TreeOverhangConfidencePrimaryRoof float64     `json:"treeOverhangConfidencePrimaryRoof"`
	TreeOverhangPresenceConfidence    float64     `json:"treeOverhangPresenceConfidence"`
	TreeOverhangAreaPrimaryRoof       float64     `json:"treeOverhangAreaPrimaryRoof"`
	TreeOverhangTotalClippedArea      float64     `json:"treeOverhangTotalClippedArea"`
	TreeOverhangTotalUnClippedArea    float64     `json:"treeOverhangTotalUnClippedArea"`
	TreeOverhangPresent               bool        `json:"treeOverhangPresent"`
	TreeOverhangCount                 int         `json:"treeOverhangCount"`
	SteepWasteFactor                  uint        `json:"-"`
	FlatWasteFactor                   uint        `json:"-"`
	RoofArea                          []*RoofArea `json:"-"`
	FailureStatus                     string      `json:"failureStatus,omitempty"`
	Note                              *string     `json:"note,omitempty"`
}

func Estimate(ctx context.Context, data *Data) (*EstimateInfo, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	//
	// There is problem with tree-overhang and the pitch degree coming null.
	// Solution to above issue is attempt past surveys. But there is chance of error in doing so.
	// So, let's attempt past survey on any error use first survey response
	//
	var res *Response
	var firstRes *Response
	var geo *Geometry
	var err error

	until := time.Now().Format(time.DateOnly)
	count := 1
	for { // max 5 attempts to do
		if count > 5 {
			break
		}

		// get parcel rollups by its address
		log.Info("=> attempt %d", count)
		res, err = addressRollups(ctx, data.StreetAddress, data.City, data.State, data.Zip, parcelCountry, until)
		if err != nil {
			// error on rollups
			log.Warn("attempt %d failed with error: %s", count, err.Error())

			// nearmap rollups search can return errors:
			//  404 Not Found
			//  500 Internal Server Error
			// so if its 1st attempt then let's try with regrid
			if count == 1 {
				log.Info("trying with geometry coordinates")
				// get parcel geometry coordinates
				geo, err = parcelGeometry(ctx, data.Latitude, data.Longitude)
				if err != nil {
					err = log.Wrap(err, "failed on parcelGeometry")
					break
				}

				// try getting rollups with parcel geometry
				res, err = geometryRollups(ctx, geo)
				if err != nil {
					err = log.Wrap(err, "failed on geometryRollups")
					break
				}

			} else if count > 1 && firstRes != nil {
				// check if 2nd onwards attempt
				log.Info("=> attempt %d failed, using response from 1st attempt", count)
				// let's use the first survey response
				res = firstRes
				err = nil
			}

			break
		} else if firstRes == nil {
			// set first survey response
			firstRes = res
		}

		overhang := readBool(TreeOverhangPresent, res.Rollups)
		pitchNil := isNil(PitchDegreePrimaryRoof, res.Rollups)
		log.Info("=> survey date: %s, tree overhang: %v, pitch nil: %v", res.SurveyDate, overhang, pitchNil)
		if overhang || pitchNil {
			var t time.Time
			// read current survey date
			if t, err = time.Parse(time.DateOnly, res.SurveyDate); err != nil {
				err = log.Wrap(err, "failed on time.Parse(res.SurveyDate)")
				break
			}

			// go or earlier survey, let's go a month ago survey date(it will be enough to fetch the previous survey)
			t = util.LastMonth(t)
			until = t.Format(time.DateOnly)
			count++
			continue
		}

		break
	}

	if err != nil {
		return nil, log.Wrap(err, "nearmap api error")
	}

	// ✦ price calc
	log.Info("calculating price")
	p, err := getPrice(ctx, data.EstimateType, pricingData(data), res.Rollups)
	if err != nil {
		return nil, log.Wrap(err, "price calculation error")
	}

	// parcel bounds
	var bounds []schema.Point
	if geo != nil {
		bounds = parcelGeometryPoints(geo)
	} else {
		bounds = rectangle(
			schema.Point{
				Lat: res.ExtentMinY,
				Lng: res.ExtentMinX,
			},
			schema.Point{
				Lat: res.ExtentMaxY,
				Lng: res.ExtentMaxX,
			})
	}

	// return
	return &EstimateInfo{
		Response: res,
		Price:    p,
		Rollups:  res.Rollups,
		Boundary: bounds,
	}, nil
}

func pricingData(j *Data) *pricing.Data {
	return &pricing.Data{
		State:                      j.State,
		Zip:                        j.Zip,
		Lat:                        j.Latitude,
		Lng:                        j.Longitude,
		CurrentMaterial:            j.CurrentMaterial,
		NewRoofingMaterial:         j.NewRoofingMaterial,
		LowSlope:                   j.LowSlope,
		CurrentMaterialLowSlope:    j.CurrentMaterialLowSlope,
		NewRoofingMaterialLowSlope: j.NewRoofingMaterialLowSlope,
		Redeck:                     j.Redeck,
		Layers:                     j.Layers,
		Layer2Material:             j.Layer2Material,
		Layer3Material:             j.Layer3Material,
		PartialPercentage:          j.PartialPercentage,
		ExtraCharge:                j.ExtraCharge,
	}
}

func FillEstimateDetail(rollup map[string]*Rollup, primaryRoof bool) *EstimateDetail {

	pd := readFloat(PitchDegreePrimaryRoof, rollup) // pitch degrees of primary roof
	pitch := roofingPitch(pd)

	data := new(EstimateDetail)
	data.RoofCount = int(readFloat(roofCount, rollup))
	data.PrimaryPitchInDegrees = pd
	data.PrimaryPitch = pitch

	data.FlatArea = num.RoundFloat(readFloat(flat, rollup), precision)
	data.FlatRatio = num.RoundFloat(readFloat(flatRatio, rollup), precision)

	// material primary roof area
	data.ShingleArea = totalSqFt(pitch, num.RoundFloat(readFloat(shingle, rollup), precision))
	data.TileArea = totalSqFt(pitch, num.RoundFloat(readFloat(tile, rollup), precision))
	data.MetalArea = totalSqFt(pitch, num.RoundFloat(readFloat(metal, rollup), precision))

	// roof type primary roof area
	data.GableArea = totalSqFt(pitch, num.RoundFloat(readFloat(gable, rollup), precision))
	data.DutchGableArea = totalSqFt(pitch, num.RoundFloat(readFloat(dutch, rollup), precision))
	data.HipArea = totalSqFt(pitch, num.RoundFloat(readFloat(hip, rollup), precision))
	data.TurretArea = totalSqFt(pitch, num.RoundFloat(readFloat(turret, rollup), precision))

	// in case measurement is called for all structures
	if !primaryRoof {
		// use 'total unclipped area' if its grater than primary area

		// material total area =>
		st := totalSqFt(pitch, num.RoundFloat(readFloat(shingleT, rollup), precision))
		if st > data.ShingleArea {
			data.ShingleArea = st
		}
		tt := totalSqFt(pitch, num.RoundFloat(readFloat(tileT, rollup), precision))
		if tt > data.TileArea {
			data.TileArea = tt
		}
		mt := totalSqFt(pitch, num.RoundFloat(readFloat(metalT, rollup), precision))
		if mt > data.MetalArea {
			data.MetalArea = mt
		}

		// roof type total area =>
		gat := totalSqFt(pitch, num.RoundFloat(readFloat(gableT, rollup), precision))
		if gat > data.GableArea {
			data.GableArea = gat
		}
		dga := totalSqFt(pitch, num.RoundFloat(readFloat(dutchT, rollup), precision))
		if dga > data.DutchGableArea {
			data.DutchGableArea = dga
		}
		ha := totalSqFt(pitch, num.RoundFloat(readFloat(hipT, rollup), precision))
		if ha > data.HipArea {
			data.HipArea = ha
		}
		ta := totalSqFt(pitch, num.RoundFloat(readFloat(turretT, rollup), precision))
		if ta > data.TurretArea {
			data.TurretArea = ta
		}
	}

	// material dominance
	if readBool(shingleDominant, rollup) {
		data.DominantRoofMaterial = "shingle"
		data.DominantRoofMaterialID = 1
	}
	if readBool(tileDominant, rollup) {
		data.DominantRoofMaterial = "tile"
		data.DominantRoofMaterialID = 3
	}
	if readBool(metalDominant, rollup) {
		data.DominantRoofMaterial = "metal"
		data.DominantRoofMaterialID = 2
	}

	// material ratio
	data.ShingleRatio = num.RoundFloat(readFloat(shingleRatio, rollup), precision)
	data.TileRatio = num.RoundFloat(readFloat(tileRatio, rollup), precision)
	data.MetalRatio = num.RoundFloat(readFloat(metalRatio, rollup), precision)

	data.RoofMaterialRatioTotal = data.ShingleRatio + data.MetalRatio + data.TileRatio
	data.RoofMaterialSurfaceAreaTotal = data.ShingleArea + data.MetalArea + data.TileArea

	// roof type ratio
	data.GableRatio = num.RoundFloat(readFloat(gableRatio, rollup), precision)
	data.DutchGableRatio = num.RoundFloat(readFloat(dutchRatio, rollup), precision)
	data.HipRatio = num.RoundFloat(readFloat(hipRatio, rollup), precision)
	data.TurretRatio = num.RoundFloat(readFloat(turretRatio, rollup), precision)

	data.RoofTypeRatioTotal = data.FlatRatio + data.GableRatio + data.HipRatio + data.DutchGableRatio + data.TurretRatio
	data.RoofTypeSurfaceAreaTotal = data.FlatArea + data.GableArea + data.HipArea + data.DutchGableArea + data.TurretArea

	// tree overhang
	data.TreeOverhangAreaPrimaryRoof = num.RoundFloat(readFloat(treeOH, rollup), precision)
	data.TreeOverhangRatioPrimaryRoof = readFloat(treeOHRatio, rollup)
	data.TreeOverhangConfidencePrimaryRoof = readFloat(treeOHConfidence, rollup)
	data.TreeOverhangPresenceConfidence = readFloat(treeOHPConfidence, rollup)
	data.TreeOverhangTotalClippedArea = num.RoundFloat(readFloat(treeOHClipped, rollup), precision)
	data.TreeOverhangTotalUnClippedArea = num.RoundFloat(readFloat(treeOHUnClipped, rollup), precision)
	data.TreeOverhangPresent = readBool(treeOHPresent, rollup)
	data.TreeOverhangCount = int(readFloat(treeOHCount, rollup))

	// total unclipped area
	data.TotalUnclippedArea = num.RoundFloat(readFloat(unclippedArea, rollup), precision)

	// waste factor
	if data.HipRatio+data.DutchGableRatio+data.TurretRatio > .5 {
		// When (Hip Ratio + Dutch Gable Ratio + Turret Ratio) > .5
		// add 15% waste to anything that is not flat.
		data.SteepWasteFactor = material.WasteFactorSteep
	} else {
		// 10% waste needs to be added by default.
		data.SteepWasteFactor = material.WasteFactorDefault
	}
	// By default, always add a 10% waste to flat.
	data.FlatWasteFactor = material.WasteFactorDefault

	//
	// ROOF area, Steep Slope & Flat
	//
	var steepArea float64
	var des string
	materialSqFt := data.ShingleArea + data.TileArea + data.MetalArea
	typeSqFt := data.GableArea + data.DutchGableArea + data.HipArea + data.TurretArea

	// use max in material & roof type
	if materialSqFt > typeSqFt {
		steepArea = materialSqFt
		des = "primary roof area"
	} else {
		steepArea = typeSqFt
		des = "primary roof area(roof types)"
	}

	if steepArea > 0 {
		data.RoofArea = append(data.RoofArea, &RoofArea{
			SqFt:        steepArea,
			Description: des,
		})
	}

	if data.FlatArea > 0 {
		data.RoofArea = append(data.RoofArea, &RoofArea{
			SqFt:        data.FlatArea,
			Description: "primary roof flat area",
			IsFlat:      true,
		})
	}

	return data
}

func isNil(key string, data map[string]*Rollup) bool {
	if v, ok := data[key]; ok && v.Value != nil {
		return false
	}

	return true
}

func readFloat(key string, data map[string]*Rollup) float64 {
	if v, ok := data[key]; ok && v.Value != nil {
		return v.Value.(float64)
	}

	return 0
}

func readBool(key string, data map[string]*Rollup) bool {
	if v, ok := data[key]; ok && v.Value != nil {
		return v.Value.(bool)
	}

	return false
}
