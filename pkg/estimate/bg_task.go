package estimate

import (
	"context"
	"roofix/ent"
	"roofix/pkg/nearmap"
	"roofix/pkg/task"
	"roofix/pkg/util/log"
	"roofix/pkg/util/parse"
)

type BGTask struct {
	ID    string `json:"id"`
	Input Input  `json:"input"`
}

func (t BGTask) Name() task.Name {
	return task.Estimate
}

func (t BGTask) Process(ctx context.Context, progress chan string, done chan error) {
	db := ent.GetClient()
	defer db.CloseClient()

	progress <- "material price estimation is in progress"
	est, err := materialPricing(ctx, &t.Input)
	if err != nil {
		markAsFailed(ctx, db, t.ID, err.Error())
		notifyFailed(ctx, t.ID)
		done <- err
		return
	}

	p := est.Price
	// update estimate data
	err = UpdateEstimation(ctx, db, t.ID, &Estimation{
		Estimator:   nearmap.ProviderName,
		SQ:          p.TotalSq,
		Pitch:       p.Pitch,
		Price:       p.Amt,
		Summary:     p.Summary,
		RawResponse: parse.StructToMap(est.Response),
		Bounds:      est.Boundary,
	})
	if err != nil {
		log.Error(err)
		done <- err
		return
	}

	// estimate.pdf
	genPdf(ctx, t.ID, &t.Input, p)
	// notifications
	notifyCompleted(ctx, t.ID)

	progress <- "successfully completed material price estimation"
	done <- nil
}
