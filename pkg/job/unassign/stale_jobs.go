package unassign

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"roofix/config"
	"roofix/ent"
	entJob "roofix/ent/job"
	"roofix/ent/jobprogresshistory"
	entPartner "roofix/ent/partner"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/job"
	"roofix/pkg/util"
	"roofix/pkg/util/log"
	"roofix/server/model"
	"time"
)

type PartnerJobs struct {
	ID   string
	Name string
	Jobs []*JobInfo
}

type JobInfo struct {
	ID          string
	Owner       string
	Address     string
	PartnerName string
}

func StaleJobs(ctx context.Context) error {
	log.Info("== Un Assign Stale Jobs ==")

	now := config.TimeNowAppTZ()
	log.Info(" UTC TimeInAppTZ: %s", now.UTC().Format(time.DateTime))
	log.Info(" Business TimeInAppTZ: %s", now.Format(time.DateTime))

	db := ent.GetClient()
	defer db.CloseClient()

	jobs, err := expiringToday(ctx, db)
	if err != nil {
		return err
	} else if len(jobs) == 0 {
		log.Info("no stale job found")
		return nil
	}

	log.Info("%d job(s) found")
	pjMap := make(map[string][]*JobInfo) // partner job map
	for _, j := range jobs {
		pID := j.Edges.RoofingPartner.ID
		ji := &JobInfo{
			ID:          j.ID,
			Owner:       fmt.Sprintf("%s %s", j.Edges.HomeOwner.FirstName, j.Edges.HomeOwner.LastName),
			Address:     job.Address(j),
			PartnerName: j.Edges.RoofingPartner.Name,
		}
		if p, ok := pjMap[pID]; ok {
			pjMap[pID] = append(p, ji)
		} else {
			pjMap[pID] = []*JobInfo{ji}
		}
	}

	h := now.Hour()
	if h < 10 {
		return willUnAssignToday(ctx, pjMap)
	} else {
		return unAssign(ctx, db, pjMap)
	}
}

// willClearToday will trigger partner notification with list of jobs that will clear by end of day today
func willUnAssignToday(ctx context.Context, pjMap map[string][]*JobInfo) error {
	for p, jobs := range pjMap {
		notifyWillUnAssign(ctx, p, jobs)
	}
	return nil
}

func unAssign(ctx context.Context, db *ent.Client, pjMap map[string][]*JobInfo) error {
	var pj []*PartnerJobs
	expiryMsg := "invitation time expired, removed roofing partner"
	for pID, jobs := range pjMap {
		var uaJobs []*JobInfo
		for _, j := range jobs {
			tx, _ := db.BeginTx(ctx, nil)

			// update invite
			err := tx.JobAssignmentHistory.
				Create().
				SetStatus(enum.JobAssignmentStatusExpired).
				SetNote("System Note: " + expiryMsg).
				SetJobID(j.ID).
				SetPartnerID(pID).
				Exec(ctx)
			if err != nil {
				log.Error(err)
				continue
			}

			// update job
			err = job.RemoveRoofingPartner(ctx, tx, j.ID, expiryMsg)
			if err != nil {
				log.Error(err)
				continue
			}
			uaJobs = append(uaJobs, j)
		}

		notifyUnassigned(ctx, pID, uaJobs)
		pj = append(pj, &PartnerJobs{
			ID:   pID,
			Name: jobs[0].PartnerName,
			Jobs: uaJobs,
		})
	}

	// alert admins
	admins := account.ToNotify(ctx, model.AdminNotifyTopicJobInvitationExpired.String())
	for _, a := range admins {
		notifyUnassignedToAdmin(ctx, a, pj)
	}
	return nil
}

func expiringToday(ctx context.Context, db *ent.Client) ([]*ent.Job, error) {
	eod := util.EndOfDay(config.TimeNowAppTZ()).UTC()
	return db.Job.Query().
		WithRoofingPartner(func(p *ent.PartnerQuery) { p.Select(entPartner.FieldName) }).
		WithHomeOwner().
		Where(func(j *sql.Selector) {
			j.Where(sql.And(
				sql.EQ(j.C(entJob.FieldProgress), enum.JobProgressNew),
				sql.EQ(fmt.Sprintf(
					"(select count(jph.id)from %s as jph where jph.%s = %s)",
					jobprogresshistory.Table,
					jobprogresshistory.JobColumn,
					j.C(entJob.FieldID)), 0),
				sql.LTE(j.C(entJob.FieldRoofingPartnerFlagAt), eod),
			))
		}).
		Select(
			entJob.FieldID,
			entJob.FieldPrice,
		).
		All(ctx)
}
