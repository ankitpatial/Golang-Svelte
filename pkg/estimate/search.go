package estimate

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/ent/estimate"
	"roofix/ent/homeowner"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/util"
	"roofix/pkg/util/convert"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

func Search(
	ctx context.Context, partnerID, userID *string, statusIn []enum.EstimateStatus, search *string, dtRange []string,
	page model.PageInput,
) (*ent.EstimateConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Debug().Estimate.Query().
		WithPartner(func(p *ent.PartnerQuery) { p.Select(partner.FieldID, partner.FieldName) }).
		WithCreator(func(u *ent.UserQuery) { u.Select(user.FieldID, user.FieldFirstName, user.FieldLastName) }).
		WithCreatorAPI(func(api *ent.ApiUserQuery) { api.Select(apiuser.FieldUsername) }).
		Where(func(j *sql.Selector) {
			// partner or creator check
			if partnerID != nil || userID != nil {
				var p []*sql.Predicate
				// requester partner filter
				if partnerID != nil {
					p = append(p, sql.EQ(j.C(estimate.PartnerColumn), *partnerID))
				}

				if userID != nil {
					p = append(p, sql.EQ(j.C(estimate.CreatorColumn), *userID))  // creator
					p = append(p, sql.EQ(j.C(estimate.SalesRepColumn), *userID)) // sales Rep
				}

				j.Where(sql.Or(p...))
			}

			if len(dtRange) == 2 {
				f := util.ParseISODate(dtRange[0])
				t := util.ParseISODate(dtRange[1])
				j.Where(sql.And(
					sql.GTE(j.C(estimate.FieldCreatedAt), f),
					sql.LTE(j.C(estimate.FieldCreatedAt), t),
				))
			}

			// status in
			if len(statusIn) > 0 {
				j.Where(sql.In(j.C(estimate.FieldStatus), convert.ToAny(statusIn)...))
			}

			// search
			if s := str.TrimSpace(search); s != "" {
				owner := sql.Table(homeowner.Table)
				j.Join(owner).On(j.C(estimate.HomeOwnerColumn), owner.C(homeowner.FieldID))
				j.Where(sql.Or(
					sql.ContainsFold(owner.C(homeowner.FieldFirstName), s),
					sql.ContainsFold(owner.C(homeowner.FieldLastName), s),
					sql.ContainsFold(owner.C(homeowner.FieldStreetNumber), s),
					sql.ContainsFold(owner.C(homeowner.FieldStreetName), s),
				))
			}
		})

	// order-by
	orderBy := &ent.EstimateOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.EstimateOrderFieldCreatedAt,
	}

	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithEstimateOrder(orderBy))
}
