/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package job

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/homeowner"
	"roofix/ent/job"
)

type RegionCount struct {
	RegionID uint8 `json:"region_id"`
	Count    int
}

type StateCount struct {
	State string
	Count int
}

func OverviewByRegion(ctx context.Context) ([]*RegionCount, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var res []*RegionCount
	err := db.Job.
		Query().
		GroupBy(job.FieldRegionID).
		Aggregate(ent.Count()).
		Scan(ctx, &res)

	return res, err
}

func OverviewState(ctx context.Context) ([]*StateCount, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	var res []*StateCount
	err := db.Job.
		Query().
		Where(func(j *sql.Selector) {
			// table homeowner
			ho := sql.Table(homeowner.Table)
			j.Join(ho).On(j.C(job.HomeOwnerColumn), ho.C(homeowner.FieldID))
			j.GroupBy(ho.C(homeowner.FieldState))
			j.OrderBy(sql.Asc(ho.C(homeowner.FieldState)))
		}).
		Aggregate(ent.Count()).
		Scan(ctx, &res)

	return res, err
}
