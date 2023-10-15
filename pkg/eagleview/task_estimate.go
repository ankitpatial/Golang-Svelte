/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/pkg/audit"
	"roofix/pkg/task"
	"roofix/pkg/util/uid"
)

type TaskEstimate struct {
	JobID string `json:"jobId"`
}

func (t TaskEstimate) Name() task.Name {
	return task.EagleViewEstimate
}

func (t TaskEstimate) Process(ctx context.Context, progress chan string, done chan error) {
	progress <- "processing EagleView estimate request"
	db := ent.GetClient()
	defer db.CloseClient()

	// pull job
	j, err := db.Job.Get(ctx, t.JobID)
	if err != nil {
		done <- err
		return
	}

	done <- estimate(ctx, db, j, progress)
}

func estimate(ctx context.Context, db *ent.Client, job *ent.Job, progress chan string) error {
	action := audit.EstEagleView
	progress <- "submitting estimate request to EagleView API"
	jobID := job.ID
	refID := uid.ULID()
	//res, err := PlaceOrder(ctx, job, refID)
	//if err != nil {
	//	progress <- "failed to submit estimate request to EagleView API"
	//	audit.OpFailed(ctx, action, err)
	//	return err
	//}

	progress <- "successfully submitted estimate request to EagleView API"
	desc := fmt.Sprintf("submitted, jobID: %s refID: %s", jobID, refID)
	//err = pkgJob.CreateEstimate(ctx, db, jobID, refID, ProviderName, res.ID, res.ReportIDs[0])
	//if err != nil {
	//	audit.OpSuccess(ctx, action, desc+" but failed to save estimate date to db")
	//	return err
	//}

	audit.OpSuccess(ctx, action, desc)
	// set status to estimating
	//pkgJob.Estimating(ctx, jobID)
	return nil
}

func RequestEstimate(ctx context.Context, jobID string) error {
	t := &TaskEstimate{JobID: jobID}
	return task.Submit(ctx, t)
}
