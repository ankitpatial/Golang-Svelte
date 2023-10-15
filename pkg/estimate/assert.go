package estimate

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/estimate"
	"roofix/pkg/account"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
)

func AssertAccess(ctx context.Context, id string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := accessFilter(ctx, db.Estimate.Query().Where(estimate.ID(id)))
	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.JobNotAssigned)
	}

	return nil
}

// AssertAdminOrCompanyAdmin job access by checking context user role.
// will assert context-users job access
func AssertAdminOrCompanyAdmin(ctx context.Context, id string) error {
	u := account.CtxUser(ctx)
	if !u.IsAdmin && !u.IsCompanyAdmin {
		return msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Estimate.Query().Where(estimate.ID(id))

	// company-ADMIN
	if u.IsCompanyAdmin {
		qry.Where(func(j *sql.Selector) {
			j.Where(sql.EQ(j.C(estimate.PartnerColumn), u.Partner.ID))
		})
	}

	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.NotFound, "estimate")
	}

	return nil
}

func accessFilter(ctx context.Context, qry *ent.EstimateQuery) *ent.EstimateQuery {
	u := account.CtxUser(ctx)

	// none ADMIN check
	if !u.IsAdmin {
		qry.Where(func(e *sql.Selector) {
			// creator filter
			pred := []*sql.Predicate{
				sql.EQ(e.C(estimate.CreatorColumn), u.ID),
				sql.EQ(e.C(estimate.SalesRepColumn), u.ID),
			}

			p := u.Partner
			if u.IsCompanyAdmin && p != nil {
				pred = append(pred, sql.EQ(e.C(estimate.PartnerColumn), p.ID))
			}
			e.Where(sql.Or(pred...))
		})
	}

	return qry
}
