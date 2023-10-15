package installation

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/installationjob"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/notification"
	"roofix/pkg/util/log"
)

func NotifyCreated(ctx context.Context, id string) {
	msg := &notification.Message{
		Channel: enum.ChannelInstallationJob,
		Topic:   enum.TopicCreated,
		Title:   "Installation Job Created",
		Message: "New installation job is created.",
		Data:    map[string]interface{}{},
	}

	notify(ctx, id, msg)
}

func NotifyApproved(ctx context.Context, id, by string) {
	msg := &notification.Message{
		Channel: enum.ChannelInstallationJob,
		Topic:   enum.TopicUpdated,
		Title:   "Installation Job Approved",
		Message: fmt.Sprintf("Installation Job was approved by %s", by),
		Data:    map[string]interface{}{},
	}

	notify(ctx, id, msg)
}

func NotifyDenied(ctx context.Context, id, by string) {
	msg := &notification.Message{
		Channel: enum.ChannelInstallationJob,
		Topic:   enum.TopicUpdated,
		Title:   "Installation Job Denied",
		Message: fmt.Sprintf("Installation Job was denied by %s", by),
		Data:    map[string]interface{}{},
	}

	notify(ctx, id, msg)
}

func NotifyRemoveDeny(ctx context.Context, id, by string) {
	msg := &notification.Message{
		Channel: enum.ChannelInstallationJob,
		Topic:   enum.TopicUpdated,
		Title:   "Installation Job Remove Denied",
		Message: fmt.Sprintf("Installation Job denied status was removed by %s", by),
		Data:    map[string]interface{}{},
	}

	notify(ctx, id, msg)
}

func notify(ctx context.Context, id string, msg *notification.Message) {
	db := ent.GetClient()
	defer db.CloseClient()

	msg.RefID = id
	// pull job
	j, err := db.InstallationJob.Query().
		WithCreator(func(c *ent.UserQuery) {
			c.Select(user.FieldID)
		}).
		WithRequestingPartner(func(p *ent.PartnerQuery) {
			p.Select(partner.FieldID)
		}).
		Where(installationjob.ID(id)).
		Select(
			installationjob.FieldID,
			installationjob.CreatorColumn,
			installationjob.RequestingPartnerColumn,
			installationjob.FieldOwnerName,
			installationjob.FieldOwnerAddress,
			installationjob.FieldApproval,
			installationjob.FieldStatus,
		).
		Only(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	// socket connection filters
	var to []notification.To
	// to: ctx user & admin
	to = append(
		to,
		notification.ToUser(account.CtxUserID(ctx), j.Edges.Creator.ID),
		notification.ToAdmins(),
	)

	if j.Edges.RequestingPartner != nil {
		to = append(to, notification.ToCompanyAdmins(j.Edges.RequestingPartner.ID))
	}

	msg.Data["id"] = id
	msg.Data["owner"] = j.OwnerName
	msg.Data["owner"] = j.OwnerAddress

	notification.SendAndSave(ctx, msg, to...)
}
