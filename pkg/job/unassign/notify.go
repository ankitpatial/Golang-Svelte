package unassign

import (
	"context"
	"fmt"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/partner"
	"roofix/pkg/util/log"
	"roofix/template"
)

func notifyWillUnAssign(ctx context.Context, pID string, jobs []*JobInfo) {
	count := len(jobs)
	if count == 0 {
		log.Info("jobs list is empty")
		return
	}

	var subject string
	if count > 1 {
		subject = "These Jobs Needs Your Attention"
	} else {
		subject = "This Job Needs Your Attention"
	}

	pc, err := partner.PrimaryContact(ctx, pID)
	if err != nil {
		log.Error(err)
		return
	}

	// notify
	mailer.Send(ctx, &mailer.Message{
		To:      []string{fmt.Sprintf("%s<%s>", pc.Name, pc.Email)},
		Subject: subject,
		Tpl:     template.EmailJobWillUnassign,
		Modal: map[string]interface{}{
			"partner": jobs[0].PartnerName,
			"jobs":    jobs,
		},
	})
}

func notifyUnassigned(ctx context.Context, pID string, jobs []*JobInfo) {
	count := len(jobs)
	if count == 0 {
		log.Info("jobs list is empty")
		return
	}

	var subject string
	if count > 1 {
		subject = "These Jobs Needs Your Attention"
	} else {
		subject = "This Job Needs Your Attention"
	}

	pc, err := partner.PrimaryContact(ctx, pID)
	if err != nil {
		log.Error(err)
		return
	}

	// notify
	mailer.Send(ctx, &mailer.Message{
		To:      []string{fmt.Sprintf("%s<%s>", pc.Name, pc.Email)},
		Subject: subject,
		Tpl:     template.EmailJobUnassigned,
		Modal: map[string]interface{}{
			"partner": jobs[0].PartnerName,
			"isMulti": count > 1,
			"jobs":    jobs,
		},
	})
}

func notifyUnassignedToAdmin(ctx context.Context, admin *account.EmailAddress, list []*PartnerJobs) {
	count := len(list)
	if count == 0 {
		log.Info("PartnerJobs list is empty")
		return
	}

	// notify
	mailer.Send(ctx, &mailer.Message{
		To:      []string{fmt.Sprintf("%s<%s>", admin.Name, admin.Email)},
		Subject: "Unassigned Job(s)",
		Tpl:     template.EmailJobUnassignedAdmin,
		Modal: map[string]interface{}{
			"name": admin.Name,
			"list": list,
		},
	})
}

//func notifyUnassignedToAdmin(ctx context.Context, j []*ent.Job, admins []*account.EmailAddress) {
//	if j == nil {
//		logger.Warn("job data was nil")
//		return
//	}
//
//	// email admin
//	users := account.ToNotify(ctx, model.AdminNotifyTopicJobInvitationExpired.String())
//	for _, u := range users {
//		mailer.Send(ctx, &mailer.Message{
//			To:      []string{u.Email},
//			Subject: "Job Unassigned",
//			Tpl:     mailer.EmailJobUnassignedAdmin,
//			Modal: map[string]interface{}{
//				"name":    u.Name,
//				"partner": j.Edges.Partner.Name,
//				"owner":   j.OwnerFirstName + " " + j.OwnerLastName,
//				"address": job.AddressWithBr(j),
//			},
//		})
//	}
//
//	// email partner Primary Contact
//}
