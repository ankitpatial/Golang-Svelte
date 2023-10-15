/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package job

import (
	"context"
	"errors"
	"fmt"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/job"
	entPartner "roofix/ent/partner"
	"roofix/ent/user"
	"roofix/mailer"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/pkg/util/num"
	"roofix/template"
	"strings"
	"time"
)

type DatesInput struct {
	MaterialDate   *time.Time `json:"materialDate"`
	RemoveDate     *time.Time `json:"removeDate"`
	InstallDate    *time.Time `json:"installDate"`
	CompletionDate *time.Time `json:"completionDate"`
}

// AssignPartner a partner to a job and set its status as Assigned
// send email to partner primary contact, ask to accept or reject
// return job:partnerID
func AssignPartner(ctx context.Context, jobID, partnerID string) (string, error) {
	//
	// ROOF-33, There is no accept/deny dashboard or option.  When we assign a job to roofing partner they automatically
	//          received it in their jobs dashboard, they don’t have the option to accept our deny it
	//

	db := ent.GetClient()
	defer db.CloseClient()

	// check Job exist
	j, err := db.Job.Query().WithHomeOwner().Where(job.ID(jobID)).Only(ctx)
	if err != nil {
		log.Error(err)
		return "", msg.AsError(msg.NotFound, "Job")
	}

	// check Partner exists
	partner, err := db.Partner.Query().
		WithContacts(func(c *ent.UserQuery) {
			c.Where(user.StatusEQ(enum.AccountStatusActive))
			c.Select(user.FieldEmail, user.FieldFirstName, user.FieldLastName)
		}).
		Where(entPartner.ID(partnerID)).
		Only(ctx)
	if err != nil {
		log.Error(err)
		return "", msg.AsError(msg.NotFound, "Partner")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	// save Job Dates & Partner info
	now := time.Now().UTC()
	twoBD := addBusinessDays(config.TimeNowAppTZ(), 2)
	progress := enum.JobProgressNew
	err = tx.Job.UpdateOneID(jobID).
		SetRoofingPartnerID(partnerID).
		SetRoofingPartnerAssignedAt(now).
		SetRoofingPartnerFlagAt(twoBD.UTC()).
		SetNillableProgressFlagAt(FlagAtUTC(progress)).
		Exec(ctx)
	if err != nil {
		return "", err
	}

	// job status history
	err = tx.JobActivity.Create().
		SetJobID(jobID).
		SetDescription(fmt.Sprintf("assigned partner %s", partner.Name)).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		return "", err
	}

	// job assignment history
	assign, err := tx.JobAssignmentHistory.Create().
		SetJobID(jobID).
		SetPartnerID(partnerID).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		if ent.IsConstraintError(err) && strings.Contains(err.Error(), "Duplicate entry") {
			return "", errors.New("already had an assignment history")
		}

		return "", err
	}

	// email
	if len(partner.Edges.Contacts) > 0 {
		var to []string
		for _, e := range partner.Edges.Contacts {
			to = append(to, e.Email)
		}

		mailer.Send(ctx, &mailer.Message{
			To:      to,
			Subject: "New Job Assigned",
			Tpl:     template.EmailNewJobAssigned,
			Modal: map[string]interface{}{
				"owner":       j.Edges.HomeOwner.FirstName + " " + j.Edges.HomeOwner.LastName,
				"partnerName": partner.Name,
				"jobAddress":  Address(j),
				"jobPrice":    num.FormatMoney(j.Price),
				"pathname":    fmt.Sprintf("/%s/jobs", strings.ToLower(partner.Type.String())),
			},
		})
	} else {
		log.Warn("no partner contact found to send email")
	}

	// commit tx
	if err := tx.Commit(); err != nil {
		return "", err
	}

	// done!
	return assign.ID, nil
}
