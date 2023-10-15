package job

import (
	"context"
	"fmt"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/job"
	"roofix/ent/jobprogresshistory"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/notification"
	"roofix/pkg/util/log"
	"roofix/server/model"
	"time"
)

type StepLimit struct {
	step enum.JobProgress
	days uint
}

var progressLimits = []StepLimit{
	{enum.JobProgressCustomerContacted, 1},
	{enum.JobProgressJobDetailsConfirmed, 3},
	{enum.JobProgressPermitting, 1},
	{enum.JobProgressScheduled, 3},
	{enum.JobProgressInProgress, 1},
	{enum.JobProgressInstalled, 2},
	{enum.JobProgressInvoiced, 1},
}

func nextStep(progress enum.JobProgress) *enum.JobProgress {
	var ns enum.JobProgress
	switch progress {
	case enum.JobProgressNew:
		ns = enum.JobProgressCustomerContacted
		return &ns
	case enum.JobProgressCustomerContacted:
		ns = enum.JobProgressJobDetailsConfirmed
		return &ns
	case enum.JobProgressJobDetailsConfirmed:
		ns = enum.JobProgressPermitting
		return &ns
	case enum.JobProgressPermitting:
		ns = enum.JobProgressScheduled
		return &ns
	case enum.JobProgressScheduled:
		ns = enum.JobProgressInProgress
		return &ns
	case enum.JobProgressInProgress:
		ns = enum.JobProgressInstalled
		return &ns
	case enum.JobProgressInstalled:
		ns = enum.JobProgressInvoiced
		return &ns
	}

	return nil
}

func progressStepName(progress enum.JobProgress) string {
	switch progress {
	case enum.JobProgressNew:
		return "New"
	case enum.JobProgressCustomerContacted:
		return "Customer Contacted"
	case enum.JobProgressJobDetailsConfirmed:
		return "Job Details Confirmed"
	case enum.JobProgressPermitting:
		return "Permitting"
	case enum.JobProgressScheduled:
		return "Scheduled"
	case enum.JobProgressInProgress:
		return "In Progress"
	case enum.JobProgressInstalled:
		return "Installed"
	}

	return ""
}

func ProgressUpdate(
	ctx context.Context, id string, p enum.JobProgress, stepComplete bool, note string, data *model.ProgressInput,
) error {
	db := ent.GetClient()
	defer db.CloseClient()

	// fetch job
	// to check current progress status
	j := withCreatorAndRequester(ctx, id)
	if j == nil {
		return msg.AsError(msg.FailedToFind, "Job")
	}

	// do not change progress if status is Delayed
	if j.Progress != nil && *j.Progress == enum.JobProgressDelayed {
		return ErrCanNotUpdateDelayedJobProgress
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}

	var jUpdateQry *ent.JobUpdateOne

	// on step complete, set job progress
	if stepComplete {
		jUpdateQry = tx.Job.UpdateOneID(id)
		jUpdateQry.
			SetProgress(p).
			SetProgressAt(time.Now().UTC())

		if flagAt := FlagAtUTC(p); flagAt != nil {
			jUpdateQry.SetNillableProgressFlagAt(flagAt) // set flag at
		} else {
			jUpdateQry.ClearProgressFlagAt() // clear flag at
		}
	}

	// job data
	if data != nil {
		if jUpdateQry == nil {
			jUpdateQry = tx.Job.UpdateOneID(id)
		}

		if data.ShingleColor != nil {
			jUpdateQry.SetShingleColor(*data.ShingleColor)
		}
		if data.InstallDate != nil {
			jUpdateQry.SetInstallDate(config.TimeInAppTZ(*data.InstallDate).UTC())
		}

		if data.InspectionDate != nil {
			jUpdateQry.SetInspectionDate(config.TimeInAppTZ(*data.InspectionDate).UTC())
		}

		if data.CompletionDate != nil {
			jUpdateQry.SetCompletionDate(config.TimeInAppTZ(*data.CompletionDate).UTC())
		}

		jUpdateQry.SetNillablePermitRequired(data.PermitRequired)
		jUpdateQry.SetNillableInspectionRequired(data.InspectionRequired)
	}

	// save progress status
	if err = saveProgress(ctx, tx, jUpdateQry, id, p, note, stepComplete); err != nil {
		return err // error on save
	}

	// ws & push notifications
	notifyAll(ctx, j, &notification.Message{
		Channel: enum.ChannelJob,
		Topic:   enum.TopicProgress,
		RefID:   id,
		Title:   "Job Progress Update",
		Data: map[string]interface{}{
			"id":      j.ID,
			"status":  p,
			"message": fmt.Sprintf("Completed step %s", progressStepName(p)),
		},
	})

	return nil
}

func ProgressDelay(ctx context.Context, id string, flag bool, reason string) error {
	delayed := enum.JobProgressDelayed

	// fetch job
	// to check current progress status
	j := withCreatorAndRequester(ctx, id)
	if j == nil {
		return msg.AsError(msg.FailedToFind, "Job")
	}
	//
	// flag = true, mark job as Delayed
	// flag = false, remove job delay
	//
	if flag && j.Progress != nil && *j.Progress == delayed {
		//return, job already in Delayed status
		log.Warn("ignored attempt to set job(%s) progress delay, because job is already in 'Delayed' progress status", j.ID)
		return nil
	}

	if !flag && j.Progress != nil && *j.Progress != delayed {
		//return, job is not in Delayed status
		log.Warn("ignored attempt to remove job(%s) delay, because job is ont in 'Delayed' progress status", j.ID)
		return nil
	}

	db := ent.GetClient()
	// fetch last completed step from job-progress-history
	// this will be used get the active job progress status
	last, err := db.JobProgressHistory.Query().
		Where(
			jobprogresshistory.HasJobWith(job.ID(id)),
			jobprogresshistory.Complete(true),
			jobprogresshistory.StatusNotIn(enum.JobProgressDelayed),
		).
		Order(ent.Desc(jobprogresshistory.FieldCreatedAt)).
		First(ctx)
	if err != nil && !ent.IsNotFound(err) { // unexpected error
		return err
	}

	// last progress step
	var ls enum.JobProgress
	if last != nil {
		ls = last.Status
	} else { // no record found from history
		ls = enum.JobProgressNew
	}

	// return if next progress step is not conclusive
	var ns *enum.JobProgress
	if ns = nextStep(ls); ns == nil {
		return ErrProgressUnknown
	}

	// begin TX
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}

	var jUpdateQry *ent.JobUpdateOne
	var note string
	var p enum.JobProgress
	var pMsg string
	if flag {
		// set job progress is delayed
		jUpdateQry = tx.Job.UpdateOneID(id).
			SetProgress(delayed).
			SetProgressAt(time.Now().UTC()).
			ClearProgressFlagAt()

		p = enum.JobProgressDelayed
		note = fmt.Sprintf("Delayed, %s", reason)
		pMsg = fmt.Sprintf("Delayed step %q reason %q", progressStepName(*ns), reason)
	} else {
		// set job progress as it was before delay
		// with new time stamp & new flag at
		jUpdateQry = tx.Job.UpdateOneID(id).
			SetProgress(ls).
			SetProgressAt(time.Now().UTC()).
			SetNillableProgressFlagAt(FlagAtUTC(ls))

		p = *ns
		note = fmt.Sprintf("Resumed, %s", reason)
		pMsg = fmt.Sprintf("Resumed step %q reason %q", progressStepName(*ns), reason)
	}

	if err = saveProgress(ctx, tx, jUpdateQry, id, p, note, flag); err != nil {
		return err
	}

	// ws & push notifications
	notifyAll(ctx, j, &notification.Message{
		Channel: enum.ChannelJob,
		Topic:   enum.TopicProgress,
		RefID:   id,
		Title:   "Job Progress Delay",
		Data: map[string]interface{}{
			"id":        j.ID,
			"status":    *ns,
			"message":   pMsg,
			"isDelayed": flag,
		},
	})
	return nil
}

// saveProgress will save data to Job & JobProgressHistory
func saveProgress(
	ctx context.Context, tx *ent.Tx, updateJob *ent.JobUpdateOne,
	id string, p enum.JobProgress, note string, complete bool,
) error {
	// save Job Data
	if updateJob != nil {
		if err := updateJob.Exec(ctx); err != nil {
			return err
		}
	}

	// save Progress History
	if err := saveJobProgressHistory(ctx, tx, id, p, note, complete); err != nil {
		return err
	}

	// SAVE ALL CHANGES
	return tx.Commit()
}

func saveJobProgressHistory(ctx context.Context, tx *ent.Tx, jobID string, p enum.JobProgress, note string, isComplete bool) error {
	// save Progress History
	err := tx.JobProgressHistory.Create().
		SetJobID(jobID).
		SetStatus(p).
		SetNote(note).
		SetComplete(isComplete).
		SetCreatorID(account.CtxUserID(ctx)).
		Exec(ctx)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
