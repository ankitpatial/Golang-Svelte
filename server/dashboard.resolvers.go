package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/pkg/job"
	"roofix/pkg/postal"
	"roofix/server/model"
)

// Overview is the resolver for the overview field.
func (r *queryResolver) Overview(ctx context.Context, f model.Filter) (*model.Overview, error) {
	res := &model.Overview{ID: f}
	switch f {
	case model.FilterByRegion:
		items, err := job.OverviewByRegion(ctx)
		if err != nil {
			return nil, err
		}

		for idx, item := range items {
			var r postal.Region
			r.FromUInt8(item.RegionID)
			res.Total += item.Count
			res.Items = append(res.Items, &model.OverviewItem{
				ID:    fmt.Sprintf("%d", idx),
				Name:  r.String(),
				Count: item.Count,
			})
		}
	case model.FilterByState:
		items, err := job.OverviewState(ctx)
		if err != nil {
			return nil, err
		}

		for idx, item := range items {
			res.Total += item.Count
			res.Items = append(res.Items, &model.OverviewItem{
				ID:    fmt.Sprintf("%d", idx),
				Name:  item.State,
				Count: item.Count,
			})
		}

	}

	return res, nil
}
