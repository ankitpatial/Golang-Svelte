/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package job

import (
	"context"
	"roofix/ent"
	"roofix/pkg/account"
)

// RemoveRoofingPartner set new job status, clears Assigned Partner and commit tx, tx will be rolled back on any error
func RemoveRoofingPartner(ctx context.Context, tx *ent.Tx, jobID string, note string) error {
	err := tx.Job.UpdateOneID(jobID).ClearRoofingPartner().Exec(ctx)
	if err != nil {
		return err
	}

	// save status history
	err = NewActivity(ctx, tx, jobID, note)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewActivity(ctx context.Context, tx *ent.Tx, jobID string, note string) error {
	// update status history
	qry := tx.JobActivity.Create().SetJobID(jobID).SetDescription(note)
	if uID := account.CtxUserID(ctx); uID != "" {
		qry.SetCreatorID(uID)
	} else if apiUID := account.CtxApiUserID(ctx); apiUID != "" {
		qry.SetCreatorAPIID(apiUID)
	}

	return qry.Exec(ctx)
}
