package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/ent/homeowner"
	entJob "roofix/ent/job"
	"roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/job"
	"roofix/pkg/nearmap"
	"roofix/pkg/partner"
	"roofix/pkg/pricing"
	"roofix/pkg/util/ptr"
	"roofix/server/generated"
	"roofix/server/model"
	"time"

	"entgo.io/contrib/entgql"
)

// ProgressFlagged is the resolver for the progressFlagged field.
func (r *jobResolver) ProgressFlagged(ctx context.Context, obj *ent.Job) (*bool, error) {
	if obj == nil || obj.Progress == nil || obj.ProgressFlagAt == nil {
		return ptr.Bool(false), nil
	}

	p := obj.Progress
	if p == nil {
		n := enum.JobProgressNew
		p = &n
	}

	flag := time.Now().UTC().After(obj.ProgressFlagAt.UTC())
	return ptr.Bool(flag), nil
}

// Contractor is the resolver for the contractor field.
func (r *jobResolver) Contractor(ctx context.Context, obj *ent.Job) (*model.Entity, error) {
	if obj == nil || obj.Edges.RoofingPartner == nil {
		return nil, nil
	}

	p := obj.Edges.RoofingPartner
	return &model.Entity{
		ID:   p.ID,
		Name: p.Name,
	}, nil
}

// CreatorName is the resolver for the creatorName field.
func (r *jobResolver) CreatorName(ctx context.Context, obj *ent.Job) (string, error) {
	if obj == nil || (obj.Edges.Creator == nil && obj.Edges.CreatorAPI == nil) {
		return "", nil
	}

	if obj.Edges.Creator != nil {
		c := obj.Edges.Creator
		return fmt.Sprintf("%s %s", c.FirstName, c.LastName), nil
	}

	au := obj.Edges.CreatorAPI
	return au.Username, nil
}

// EpcName is the resolver for the epcName field.
func (r *jobResolver) EpcName(ctx context.Context, obj *ent.Job) (string, error) {
	if obj == nil || obj.Edges.Epc == nil {
		return obj.CompanyName, nil
	}

	return obj.Edges.Epc.Name, nil
}

// AssignPartnerToJob is the resolver for the assignPartnerToJob field.
func (r *mutationResolver) AssignPartnerToJob(ctx context.Context, jobID string, partnerID string) (string, error) {
	action := audit.JobUpdate
	if id, err := job.AssignPartner(ctx, jobID, partnerID); err != nil {
		audit.OpFailed(ctx, action, err)
		return "", err
	} else {
		audit.OpSuccess(ctx, action, fmt.Sprintf("assign partnerID: %s to jobID: %s", partnerID, jobID))
		job.NotifyPartnerAssigned(ctx, jobID, partnerID)
		return id, err
	}
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*ent.Job, error) {
	// assert user access to job
	if err := job.AssertAccess(ctx, id); err != nil {
		return nil, err
	}

	return ent.GetClient().Job.Query().
		WithCreator(func(u *ent.UserQuery) { u.Select(user.FieldFirstName, user.FieldLastName) }).
		WithCreatorAPI(func(au *ent.ApiUserQuery) { au.Select(apiuser.FieldUsername) }).
		Where(entJob.ID(id)).
		Only(ctx)
}

// MyJob is the resolver for the myJob field.
func (r *queryResolver) MyJob(ctx context.Context, id string) (*ent.Job, error) {
	return r.Job(ctx, id)
}

// JobGeoCode is the resolver for the jobGeoCode field.
func (r *queryResolver) JobGeoCode(ctx context.Context, id string) ([]*model.Point, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	j, err := db.Job.Query().
		WithHomeOwner(func(ho *ent.HomeOwnerQuery) {
			// select lat, lng
			ho.Select(homeowner.FieldLatitude, homeowner.FieldLongitude)
		}).
		Where(entJob.ID(id)).
		Select(entJob.FieldID).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	// get parcel geometry coordinates
	geo, err := nearmap.ParcelGeometryPolygon(ctx, j.Edges.HomeOwner.Latitude, j.Edges.HomeOwner.Longitude)
	if err != nil {
		return nil, err
	}

	var res []*model.Point
	for _, p := range geo {
		res = append(res, &model.Point{
			Lat: p.Lat,
			Lng: p.Lng,
		})
	}

	return res, nil
}

// UnassignedJobs is the resolver for the jobs field.
func (r *queryResolver) UnassignedJobs(ctx context.Context, progress *enum.JobProgress, search *string, betweenDates []string, page model.PageInput, orderBy *ent.JobOrder) (*ent.JobConnection, error) {
	var progressIn []enum.JobProgress
	if progress != nil {
		progressIn = append(progressIn, *progress)
	}

	isAdmin, uID, pID, _ := account.Info(ctx)
	if isAdmin {
		return job.Search(ctx, nil, nil, progressIn, nil, search, betweenDates, ptr.Bool(false), page, orderBy)
	} else {
		return job.Search(ctx, pID, uID, progressIn, nil, search, betweenDates, ptr.Bool(false), page, orderBy)
	}
}

// AssignedJobs is the resolver for the jobs field.
func (r *queryResolver) AssignedJobs(ctx context.Context, progress *enum.JobProgress, search *string, betweenDates []string, page model.PageInput, orderBy *ent.JobOrder) (*ent.JobConnection, error) {
	var progressIn []enum.JobProgress
	if progress != nil {
		progressIn = append(progressIn, *progress)
	}

	isAdmin, uID, pID, _ := account.Info(ctx)
	if isAdmin {
		return job.Search(ctx, nil, nil, progressIn, nil, search, betweenDates, ptr.Bool(true), page, orderBy)
	} else {
		return job.Search(ctx, pID, uID, progressIn, nil, search, betweenDates, ptr.Bool(true), page, orderBy)
	}
}

// JobsByProgress is the resolver for the jobsByProgress field.
func (r *queryResolver) JobsByProgress(ctx context.Context, status enum.JobProgress, search *string, betweenDates []string, page model.PageInput, orderBy *ent.JobOrder) (*ent.JobConnection, error) {
	return job.ByProgress(ctx, status, search, betweenDates, page, orderBy)
}

// PartnerJobStats is the resolver for the partnerJobStats field.
func (r *queryResolver) PartnerJobStats(ctx context.Context, search *string, partnerType enum.Partner, skip int, take int) ([]*partner.JobStats, error) {
	return partner.Stats(ctx, partnerType, search, skip, take)
}

// PartnerJobs is the resolver for the partnerJobs field.
func (r *queryResolver) PartnerJobs(ctx context.Context, partnerID string, search *string, flagged *bool, progress *enum.JobProgress, dates []string, page model.PageInput) (*ent.JobConnection, error) {
	var progressIn []enum.JobProgress
	// progress
	if progress != nil {
		progressIn = append(progressIn, *progress)
	}

	// order
	order := &ent.JobOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.JobOrderFieldCreatedAt,
	}

	isAdmin, uID, pID, _ := account.Info(ctx)
	if isAdmin {
		return job.Search(ctx, &partnerID, nil, progressIn, nil, search, dates, nil, page, order)
	}

	return job.SearchOwned(ctx, uID, pID, progressIn, flagged, search, dates, page, order)
}

// ApprovedJobs is the resolver for the approvedJobs field.
func (r *queryResolver) ApprovedJobs(ctx context.Context, search *string, progress *enum.JobProgress, dates []string, page model.PageInput) (*ent.JobConnection, error) {
	var progressIn []enum.JobProgress
	// progress
	if progress != nil {
		progressIn = append(progressIn, *progress)
	}

	// order
	order := &ent.JobOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.JobOrderFieldCreatedAt,
	}

	isAdmin, uID, pID, _ := account.Info(ctx)
	if isAdmin {
		return job.Search(ctx, nil, nil, progressIn, nil, search, dates, nil, page, order)
	}

	// search jobs
	return job.SearchOwned(ctx, uID, pID, progressIn, nil, search, dates, page, order)
}

// JobEstimates is the resolver for the jobEstimates field.
func (r *queryResolver) JobEstimates(ctx context.Context, jobID string) (*model.JobEstimates, error) {
	return nil, errors.New("not in use")
}

// Pitch is the resolver for the pitch field.
func (r *measurementResolver) Pitch(ctx context.Context, obj *pricing.Measurement, data string) error {
	if obj == nil {
		return nil
	}
	switch data {
	case "1/12":
		obj.Pitch = 1
	case "2/12":
		obj.Pitch = 2
	case "3/12":
		obj.Pitch = 3
	case "4/12":
		obj.Pitch = 4
	case "5/12":
		obj.Pitch = 5
	case "6/12":
		obj.Pitch = 6
	case "7/12":
		obj.Pitch = 7
	case "8/12":
		obj.Pitch = 8
	case "9/12":
		obj.Pitch = 9
	case "10/12":
		obj.Pitch = 10
	case "11/12":
		obj.Pitch = 11
	case "12/12":
		obj.Pitch = 12
	case "13/12":
		obj.Pitch = 13
	}
	return nil
}

// Job returns generated.JobResolver implementation.
func (r *Resolver) Job() generated.JobResolver { return &jobResolver{r} }

// Measurement returns generated.MeasurementResolver implementation.
func (r *Resolver) Measurement() generated.MeasurementResolver { return &measurementResolver{r} }

type jobResolver struct{ *Resolver }
type measurementResolver struct{ *Resolver }
