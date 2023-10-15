package survey

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
	"roofix/server/model"
	"time"
)

func Reserve(ctx context.Context, creatorID string, input model.SurveyInput) (*ent.Survey, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// Validate input data
	if input.Name == "" || input.PhoneNumber == "" || input.Address == "" {
		return nil, errors.New("all fields are required")
	}

	// pull request by ID
	id := input.ID
	s, err := ByIdAndCreatorID(ctx, db, id, creatorID)
	// status must be REQUESTING
	if s.Status != enum.SurveyStatusRequesting {
		return nil, ErrAlreadySubmitted
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, msg.AsError(msg.ServerError)
	}

	// Update existing survey with the new information
	qry := tx.Survey.UpdateOneID(id).
		SetName(input.Name).
		SetAddress(input.Address).
		SetPhone(input.PhoneNumber).
		SetStatus(enum.SurveyStatusRequested).
		SetProgress(enum.SurveyProgressScheduled).
		SetProgressAt(time.Now().UTC()).
		ClearUntil()
	// if note is there, set it
	if input.Notes != nil {
		qry.SetNotes(*input.Notes)
	}

	obj, err := qry.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating survey: %v", err)
	}

	// save entry in history also
	err = tx.SurveyProgress.Create().
		SetStatus(enum.SurveyProgressScheduled).
		SetComplete(true).
		SetNote("").
		SetSurveyID(id).
		SetCreatorID(creatorID).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		log.Error(err)
		return nil, msg.AsError(msg.ServerError)
	}

	// commit TX
	err = tx.Commit()
	if err != nil {
		log.Error(err)
		return nil, msg.AsError(msg.ServerError)
	}

	return obj, nil
}
