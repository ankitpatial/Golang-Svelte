package estimate

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/estimate"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

func Deny(ctx context.Context, input *model.DenyEstimateInput, userID, apiUserID *string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	// pull estimate and make sure it's in status is pending
	est, err := db.Estimate.Query().Where(estimate.ID(input.ID)).Select(estimate.FieldID, estimate.FieldStatus).Only(ctx)
	if err != nil {
		return err
	} else if est.Status != enum.EstimateStatusPending {
		return fmt.Errorf("estimate status is not Pending")
	}

	// begin transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// update estimate status
	estStatus := enum.EstimateStatusDenied
	err = tx.Estimate.UpdateOneID(input.ID).SetStatus(estStatus).Exec(ctx)
	if err != nil {
		return err
	}

	// create estimate activity
	err = tx.EstimateActivity.Create().
		SetEstimateID(input.ID).
		SetDescription(fmt.Sprintf("changed status to %s, reason: %s", estStatus, str.Val(input.Note))).
		SetNillableCreatorID(userID).
		SetNillableCreatorAPIID(apiUserID).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func UndoDeny(ctx context.Context, id, note string, userID, apiUserID *string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	// pull estimate and make sure it's in status is denied
	est, err := db.Estimate.Query().Where(estimate.ID(id)).Select(estimate.FieldID, estimate.FieldStatus).Only(ctx)
	if err != nil {
		return err
	} else if est.Status != enum.EstimateStatusDenied {
		return fmt.Errorf("estimate status is not Denied")
	}

	// begin transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// update estimate status
	estStatus := enum.EstimateStatusPending
	err = tx.Estimate.UpdateOneID(id).SetStatus(estStatus).Exec(ctx)
	if err != nil {
		return err
	}

	// create estimate activity
	err = tx.EstimateActivity.Create().
		SetEstimateID(id).
		SetDescription(note).
		SetNillableCreatorID(userID).
		SetNillableCreatorAPIID(apiUserID).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
