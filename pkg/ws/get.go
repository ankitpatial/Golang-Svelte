/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package ws

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/usersession"
	"roofix/pkg/util/convert"
)

// GetUserSocket by UserIDs or app Routes
func GetUserSocket(ctx context.Context, userIDs []string) ([]string, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.UserSession.Query().WithSockets().
		Where(func(sess *sql.Selector) {
			sess.Where(sql.In(sess.C(usersession.UserColumn), convert.ToAny(userIDs)...))
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
