package notification

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/pkg/util/app"
	"roofix/pkg/util/log"
)

func saveToDB(ctx context.Context, msg *Message) {
	db := ent.GetClient()
	defer db.CloseClient()

	log.Info("saving message to DB...")

	// create channel
	channelID, err := createWithSubscribers(ctx, db, msg)
	if err != nil {
		log.Error(err)
		return
	}

	// save message
	qry := db.ChannelMessage.Create().
		SetChannelID(channelID).
		SetTitle(msg.Title).
		SetMessage(msg.Message).
		SetPrivate(false)

	if u := app.CtxUser(ctx); u != nil {
		qry.
			SetFromID(u.ID).
			SetFromName(fmt.Sprintf("%s %s", u.FirstName, u.LastName))
	}

	if api := app.CtxAPIUser(ctx); api != nil {
		if api.Name == "" {
			if r, e1 := db.ApiUser.Get(ctx, api.ID); e1 != nil {
				api.Name = "N/A"
			} else {
				api.Name = r.Username
			}
		}
		qry.
			SetFromAPIUserID(api.ID).
			SetFromName(api.Name)
	}

	m, err := qry.Save(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("saved message(%s) to DB", m.ID)
}
