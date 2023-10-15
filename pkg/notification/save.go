package notification

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/channel"
	"roofix/pkg/enum"
	"roofix/pkg/util/crypt"
)

func createWithSubscribers(ctx context.Context, db *ent.Client, ms *Message) (string, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	var exist bool
	// create notification channel
	id := crypt.MD5Hash(fmt.Sprintf("%s%s%s", ms.Channel, ms.Topic, ms.RefID))
	exist, err = tx.Channel.Query().Where(channel.ID(id)).Exist(ctx)
	if err != nil {
		return "", err
	}
	// exists
	if exist {
		return id, nil
	}

	// create a notification channel
	_, err = tx.Channel.Create().SetID(id).SetName(ms.Channel).SetTopic(ms.Topic).SetRefID(ms.RefID).Save(ctx)
	if err != nil {
		return "", err
	}

	// add subscribers
	bulk := make([]*ent.ChannelSubCreate, 0, len(ms.Audience.UserIDs)+len(ms.Audience.PartnerIDs)+1)
	for _, uID := range ms.Audience.UserIDs {
		if uID == "" {
			continue
		}
		bulk = append(bulk, tx.ChannelSub.Create().SetChannelID(id).SetUserID(uID))
	}
	for _, pID := range ms.Audience.PartnerIDs {
		bulk = append(bulk, tx.ChannelSub.Create().SetChannelID(id).SetPartnerID(pID))
	}

	if ms.Audience.Admins {
		bulk = append(bulk, tx.ChannelSub.Create().SetChannelID(id).SetRole(enum.RoleAdmin))
	}

	_, err = tx.ChannelSub.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	return id, tx.Commit()
}
