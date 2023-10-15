package estimate

import (
	"context"
	"fmt"
	"roofix/ent"
	entEstimate "roofix/ent/estimate"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/notification"
	"roofix/pkg/util/log"
)

func notifyCompleted(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelEstimate,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Estimate Completed",
		Data: map[string]interface{}{
			"id":     id,
			"status": enum.EstimateStatusPending,
		},
	}, true)
}

func notifyFailed(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelEstimate,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Estimate Failed",
		Data: map[string]interface{}{
			"id":     id,
			"status": enum.EstimateStatusFailed,
		},
	}, false)
}

func NotifyApproved(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelEstimate,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Estimate Approved",
		Data: map[string]interface{}{
			"id":     id,
			"status": enum.EstimateStatusApproved,
		},
	}, true)
}

func NotifyDenied(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelEstimate,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Estimate Denied",
		Data: map[string]interface{}{
			"id":     id,
			"status": enum.EstimateStatusDenied,
		},
	}, true)
}

func NotifyEstimateUnDenied(ctx context.Context, id string) {
	notify(ctx, id, &notification.Message{
		Channel: enum.ChannelEstimate,
		Topic:   enum.TopicStatusChange,
		RefID:   id,
		Title:   "Estimate Denial Has Been Reversed",
		Data: map[string]interface{}{
			"id":     id,
			"status": enum.EstimateStatusPending,
		},
	}, false)
}

func notify(ctx context.Context, jobID string, msg *notification.Message, saveToDB bool) {
	// pull
	db := ent.GetClient()
	defer db.CloseClient()

	j, err := db.Estimate.Query().
		WithHomeOwner().
		WithCreator(func(c *ent.UserQuery) {
			c.Select(user.FieldID)
		}).
		WithPartner(func(p *ent.PartnerQuery) {
			p.Select(partner.FieldID)
		}).
		Where(entEstimate.ID(jobID)).
		Select(
			entEstimate.FieldID,
		).
		Only(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	notifyAll(ctx, j, msg, saveToDB)
}

func notifyAll(ctx context.Context, j *ent.Estimate, msg *notification.Message, saveToDB bool) {
	ho := j.Edges.HomeOwner
	owner := fmt.Sprintf("%s %s", ho.FirstName, ho.LastName)
	address := addressWithNL(ho)
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
	if j.Edges.Partner != nil {
		to = append(to, notification.ToCompanyAdmins(j.Edges.Partner.ID))
	}

	if saveToDB {
		notification.SendAndSave(ctx, msg, to...)
	} else {
		notification.Send(ctx, msg, to...)
	}
}

func addressWithNL(j *ent.HomeOwner) string {
	return fmt.Sprintf("%s %s, %s\n%s %s", j.StreetNumber, j.StreetName, j.City, j.State, j.Zip)
}
