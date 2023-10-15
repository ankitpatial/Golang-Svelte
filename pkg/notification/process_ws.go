package notification

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/ent/usersession"
	"roofix/pkg/enum"
	"roofix/pkg/util/app"
	"roofix/pkg/util/convert"
	"roofix/pkg/util/log"
	"roofix/pkg/ws"
)

func wsMessage(ctx context.Context, msg *Message) {
	connIDs, err := wsConnIDs(ctx, msg.Audience)
	if err != nil {
		log.Error(err)
		return
	}

	if len(connIDs) == 0 {
		log.Warn("abort wsMessage, not connIDs list is empty")
		return
	}

	if msg.SaveToDB {
		msg.Data["isNotification"] = true
	}

	if u := app.CtxUser(ctx); u != nil {
		msg.Data["ctxUserID"] = u.ID
	}

	ws.Send(ctx, connIDs, ws.Message{
		Channel: msg.Channel,
		Topic:   msg.Topic,
		Title:   msg.Title,
		Message: msg.Message,
		Data:    msg.Data,
	})
}

func wsConnIDs(ctx context.Context, a Audience) ([]string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// notify socket connections
	qry := db.UserSession.Query().WithSockets().Where(func(sess *sql.Selector) {
		u := sql.Table(user.Table)
		pc := sql.Table(partnercontact.Table)

		sess.Join(u).On(sess.C(usersession.UserColumn), u.C(user.FieldID))
		sess.LeftJoin(pc).On(sess.C(usersession.PartnerContactColumn), pc.C(partnercontact.FieldID))

		var preds []*sql.Predicate

		// admins
		if a.Admins {
			preds = append(preds, sql.EQ(u.C(user.FieldRole), enum.RoleAdmin))
		}

		// userIDs
		if len(a.UserIDs) > 0 {
			preds = append(preds, sql.In(sess.C(usersession.UserColumn), convert.ToAny(a.UserIDs)...))
		}

		// company admins
		if len(a.PartnerIDs) > 0 {
			preds = append(preds, sql.And(
				sql.In(sess.C(usersession.PartnerColumn), convert.ToAny(a.PartnerIDs)...),
				sql.EQ(pc.C(partnercontact.FieldRole), enum.PartnerContactRoleAdmin),
			))
		}

		sess.Where(sql.Or(preds...))
	})

	if res, err := qry.Select(usersession.FieldID).All(ctx); err != nil {
		return nil, err
	} else {
		return ids(res), err
	}
}

func ids(list []*ent.UserSession) []string {
	connIDs := make([]string, 0, len(list))
	for _, r := range list {
		if r.Edges.Sockets == nil {
			continue
		}

		for _, s := range r.Edges.Sockets {
			connIDs = append(connIDs, s.ID)
		}
	}

	return connIDs
}
