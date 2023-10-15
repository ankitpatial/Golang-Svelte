package estimate

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
)

func UpdateEstimation(
	ctx context.Context, db *ent.Client, id string, inp *Estimation,
) error {
	// update order
	err := db.Estimate.UpdateOneID(id).
		SetEstimator(inp.Estimator).
		SetStatus(enum.EstimateStatusPending).
		SetPrice(inp.Price).
		SetPrimaryPitch(inp.Pitch).
		SetTotalSquares(inp.SQ).
		SetPriceSummary(inp.Summary).
		SetBounds(inp.Bounds).
		SetEstimatorRawResponse(inp.RawResponse).
		Exec(ctx)

	if err != nil {
		return err
	}
	return err
}

// markAsFailed set status to failed with failure reason
func markAsFailed(ctx context.Context, db *ent.Client, id, reason string) {
	// update order
	err := db.Estimate.UpdateOneID(id).SetStatus(enum.EstimateStatusFailed).SetFailureReason(reason).Exec(ctx)
	if err != nil {
		log.Error(err)
	}
}
