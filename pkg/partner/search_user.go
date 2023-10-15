package partner

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/server/model"
	"strings"
)

func SearchUsers(ctx context.Context, ctxUID, partnerID, search string, page *model.PageInput) (*ent.UserConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	orderBy := &ent.UserOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.UserOrderFieldFirstName,
	}
	return db.User.Query().
		Where(func(u *sql.Selector) {
			p := sql.Table(partner.Table)
			pc := sql.Table(partnercontact.Table)
			u.
				Join(pc).On(u.C(user.FieldID), pc.C(partnercontact.UserColumn)).
				Join(p).On(pc.C(partnercontact.FieldPartnerID), p.C(partner.FieldID))

			u.Where(sql.NEQ(u.C(user.FieldID), ctxUID))
			u.Where(sql.EQ(p.C(partner.FieldID), partnerID))

			q := strings.TrimSpace(search)
			if len(q) > 0 {
				u.Where(sql.Or(
					sql.ContainsFold(u.C(user.FieldEmail), q),
					sql.ContainsFold(u.C(user.FieldFirstName), q),
					sql.ContainsFold(u.C(user.FieldLastName), q),
					sql.ContainsFold(u.C(user.FieldPhone), q),
				))
			}

		}).Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithUserOrder(orderBy))
}
