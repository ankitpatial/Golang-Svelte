/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"roofix/ent"
	ej "roofix/ent/job"
	"roofix/ent/jobactivity"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/estimate"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/server/router/response"
	"time"
)

type jobDetail struct {
	Status        enum.JobProgress `json:"status,omitempty"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	StatusHistory []*history       `json:"statusHistory"`
}
type history struct {
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// getJob will return the job status
func getJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.JobDetail
	apiUID := account.CtxApiUserID(ctx)

	// make sure we JobID
	id := chi.URLParam(r, "id")
	if id == "" {
		e := errors.New("param jobID is missing")
		audit.ApiOpFailed(ctx, action, apiUID, e)
		response.BadRequest(w, e)
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// get job
	j, err := db.Job.Query().
		Select(ej.FieldProgress, ej.FieldCreatedAt, ej.FieldUpdatedAt).
		Where(ej.ID(id)).
		WithActivities(func(a *ent.JobActivityQuery) {
			a.Order(ent.Desc(jobactivity.FieldCreatedAt))
		}).
		WithCreatorAPI().
		First(ctx)

	// server error
	if err != nil && ent.IsNotFound(err) {
		e := errors.New(fmt.Sprintf("invalid jobID: %s", id))
		audit.ApiOpFailed(ctx, action, apiUID, e)
		response.BadRequest(w, e)
		return
	} else if err != nil {
		log.Error(err)
		audit.ApiOpFailed(ctx, action, apiUID, err)
		response.ServerError(w, msg.AsError(msg.FailedToFind, "Job"))
		return
	}

	// make sure Job is related to Api User
	creator := j.Edges.CreatorAPI
	if creator == nil || creator.ID != apiUID {
		err := errors.New(fmt.Sprintf("access denied to jobID: %s", id))
		audit.ApiOpFailed(ctx, action, apiUID, err)
		response.BadRequest(w, err)
		return
	}

	// audit log
	audit.ApiOpSuccess(ctx, action, apiUID, "jobID: "+id)

	// prepare response
	var status enum.JobProgress
	var hs []*history
	for _, s := range j.Edges.Activities {
		hs = append(hs, &history{
			Status:    s.Description,
			CreatedAt: s.CreatedAt,
		})
	}

	// return response
	response.Ok(w, &jobDetail{
		Status:        status,
		CreatedAt:     j.CreatedAt,
		StatusHistory: hs,
	})
}

// postJob will accept create job request from external source
func postJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.JobCreate
	apiUID := account.CtxApiUserID(ctx)

	// bind modal
	var inp estimate.Input
	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		audit.ApiOpFailed(ctx, action, apiUID, err)
		response.BadRequest(w, err)
		return
	}

	// TODO: need modifications as per create estimate changes
	response.BadRequest(w, errors.New("not functional"))
	return

	// save to DB
	//id, err := job.Create(ctx, &inp, enum.JobSourceExternal)
	//if err != nil {
	//	audit.ApiOpFailed(ctx, action, apiUID, err)
	//	response.BadRequest(w, err)
	//	return
	//}
	//
	//m := "created jobID: " + id

	// estimate measurements
	// err = nearmap.SubmitEstimate(ctx, id)
	//if err != nil {
	//	log.Error(err)
	//}

	//audit.Operation(ctx, action, m, err)
	//response.Created(w, map[string]string{
	//	"jobId": id,
	//})
}
