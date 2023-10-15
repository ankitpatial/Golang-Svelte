package partner

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/pkg/enum"
)

func PrimaryContact(ctx context.Context, id string) (*Contact, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	t := enum.PartnerContactPrimary
	var res []*Contact
	err := db.User.Query().
		Where(func(u *sql.Selector) {
			contactSelector(u, id, &t)
		}).
		Select().
		Scan(ctx, &res)

	if err != nil {
		return nil, err
	} else if len(res) == 0 {
		return nil, nil
	}

	return res[0], err
}

func contactSelector(u *sql.Selector, partnerID string, pt *enum.PartnerContact) {
	pc := sql.Table(partnercontact.Table)
	p := sql.Table(partner.Table)
	u.Join(pc).On(u.C(user.FieldID), pc.C(partnercontact.UserColumn))
	u.Join(p).On(pc.C(partnercontact.PartnerColumn), p.C(partner.FieldID))

	u.Where(sql.EQ(pc.C(partnercontact.PartnerColumn), partnerID))
	if pt != nil {
		u.Where(sql.EQ(pc.C(partnercontact.FieldType), *pt))
	}

	u.Select(
		fmt.Sprintf("%s as id", u.C(user.FieldID)),
		fmt.Sprintf("%s as email", u.C(user.FieldEmail)),
		fmt.Sprintf("concat_ws(' ', %s, %s) as name", u.C(user.FieldFirstName), u.C(user.FieldLastName)),
		fmt.Sprintf("%s as phone", u.C(user.FieldPhone)),
		fmt.Sprintf("%s as type", pc.C(partnercontact.FieldType)),
		fmt.Sprintf("%s as partnerID", p.C(partner.FieldID)),
		fmt.Sprintf("%s as partnerName", p.C(partner.FieldName)),
	)
}
