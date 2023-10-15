package job

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/job"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/notification"
	"roofix/pkg/util/log"
)

type NotifyProgressData struct {
	Name    string
	Email   string
	Address string
	Message string
}

func NotifyPartnerAssigned(ctx context.Context, id, partnerID string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelJob,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Roofing Partner Assigned",
		Data: map[string]interface{}{
			"id":        id,
			"partnerID": partnerID,
		},
		Audience: notification.Audience{
			PartnerIDs: []string{
				partnerID, // assigned to partner
			},
		},
	})
}

func NotifyNoteCreated(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelJobNote,
		Topic:   enum.TopicCreated,
		RefID:   id,
		Title:   "Job Note Created",
		Data: map[string]interface{}{
			"id": id,
		},
	})
}
func NotifyNoteUpdated(ctx context.Context, jobID, noteID string) {
	notify(ctx, jobID, &notification.Message{
		Channel: enum.ChannelJobNote,
		Topic:   enum.TopicUpdated,
		RefID:   jobID,
		Title:   "Job Note Edited",
		Data: map[string]interface{}{
			"id":     jobID,
			"noteID": noteID,
		},
	})
}

func NotifyDocsChange(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelJob,
		Topic:   enum.TopicFileUpload,
		RefID:   id,
		Title:   "File Upload",
		Data: map[string]interface{}{
			"id": id,
		},
	})
}

func notify(ctx context.Context, jobID string, msg *notification.Message) {
	// pull
	db := ent.GetClient()
	defer db.CloseClient()

	j, err := db.Job.Query().
		WithCreator(func(c *ent.UserQuery) { c.Select(user.FieldID) }).
		WithRequester(func(p *ent.PartnerQuery) { p.Select(partner.FieldID) }).
		WithHomeOwner().
		Where(job.ID(jobID)).
		Select(job.FieldID).
		Only(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	notifyAll(ctx, j, msg)
}

func notifyAll(ctx context.Context, j *ent.Job, msg *notification.Message) {
	var owner string
	if j.Edges.HomeOwner != nil {
		owner = fmt.Sprintf("%s %s", j.Edges.HomeOwner.FirstName, j.Edges.HomeOwner.LastName)
	}

	address := AddressWithNL(j)
	msg.Data["owner"] = owner
	msg.Data["address"] = address
	msg.Message = fmt.Sprintf("%s\n%s", owner, address)

	// socket connection filters
	to := make([]notification.To, 0, 4)
	// to: ctx user & admin
	to = append(to, notification.ToUser(account.CtxUserID(ctx)), notification.ToAdmins())
	// to: creator of job
	if j.Edges.Creator != nil {
		to = append(to, notification.ToUser(j.Edges.Creator.ID))
	}
	// to: requester of job (Company to whom creator belongs)
	if j.Edges.Requester != nil {
		to = append(to, notification.ToCompanyAdmins(j.Edges.Requester.ID))
	}

	notification.SendAndSave(ctx, msg, to...)
}
