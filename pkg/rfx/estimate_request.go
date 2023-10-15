package rfx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/estimate"
	"roofix/pkg/homeOwner"
	"roofix/pkg/material"
	"roofix/pkg/nearmap"
	pkgPartner "roofix/pkg/partner"
	price "roofix/pkg/pricing"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/storage"
	"roofix/pkg/util/str"
	"roofix/pkg/util/uid"
	"roofix/server/router/response"
)

type EstimateInput struct {
	ExternalID                 *string           `json:"externalID"`
	CompanyID                  *string           `json:"companyID"`
	HomeOwner                  homeOwner.Input   `json:"homeOwner"`
	SalesRep                   estimate.SalesRep `json:"salesRep"`
	CurrentMaterial            int               `json:"currentMaterial"`
	NewRoofingMaterial         *int              `json:"newRoofingMaterial"`
	LowSlope                   bool              `json:"lowSlope"`
	CurrentMaterialLowSlope    *int              `json:"currentMaterialLowSlope"`
	NewRoofingMaterialLowSlope *int              `json:"newRoofingMaterialLowSlope"`
	Redeck                     bool              `json:"redeck"`
	Layers                     int               `json:"layers"`
	Layer2Material             *int              `json:"layer2Material"`
	Layer3Material             *int              `json:"layer3Material"`
	MeasurementType            *int              `json:"measurementType"`
	PartialPercentage          *int              `json:"partialPercentage"`
}

// EstimateRequestHandler from roofix bubble app
func EstimateRequestHandler(w http.ResponseWriter, r *http.Request) {
	var id, note *string
	ctx := r.Context()
	action := audit.EstRequest
	apiUID := account.CtxApiUserID(ctx)
	failed := func(err error) {
		if id != nil {
			audit.ApiOpFailed(ctx, action, apiUID, fmt.Errorf("estimate: %s failed with error, %v", *id, err))
		} else {
			audit.ApiOpFailed(ctx, action, apiUID, err)
		}
		// return back with 200 with FailureStatus
		response.Ok(w, &nearmap.EstimateDetail{
			FailureStatus: err.Error(),
		})
	}

	// payload: read from body
	var raw map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
		failed(errors.New("failed to parse request body"))
		return
	}
	// payload: convert to JSON
	d, err := json.Marshal(raw)
	if err != nil {
		failed(errors.New("failed to bind model"))
		return
	}

	// find estimate id(externalID)
	if v, ok := raw["externalID"]; ok {
		id = ptr.Str(fmt.Sprintf("%v", v))
	}
	// in case externalID is not provided
	if str.Val(id) == "" || str.Val(id) == "null" {
		id = ptr.Str(uid.ULID())
		note = ptr.Str("** externalID was not provided")
	}

	// payload: dump to storage
	key := fmt.Sprintf("%s/%s/request.json", enum.FolderEstimates, *id)
	_ = storage.PutObject(ctx, config.DataBucket(), key, d, ptr.Str("application/json"))
	// <===

	// estimate: bind data
	var inp EstimateInput
	if err = json.Unmarshal(d, &inp); err != nil {
		failed(err)
		return
	}

	// estimate: in realtime
	inp.ExternalID = id
	if resp, err2 := estimateRequest(ctx, apiUID, &inp); err2 != nil {
		failed(err2)
		return
	} else {
		resp.Note = note
		audit.OpSuccess(ctx, action, fmt.Sprintf("created estimate: %s", *id))
		response.Ok(w, resp)
	}
}

func estimateRequest(ctx context.Context, apiUID string, inp *EstimateInput) (*nearmap.EstimateDetail, error) {
	cm := currentMaterial(inp.CurrentMaterial)
	partial := partialPercentage(inp.PartialPercentage)
	mt := measurementType(inp.MeasurementType)
	detail, err := estimate.MaterialDetailCheck(
		cm,
		newRoofingMaterial(inp.NewRoofingMaterial),
		newRoofingMaterialLowSlope(inp.NewRoofingMaterialLowSlope),
		inp.LowSlope,
		inp.Redeck,
		inp.Layers,
		layerMaterial(inp.Layer2Material, cm),
		layerMaterial(inp.Layer3Material, cm),
		&partial,
	)
	if err != nil {
		return nil, err
	}

	// instant price estimation
	est := &estimate.Input{
		ID:                inp.ExternalID,
		ExternalCompanyID: inp.CompanyID,
		HomeOwner:         &inp.HomeOwner,
		SalesRep:          &inp.SalesRep,
		MaterialDetail:    detail,
		MeasurementType:   mt,
		CtxApiUserID:      &apiUID,
	}

	db := ent.GetClient()
	defer db.CloseClient()

	coID := str.Val(inp.CompanyID)
	coName := inp.SalesRep.CompanyName
	// find partner by external CompanyID
	if coID != "" {
		co, er1 := db.Partner.Query().Where(partner.ExternalID(coID)).Select(partner.FieldID).First(ctx)
		if er1 != nil && ent.IsNotFound(er1) { // partner not found
			// this must not happen, in normal proces partner:company relation is created before estimate request is made
			//
			// create a new solar partner
			if id, er2 := pkgPartner.QuickCreateSolar(ctx, db, *inp.CompanyID, coName); er2 != nil {
				log.Error(err)
			} else {
				est.PartnerID = &id
			}
		} else if er1 != nil {
			log.Error(er1)
		} else {
			est.PartnerID = &co.ID
		}
	}

	// company vise extra charges
	if co := CompanyExtraCharges(coID); co != nil && co.ChargeType != enum.ExtraChargeNone {
		est.ExtraCharge = &price.ExtraCharge{
			Type:       co.ChargeType,
			Val:        co.ChargeVal,
			Note:       co.ChargeNote,
			Conditions: co.ChargeConditions,
		}
	}

	return estimate.RequestRealtime(ctx, est)
}

func currentMaterial(v int) material.CurrentSteepSlope {
	switch v {
	case 1:
		return material.ThreeTabShingles
	case 2:
		return material.Architectural
	case 3:
		return material.ClayTile
	case 4:
		return material.ConcreteTile
	case 5:
		return material.WoodShakes
	case 6:
		return material.MetalShakes
	case 7:
		return material.MetalTitle
	case 8:
		return material.StandingSeamMetal
	case 9:
		return material.Slate
	case 10:
		return material.MetalRPanelExpFastener
	case 11:
		return material.LowSlopeOnly
	default:
		return ""
	}
}

func newRoofingMaterial(v *int) material.NewSteepSlope {
	switch num.ReadIntPtr(v) {
	case 1:
		return material.BestValue
	case 2:
		return material.MoreExpensive
	case 3:
		return material.StandingSeamMetal
	case 4:
		return material.ConcreteTile
	case 5:
		return material.ClayTileBarrel
	case 7:
		return material.Repaper
	default:
		return ""
	}
}

func newRoofingMaterialLowSlope(v *int) material.NewLowSlope {
	switch num.ReadIntPtr(v) {
	case 1:
		return material.ModBit
	case 2:
		return material.Coating
	default:
		return ""
	}
}

// layerMaterial conversion.
// if no material is given, assume currentMaterial as material
func layerMaterial(v *int, cm material.CurrentSteepSlope) material.Layer {
	switch num.ReadIntPtr(v) {
	case 1:
		return material.ThreeTabShingles
	case 2:
		return material.Architectural
	case 5:
		return material.WoodShakes
	default:
		// Layers = 2 or 3 has no material assume that layer2Material and layer3Material are the same as currentMaterial.
		if cm != "" {
			return material.Layer(cm)
		}
		return ""
	}
}

func measurementType(v *int) enum.Measure {
	switch num.ReadIntPtr(v) {
	case 2:
		return enum.MeasurePrimaryPlusDetachedGarage
	case 3:
		return enum.MeasureAllStructuresOnParcel
	default:
		return enum.MeasurePrimaryStructureOnly
	}
}

func partialPercentage(v *int) int {
	switch num.ReadIntPtr(v) {
	case 1:
		return 30
	case 2:
		return 50
	case 3:
		return 75

	default:
		return 0
	}
}
