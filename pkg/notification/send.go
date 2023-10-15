package notification

import (
	"context"
	"roofix/config"
	"roofix/pkg/enum"
	"roofix/pkg/util/app"
	"roofix/pkg/util/log"
	"roofix/pkg/util/queue"
	"roofix/pkg/util/slice"
)

type Message struct {
	Channel  enum.Channel           `json:"channel"`
	Topic    enum.Topic             `json:"topic"`
	RefID    string                 `json:"refID"`
	Title    string                 `json:"title"`
	Message  string                 `json:"message"`
	Data     map[string]interface{} `json:"data"`
	Audience Audience               `json:"audience"`
	SaveToDB bool                   `json:"SaveToDB"`
}

type Audience struct {
	UserIDs    []string
	PartnerIDs []string
	Admins     bool
}

type Company struct {
	ID   string
	Name string
}

type To func(o *Audience)

func Send(ctx context.Context, msg *Message, to ...To) {
	log.InfoBullet("notification.Send()")
	send(ctx, msg, false, to...)
}

func SendAndSave(ctx context.Context, msg *Message, to ...To) {
	log.InfoBullet("notification.SendAndSave()")
	send(ctx, msg, true, to...)
}

func send(ctx context.Context, msg *Message, saveToDB bool, to ...To) {
	log.InfoBullet("notification.send()")

	msg.SaveToDB = saveToDB
	// set audience
	for _, a := range to {
		a(&msg.Audience)
	}

	// push to queue
	if config.IsDevEnv() {
		if err := Process(app.SetCtx(context.Background(), app.ReadCtx(ctx)), msg); err != nil {
			log.Error(err)
		}
		return
	}

	// enqueue message
	if err := queue.Send(ctx, config.NotificationQueueName(), msg); err != nil {
		log.Error(err)
	}
}

func ToUser(ids ...string) To {
	return func(m *Audience) {
		m.UserIDs = slice.Unique(append(m.UserIDs, ids...))
	}
}

func ToCompanyAdmins(partnerIDs ...string) To {
	return func(m *Audience) {
		m.PartnerIDs = slice.Unique(append(m.PartnerIDs, partnerIDs...))
	}
}

func ToAdmins() To {
	return func(m *Audience) {
		m.Admins = true
	}
}
