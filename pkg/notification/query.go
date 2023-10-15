package notification

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/ent"
	"roofix/ent/channel"
	"roofix/ent/channelmessage"
	"roofix/ent/channelmessageread"
	"roofix/ent/channelsub"
	"roofix/pkg/enum"
	"roofix/pkg/model"
)

func UnreadCount(ctx context.Context, uID string, partnerID *string, isAdmin bool) (int, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	return messageQuery(db, uID, partnerID, isAdmin, true, false).Count(ctx)
}

func Messages(uID string, partnerID *string, isAdmin bool, after *ent.Cursor, first *int) (*model.NotifyMessageConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	order := &ent.ChannelMessageOrder{
		Direction: entgql.OrderDirectionDesc,
		Field:     ent.ChannelMessageOrderFieldCreatedAt,
	}
	p, err := messageQuery(db, uID, partnerID, isAdmin, false, true).
		WithChannel(func(ch *ent.ChannelQuery) {
			ch.Select(channel.FieldID, channel.FieldName, channel.FieldTopic, channel.FieldRefID)
		}).
		// using ctx is not pulling all channel fields
		// using context.Background is
		Paginate(context.Background(), after, first, nil, nil, ent.WithChannelMessageOrder(order))

	if err != nil {
		return nil, err
	}

	res := &model.NotifyMessageConnection{
		TotalCount: p.TotalCount,
		PageInfo:   &p.PageInfo,
	}

	for _, e := range p.Edges {
		n := e.Node
		ch := n.Edges.Channel
		res.Edges = append(res.Edges, &model.NotifyMessageEdge{
			Cursor: e.Cursor,
			Node: &model.NotifyMessage{
				ID:        n.ID,
				Channel:   ch.Name,
				Topic:     ch.Topic,
				RefID:     &ch.RefID,
				Title:     n.Title,
				Message:   n.Message,
				From:      n.FromName,
				Unread:    n.Unread,
				CreatedAt: n.CreatedAt,
			},
		})
	}

	return res, nil
}

func messageQuery(db *ent.Client, uID string, partnerID *string, isAdmin, unread, selectFields bool) *ent.ChannelMessageQuery {
	return db.ChannelMessage.Query().
		Where(func(cm *sql.Selector) {
			subs := sql.Table(channelsub.Table)
			cmr := sql.Table(channelmessageread.Table)
			cm.Join(subs).On(cm.C(channelmessage.ChannelColumn), subs.C(channelsub.ChannelColumn))
			cm.LeftJoin(cmr).On(cm.C(channelmessage.FieldID), cmr.C(channelmessageread.ChannelMessageColumn))

			preds := []*sql.Predicate{
				sql.EQ(subs.C(channelsub.UserColumn), uID),
			}
			if partnerID != nil {
				preds = append(preds, sql.EQ(subs.C(channelsub.PartnerColumn), *partnerID))
			}
			if isAdmin {
				preds = append(preds, sql.EQ(subs.C(channelsub.FieldRole), enum.RoleAdmin))
			}
			cm.Where(sql.Or(preds...))

			if unread {
				cm.Where(sql.EQ(fmt.Sprintf("IFNULL(%s, 0)", cmr.C(channelmessageread.FieldRead)), 0))
			}

			if selectFields {
				ch := sql.Table(channel.Table)
				cm.Join(ch).On(ch.C(channel.FieldID), subs.C(channelsub.ChannelColumn))
				cm.Select(
					cm.C(channelmessage.FieldID),
					cm.C(channelmessage.FieldCreatedAt),
					cm.C(channelmessage.ChannelColumn),
					channelmessage.FieldTitle,
					channelmessage.FieldMessage,
					channelmessage.FieldFromName,
					fmt.Sprintf("!IFNULL(%s, 0) as unread", cmr.C(channelmessageread.FieldRead)),
				)
			}
		}).
		Unique(true)
}
