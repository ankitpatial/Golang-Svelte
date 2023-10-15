package job

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/ent/estimate"
	"roofix/ent/homeowner"
	entJob "roofix/ent/job"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/util"
	"roofix/pkg/util/str"
	"roofix/server/model"
	"strings"
)

func Search(
	ctx context.Context, partnerID, creatorID *string,
	progressIn []enum.JobProgress, progressFlagged *bool,
	search *string, dtRange []string, isAssigned *bool, page model.PageInput, orderBy *ent.JobOrder,
) (*ent.JobConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := withEdges(db).Where(func(j *sql.Selector) {
		if isAssigned != nil {
			if *isAssigned {
				// partner is assigned
				j.Where(sql.NotNull(j.C(entJob.RoofingPartnerColumn)))
			} else {
				// partner is not assigned
				j.Where(sql.IsNull(j.C(entJob.RoofingPartnerColumn)))
			}
		}

		// partner or creator check
		if partnerID != nil || creatorID != nil {
			var p []*sql.Predicate
			// requester partnerI filter
			if partnerID != nil {
				p = append(p, sql.EQ(j.C(entJob.RoofingPartnerColumn), *partnerID))
				p = append(p, sql.EQ(j.C(entJob.RequesterColumn), *partnerID))
			}
			// creator filter
			if creatorID != nil {
				p = append(p, sql.EQ(j.C(entJob.CreatorColumn), *creatorID))
			}

			j.Where(sql.Or(p...))
		}
	})

	return searchFilters(qry, progressIn, progressFlagged, search, dtRange).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithJobOrder(orderBy))
}

func SearchAssigned(
	ctx context.Context, partnerID string, progressIn []enum.JobProgress, progressFlagged *bool, search *string,
	dtRange []string, page model.PageInput, orderBy *ent.JobOrder,
) (*ent.JobConnection, error) {
	if strings.TrimSpace(partnerID) == "" {
		return nil, errors.New("param partnerID is empty")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := withEdges(db).Where(func(j *sql.Selector) {
		// partner or creator check
		j.Where(sql.EQ(j.C(entJob.RoofingPartnerColumn), partnerID))
	})

	return searchFilters(qry, progressIn, progressFlagged, search, dtRange).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithJobOrder(orderBy))
}

func SearchOwned(
	ctx context.Context, userID, partnerID *string, progressIn []enum.JobProgress, progressFlagged *bool,
	search *string, dtRange []string, page model.PageInput, orderBy *ent.JobOrder,
) (*ent.JobConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := withEdges(db).Where(func(j *sql.Selector) {
		var p []*sql.Predicate

		// creator or sales rep
		if userID != nil {
			p = append(
				p,
				sql.EQ(j.C(entJob.CreatorColumn), *userID),
				sql.EQ(j.C(entJob.SalesRepColumn), *userID),
			)
		}

		// contractor & partner
		if partnerID != nil {
			p = append(
				p,
				sql.EQ(j.C(entJob.RoofingPartnerColumn), *partnerID), // assigned partner (contractor)
				sql.EQ(j.C(entJob.RequesterColumn), *partnerID),      // requester partner
			)
		}

		j.Where(sql.Or(p...))
	})

	return searchFilters(qry, progressIn, progressFlagged, search, dtRange).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithJobOrder(orderBy))
}

func ByProgress(
	ctx context.Context, status enum.JobProgress, search *string, betweenDates []string, page model.PageInput,
	orderBy *ent.JobOrder,
) (*ent.JobConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := withEdges(db).Where(entJob.ProgressEQ(status))

	if len(betweenDates) == 2 {
		f := util.ParseISODate(betweenDates[0])
		t := util.ParseISODate(betweenDates[1])
		//logger.Info("from: %v, to: %v", f, t)
		qry.Where(entJob.ProgressAtGTE(f), entJob.ProgressAtLTE(t))
	}

	return searchFields(qry, search).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithJobOrder(orderBy))
}

// withEdges returns a query that loads all edges of the given node.
func withEdges(db *ent.Client) *ent.JobQuery {
	return db.Job.Query().
		WithHomeOwner().
		WithSalesRep(func(u *ent.UserQuery) {
			u.Select(user.FieldFirstName, user.FieldLastName, user.FieldEmail, user.FieldPhone)
		}).
		WithEstimate(func(est *ent.EstimateQuery) {
			//select material related fields
			est.Select(
				estimate.FieldCurrentMaterial,
				estimate.FieldNewRoofingMaterial,
				estimate.FieldLowSlope,
				estimate.FieldCurrentMaterialLowSlope,
				estimate.FieldNewRoofingMaterialLowSlope,
				estimate.FieldRedeck,
				estimate.FieldLayers,
				estimate.FieldLayer2Material,
				estimate.FieldLayer3Material,
				estimate.FieldPartialPercentage,
				estimate.FieldMeasureType,
			)
		}).
		WithRoofingPartner(func(p *ent.PartnerQuery) {
			p.Select(partner.FieldName)
		}).
		WithCreator(func(c *ent.UserQuery) {
			c.Select(user.FieldFirstName, user.FieldLastName)
		})
}

func searchFilters(
	qry *ent.JobQuery,
	progressIn []enum.JobProgress, progressFlagged *bool,
	search *string, dtRange []string,
) *ent.JobQuery {
	qry.Where(func(j *sql.Selector) {
		// flagged progress
		if progressFlagged != nil && *progressFlagged {
			j.Where(sql.ExprP(fmt.Sprintf("%s <= utc_timestamp", j.C(entJob.FieldProgressFlagAt))))
		}

		// date range
		if len(dtRange) == 2 {
			f := util.ParseISODate(dtRange[0])
			t := util.ParseISODate(dtRange[1])
			j.Where(sql.And(
				sql.GTE(j.C(entJob.FieldCreatedAt), f),
				sql.LTE(j.C(entJob.FieldCreatedAt), t),
			))
		}

	})

	// progres in
	if len(progressIn) > 0 {
		if idx := util.IndexOf(progressIn, enum.JobProgressNew); idx > -1 {
			// approved jobs has field progress is null, to show them for solar partner Approved Jobs we need to allow
			// so `New` progress status can include int the 'progress is null' records
			qry.Where(entJob.Or(
				entJob.ProgressIsNil(),
				entJob.ProgressIn(progressIn...),
			))
		} else {
			qry.Where(entJob.ProgressIn(progressIn...))
		}
	}

	return searchFields(qry, search)
}

func searchFields(qry *ent.JobQuery, search *string) *ent.JobQuery {
	if s := str.TrimSpace(search); s != "" {
		// search by: company name, sales rep name, customer name, state, date created, etc.
		qry.Where(func(j *sql.Selector) {
			// table homeowner
			ho := sql.Table(homeowner.Table)
			// table partner
			p := sql.Table(partner.Table)

			// join with homeowner and search homeowner fields
			j.Join(ho).On(j.C(entJob.HomeOwnerColumn), ho.C(homeowner.FieldID))
			// outer join with partner and search partner fields
			j.LeftJoin(p).On(j.C(entJob.RoofingPartnerColumn), p.C(partner.FieldID))
			// search fields
			j.Where(sql.Or(
				sql.ContainsFold(p.C(partner.FieldName), s),
				sql.ContainsFold(ho.C(homeowner.FieldFirstName), s),
				sql.ContainsFold(ho.C(homeowner.FieldLastName), s),
				sql.ContainsFold(ho.C(homeowner.FieldStreetNumber), s),
				sql.ContainsFold(ho.C(homeowner.FieldStreetName), s),
				sql.ContainsFold(ho.C(homeowner.FieldCity), s),
				sql.ContainsFold(ho.C(homeowner.FieldState), s),
				sql.ContainsFold(ho.C(homeowner.FieldZip), s),
			))
		})

		qry.Where(func(j *sql.Selector) {
			// table homeowner
			ho := sql.Table(homeowner.Table)
			// table sales rep
			sr := sql.Table(user.Table)

			// join with homeowner and search homeowner fields
			j.Join(ho).On(j.C(entJob.HomeOwnerColumn), ho.C(homeowner.FieldID))
			j.LeftJoin(sr).On(j.C(entJob.SalesRepColumn), sr.C(user.FieldID))

			j.Where(sql.Or(
				sql.Like(ho.C(homeowner.FieldFirstName), s),
				sql.Like(ho.C(homeowner.FieldLastName), s),
				sql.Like(ho.C(homeowner.FieldEmail), s),
				sql.Like(ho.C(homeowner.FieldPhone), s),
			))
		})
	}

	return qry
}
