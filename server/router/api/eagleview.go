/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"roofix/ent"
	"roofix/ent/estimate"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/eagleview"
	"roofix/pkg/util/log"
	"roofix/server/model"
	"roofix/server/router/response"
	"roofix/template"
	"strconv"
)

// eagleViewOrderStatusUpdate endpoint to receive estimate updates
// ref: https://restdoc.e	agleview.com/#OrderStatusUpdate
func eagleViewOrderStatusUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.EstEagleViewUpdate

	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.ServerError(w, err)
		return
	}
	// ensure all needed params are there.
	if !q.Has("StatusId") || !q.Has("SubStatusId") || !q.Has("RefId") || !q.Has("ReportId") {
		err = errors.New(fmt.Sprintf("required param is missing, make sure query has params: StatusId, SubStatusId, RefId & ReportId"))
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, err)
		return
	}

	sid, err := strconv.Atoi(q.Get("StatusId")) // integer
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, errors.New("failed to parse StatusId"))
		return
	}
	ssid, err := strconv.Atoi(q.Get("SubStatusId")) // integer
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, errors.New("failed to parse SubStatusId"))
		return
	}
	refID := q.Get("RefId") // string, UniqueID that was passed as a reference
	//repID, err := strconv.Atoi(q.Get("ReportId")) // integer, EagleView's UniqueID for the order.
	_, err = strconv.Atoi(q.Get("ReportId")) // integer, EagleView's UniqueID for the order.
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, errors.New("failed to parse ReportId"))
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	//check if we have referring Estimate in DB
	est := refExists(ctx, db, w, refID, action)
	if est == nil {
		return
	}

	status := eagleview.Status(uint8(sid))
	subStatus := eagleview.SubStatus(uint8(ssid))
	/* refactoring needed
	   err = db.EstimateResponse.Create().
	   	SetEstimateID(refID).
	   	SetDescription(fmt.Sprintf("Status: %s,  SubStatus: %s", status.String(), subStatus.String())).
	   	SetRaw(map[string]interface{}{
	   		"StatusId":    sid,
	   		"SubStatusId": ssid,
	   		"ReportId":    repID,
	   	}).
	   	Exec(ctx)
	   if err != nil {
	   	audit.OpFailed(ctx, action, err)
	   	response.BadRequest(w, errors.New("failed to save info"))
	   	return
	   }
	*/

	// on complete, trigger get report task
	if eagleview.StatusCompleted == status &&
		(eagleview.SubStatusSent == subStatus || eagleview.SubStatusSent20 == subStatus) {
		if err := eagleview.SubmitGetReport(ctx, refID); err != nil {
			log.Error(err)
		}
	}

	// trigger email
	notifyEagleViewEstChange(ctx, est, status.String(), subStatus.String())

	// done!
	response.Ok(w, "Success")
}

// eagleViewNeedToId endpoint to handle NeedToId
// If the report is NPA (Not Parcel Accurate) then EagleView can notify an integration partner's hosted endpoint.
// ref: https://restdoc.eagleview.com/#OrderStatusUpdate
func eagleViewNeedToId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	action := audit.EstEagleViewUpdate

	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.ServerError(w, err)
		return
	}
	// ensure all needed params are there.
	if !q.Has("RefId") || !q.Has("ReportId") || !q.Has("VerifyId") {
		err = errors.New(fmt.Sprintf("required param is missing, make sure query has params: RefId, ReportId & VerifyId"))
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, err)
		return
	}

	refID := q.Get("RefId")    // string, UniqueID that was passed as a reference
	verID := q.Get("VerifyId") // string, URL Encoded NPA Link.
	//repID, err := strconv.Atoi(q.Get("ReportId")) // integer, EagleView's UniqueID for the order.
	_, err = strconv.Atoi(q.Get("ReportId")) // integer, EagleView's UniqueID for the order.
	if err != nil {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, errors.New("failed to parse ReportId"))
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// check that Estimate exists in database
	est := refExists(ctx, db, w, refID, action)
	if est == nil {
		return
	}

	/* refactoring needed
	   err = db.EstimateResponse.Create().
	   	SetEstimateID(refID).
	   	SetDescription(verID).
	   	SetRaw(map[string]interface{}{
	   		"VerifyId": verID,
	   		"ReportId": repID,
	   	}).
	   	SetNeed(true).
	   	Exec(ctx)

	   if err != nil {
	   	audit.OpFailed(ctx, action, err)
	   	response.BadRequest(w, errors.New("failed to save info"))
	   	return
	   }
	*/
	notifyEagleViewEstNeed(ctx, est, verID)

	response.Ok(w, "Success")
}

func refExists(ctx context.Context, db *ent.Client, w http.ResponseWriter, refID string, action audit.Action) *ent.Estimate {
	// check estimate req exists for EagleView
	est, err := db.Estimate.
		Query().
		WithHomeOwner().
		Where(estimate.ID(refID)).
		First(ctx)

	if err != nil && ent.IsNotFound(err) {
		audit.OpFailed(ctx, action, err)
		response.BadRequest(w, errors.New("invalid RefId"))
		return nil
	} else if err != nil {
		audit.OpFailed(ctx, action, err)
		response.ServerError(w, errors.New("failed to find RefId"))
		return nil
	}

	return est
}

func notifyEagleViewEstChange(ctx context.Context, est *ent.Estimate, status, subStatus string) {
	owner := est.Edges.HomeOwner
	msg := &mailer.Message{
		To:      account.ToNotifyEmails(ctx, model.AdminNotifyTopicEagleViewEstimateProgress.String()),
		Subject: "EagleView Estimate Progress",
		Tpl:     template.EmailEagleViewEstChange,
		Modal: map[string]interface{}{
			"owner":        fmt.Sprintf("%s %s", owner.FirstName, owner.LastName),
			"streetNumber": owner.StreetNumber,
			"streetName":   owner.StreetName,
			"city":         owner.City,
			"state":        owner.State,
			"zip":          owner.Zip,
			"status":       status,
			"subStatus":    subStatus,
			"pathname":     fmt.Sprintf("/estimates/%s", est.ID),
		},
	}
	mailer.Send(ctx, msg)
}

func notifyEagleViewEstNeed(ctx context.Context, est *ent.Estimate, needUrl string) {
	owner := est.Edges.HomeOwner
	msg := &mailer.Message{
		To:      account.ToNotifyEmails(ctx, model.AdminNotifyTopicEagleViewEstimateProgress.String()),
		Subject: "EagleView Estimate Progress",
		Tpl:     template.EmailEagleViewEstNeed,
		Modal: map[string]interface{}{
			"owner":        fmt.Sprintf("%s %s", owner.FirstName, owner.LastName),
			"streetNumber": owner.StreetNumber,
			"streetName":   owner.StreetName,
			"city":         owner.City,
			"state":        owner.State,
			"zip":          owner.Zip,
			"needURL":      needUrl,
			"pathname":     fmt.Sprintf("/estimates/%s", est.ID),
		},
	}
	mailer.Send(ctx, msg)
}
