package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"roofix/ent"
	entEstimate "roofix/ent/estimate"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/estimate"
	"roofix/pkg/homeOwner"
	"roofix/pkg/material"
	"roofix/pkg/nearmap"
	"roofix/pkg/pricing"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/pkg/util/ptr"
	"roofix/pkg/util/str"
	"roofix/server/generated"
	"roofix/server/model"
	"strings"
)

// Layers is the resolver for the layers field.
func (r *estimateResolver) Layers(ctx context.Context, obj *ent.Estimate) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.Layers), nil
}

// MeasurementType is the resolver for the measurementType field.
func (r *estimateResolver) MeasurementType(ctx context.Context, obj *ent.Estimate) (enum.Measure, error) {
	if obj == nil {
		return "", nil
	}

	return obj.MeasureType, nil
}

// Partial is the resolver for the partial field.
func (r *estimateResolver) Partial(ctx context.Context, obj *ent.Estimate) (*int, error) {
	if obj == nil {
		return nil, nil
	}

	v := int(obj.PartialPercentage)
	return &v, nil
}

// Bounds is the resolver for the bounds field.
func (r *estimateResolver) Bounds(ctx context.Context, obj *ent.Estimate) ([]*model.Point, error) {
	if obj == nil {
		return nil, nil
	}

	var b []*model.Point
	for _, p := range obj.Bounds {
		b = append(b, &model.Point{
			Lat: p.Lat,
			Lng: p.Lng,
		})
	}
	return b, nil
}

// CompanyName is the resolver for the companyName field.
func (r *estimateResolver) CompanyName(ctx context.Context, obj *ent.Estimate) (string, error) {
	if obj == nil {
		return "", nil
	}

	if obj.Edges.Partner != nil {
		return obj.Edges.Partner.Name, nil
	}

	return obj.CompanyRefName, nil
}

// CreatorName is the resolver for the creatorName field.
func (r *estimateResolver) CreatorName(ctx context.Context, obj *ent.Estimate) (string, error) {
	if obj == nil {
		return "", nil
	}

	if obj.Edges.Creator != nil {
		return fmt.Sprintf("%s %s", obj.Edges.Creator.FirstName, obj.Edges.Creator.LastName), nil
	}

	if obj.Edges.CreatorAPI != nil {
		return obj.Edges.CreatorAPI.Username, nil
	}

	return "", nil
}

// PDF is the resolver for the pdf field.
func (r *estimateResolver) PDF(ctx context.Context, obj *ent.Estimate) (*document.InfoShort, error) {
	if obj == nil || obj.Edges.Pdf == nil {
		return nil, nil
	}

	return document.AsDocShortInfo(obj.Edges.Pdf), nil
}

// CreateEstimate is the resolver for the createJob field.
func (r *mutationResolver) CreateEstimate(ctx context.Context, input model.CreateEstimateInput) (string, error) {
	u := account.CtxUser(ctx)
	var partnerID, companyName *string
	if u.Partner != nil {
		partnerID = ptr.Str(u.Partner.ID)
		companyName = &u.Partner.PartnerName
	} else {
		companyName = ptr.Str(enum.MasterCompanyName)
	}

	detail, err := estimate.MaterialDetailCheck(
		material.CurrentSteepSlope(input.CurrentMaterial),
		material.NewSteepSlope(str.Val(input.NewRoofingMaterial)),
		material.NewLowSlope(str.Val(input.NewRoofingMaterialLowSlope)),
		true,
		input.Redeck,
		input.Layers,
		material.Layer(str.Val(input.Layer2Material)),
		material.Layer(str.Val(input.Layer3Material)),
		input.Partial,
	)
	if err != nil {
		return "", err
	}

	j := &estimate.Input{
		HomeOwner: &homeOwner.Input{
			FirstName:    input.OwnerFirstName,
			LastName:     input.OwnerLastName,
			StreetNumber: input.StreetNumber,
			StreetName:   input.StreetName,
			City:         input.City,
			State:        input.State,
			Zip:          input.Zip,
			Latitude:     num.ReadFloatPtr(input.Latitude),
			Longitude:    num.ReadFloatPtr(input.Longitude),
		},
		SalesRep: &estimate.SalesRep{
			CompanyName: str.Val(companyName),
			FirstName:   input.RepFirstName,
			LastName:    input.RepLastName,
			Email:       input.RepEmail,
			Mobile:      input.RepMobile,
		},
		MaterialDetail:  detail,
		MeasurementType: enum.Measure(input.MeasurementType),
		PartnerID:       partnerID,
		CtxUserID:       &u.ID,
	}

	id, err := estimate.Request(ctx, j)
	audit.Operation(ctx, audit.EstRequest, fmt.Sprintf("created estimate: %s", id), err)
	return id, err
}

// RequestAnEstimate is the resolver for the requestAnEstimate field.
func (r *mutationResolver) RequestAnEstimate(ctx context.Context, inp *model.EstimateRequest) (string, error) {
	action := audit.EstRequest
	// owner & parcel details
	owner := &homeOwner.Input{
		Email:        &inp.OwnerEmail,
		Phone:        &inp.OwnerPhone,
		StreetNumber: inp.StreetNumber,
		StreetName:   inp.StreetName,
		City:         inp.City,
		State:        inp.State,
		Zip:          inp.Zip,
		Latitude:     inp.Latitude,
		Longitude:    inp.Longitude,
	}
	name := strings.Split(inp.OwnerName, " ")
	switch {
	case len(name) == 1:
		owner.FirstName = name[0]
		owner.LastName = " "
	case len(name) >= 2:
		owner.FirstName = name[0]
		owner.LastName = strings.Join(name[1:], " ")
	}

	// sales rep detail.
	var partnerID *string
	u := account.CtxUser(ctx)
	rep := &estimate.SalesRep{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Mobile:    u.Phone,
	}

	if u.Partner != nil {
		partnerID = &u.Partner.ID
		rep.CompanyName = u.Partner.PartnerName
	} else {
		rep.CompanyName = enum.MasterCompanyName
	}

	// material detail
	detail, err := estimate.MaterialDetailCheck(
		estimateCurrentMaterial(inp.CurrentMaterial),
		estimateNewMaterial(inp.NewRoofingMaterial),
		material.ModBit,
		true,
		inp.Redeck,
		inp.Layers,
		"",
		"",
		inp.Partial,
	)
	if err != nil {
		return "", err
	}

	e := &estimate.Input{
		HomeOwner:      owner,
		SalesRep:       rep,
		MaterialDetail: detail,
		CtxUserID:      &u.ID,
		PartnerID:      partnerID,
	}

	// MeasurementType
	if inp.IncludeDetachedStructure != nil && *inp.IncludeDetachedStructure {
		e.MeasurementType = enum.MeasurePrimaryPlusDetachedGarage
	} else {
		e.MeasurementType = enum.MeasurePrimaryStructureOnly
	}

	// save, Job(Estimate Request) data

	if _, err = estimate.Request(ctx, e); err != nil {
		audit.OpFailed(ctx, action, err)
		return "", err
	} else {
		return "Successfully saved & estimated your request", nil
	}
}

// ApproveEstimate is the resolver for the approveEstimate field.
func (r *mutationResolver) ApproveEstimate(ctx context.Context, input model.ApproveEstimateInput) (bool, error) {
	if err := estimate.AssertAdminOrCompanyAdmin(ctx, input.ID); err != nil {
		return false, err
	}

	u := account.CtxUser(ctx)
	if u.IsCompanyAdmin {
		if input.Agree == nil {
			return false, errors.New("please ensure that you agree to the terms and conditions")
		}
	}

	if _, err := estimate.Approve(ctx, &input, &u.ID, nil); err != nil {
		log.Error(err)
		return false, err
	}

	estimate.NotifyApproved(ctx, input.ID)
	return true, nil
}

// DenyEstimate is the resolver for the denyEstimate field.
func (r *mutationResolver) DenyEstimate(ctx context.Context, input model.DenyEstimateInput) (bool, error) {
	if err := estimate.AssertAdminOrCompanyAdmin(ctx, input.ID); err != nil {
		return false, err
	}

	uID := account.CtxUserID(ctx)
	if err := estimate.Deny(ctx, &input, &uID, nil); err != nil {
		log.Error(err)
		return false, err
	}

	estimate.NotifyDenied(ctx, input.ID)
	return true, nil
}

// RemoveDenied is the resolver for the removeDenied field.
func (r *mutationResolver) RemoveDenied(ctx context.Context, id string) (bool, error) {
	if err := estimate.AssertAdminOrCompanyAdmin(ctx, id); err != nil {
		return false, err
	}

	uID := account.CtxUserID(ctx)
	if err := estimate.UndoDeny(ctx, id, "removed Denied status", &uID, nil); err != nil {
		log.Error(err)
		return false, err
	}

	// notify
	estimate.NotifyEstimateUnDenied(ctx, id)

	return true, nil
}

// TestPricing is the resolver for the testJobPricing field.
func (r *mutationResolver) TestPricing(ctx context.Context, job model.CreateEstimateInput, measure []*pricing.Measurement) (*model.PriceSummary, error) {
	j := &pricing.Data{
		State:                      job.State,
		Zip:                        job.Zip,
		Lat:                        num.ReadFloatPtr(job.Latitude),
		Lng:                        num.ReadFloatPtr(job.Longitude),
		CurrentMaterial:            job.CurrentMaterial,
		NewRoofingMaterial:         str.Val(job.NewRoofingMaterial),
		LowSlope:                   job.LowSlope,
		CurrentMaterialLowSlope:    str.Val(job.CurrentMaterialLowSlope),
		NewRoofingMaterialLowSlope: str.Val(job.NewRoofingMaterialLowSlope),
		Redeck:                     job.Redeck,
		Layers:                     uint8(job.Layers),
		Layer2Material:             str.Val(job.Layer2Material),
		Layer3Material:             str.Val(job.Layer3Material),
		PartialPercentage:          float64(num.ReadIntPtr(job.Partial)),
	}

	for _, m := range measure {
		m.CurrentMaterial = pricing.CurrentMaterial(j.CurrentMaterial)
	}

	p, err := pricing.Calculate(ctx, j, measure)
	if err != nil {
		return nil, err
	}

	return &model.PriceSummary{
		Total:   p.Total,
		Summary: p.Summary,
	}, nil
}

// Estimate is the resolver for the estimate field.
func (r *queryResolver) Estimate(ctx context.Context, id string) (*ent.Estimate, error) {
	return estimate.Get(ctx, id) // check access and get
}

// Estimates is the resolver for the estimates field.
func (r *queryResolver) Estimates(ctx context.Context, status *enum.EstimateStatus, search *string, dtRange []string, page model.PageInput) (*ent.EstimateConnection, error) {
	isAdmin, uID, pID, _ := account.Info(ctx)

	var statusIn []enum.EstimateStatus
	if status != nil { // user's status filter
		statusIn = append(statusIn, *status)
	}

	if isAdmin {
		return estimate.Search(ctx, nil, nil, statusIn, search, dtRange, page)
	} else {
		return estimate.Search(ctx, pID, uID, statusIn, search, dtRange, page)
	}
}

// NearmapResponse is the resolver for the nearmapResponse field.
func (r *queryResolver) NearmapResponse(ctx context.Context, id string, respID string) (*nearmap.ResponseAndDetail, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	res, err := db.Estimate.Query().Where(entEstimate.ID(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	var resp nearmap.Response
	d, err := json.Marshal(res.EstimatorRawResponse)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(d, &resp)
	if err != nil {
		return nil, err
	}

	primary := res.MeasureType == enum.MeasurePrimaryStructureOnly
	return &nearmap.ResponseAndDetail{
		Detail: nearmap.FillEstimateDetail(resp.Rollups, primary),
		Raw:    res.EstimatorRawResponse,
	}, nil
}

// Estimate returns generated.EstimateResolver implementation.
func (r *Resolver) Estimate() generated.EstimateResolver { return &estimateResolver{r} }

type estimateResolver struct{ *Resolver }
