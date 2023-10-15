/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package partner

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/homeowner"
	"roofix/ent/job"
	entPartner "roofix/ent/partner"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

// JobsByProgress by partnerID and Progress AccountStatus
func JobsByProgress(ctx context.Context, partnerID string, progress enum.JobProgress, search *string, page model.PageInput) (*ent.JobConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Job.Query().
		WithRoofingPartner(func(p *ent.PartnerQuery) { p.Select(entPartner.FieldName) }).
		Where(func(t1 *sql.Selector) {
			t1.Where(sql.And(
				sql.EQ(t1.C(job.RoofingPartnerColumn), partnerID),
				sql.Or(
					sql.EQ(t1.C(job.FieldProgress), progress),
				),
			))
		})

	return qryJob(ctx, qry, page, search)
}

func qryJob(ctx context.Context, qry *ent.JobQuery, page model.PageInput, search *string) (*ent.JobConnection, error) {
	q := str.Val(search)
	if q != "" {
		qry.Where(func(j *sql.Selector) {
			ho := sql.Table(homeowner.Table)
			j.Join(ho).On(j.C(job.HomeOwnerColumn), ho.C(homeowner.FieldID))
			j.Where(sql.Or(
				sql.ContainsFold(ho.C(homeowner.FieldFirstName), q),
				sql.ContainsFold(ho.C(homeowner.FieldLastName), q),
				sql.ContainsFold(ho.C(homeowner.FieldStreetNumber), q),
				sql.ContainsFold(ho.C(homeowner.FieldStreetName), q),
				sql.ContainsFold(ho.C(homeowner.FieldCity), q),
				sql.ContainsFold(ho.C(homeowner.FieldState), q),
				sql.ContainsFold(ho.C(homeowner.FieldZip), q),
			))
		})
	}

	orderBy := &ent.JobOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.JobOrderFieldCreatedAt,
	}
	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithJobOrder(orderBy))
}
