package survey

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/survey"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
)

func HasAccess(ctx context.Context, id string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	u := account.CtxUser(ctx)

	// ADMIN
	if u.Role == enum.RoleAdmin {
		return nil
	}

	//
	// CREATOR
	exist, err := db.Survey.Query().
		Where(func(j *sql.Selector) {
			j.Where(sql.And(
				sql.EQ(j.C(survey.FieldID), id),
				sql.EQ(j.C(survey.CreatedByColumn), u.ID),
			))
		}).
		Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}
	// creator of the survey
	if exist {
		return nil
	}

	//
	// PARTNER
	p := u.Partner
	if p != nil || !p.IsAdmin() {
		return msg.AsError(msg.NotAuthorized)
	}

	exist, err = db.Survey.Query().
		Where(func(j *sql.Selector) {
			j.Where(sql.And(
				sql.EQ(j.C(survey.FieldID), id),
				sql.EQ(j.C(survey.PartnerColumn), p.ID),
			))
		}).
		Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.NotFound)
	}

	return nil
}
