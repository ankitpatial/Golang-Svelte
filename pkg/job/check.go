package job

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	entJob "roofix/ent/job"
)

// IsRequester will check if the give job was submitted by the give partner org
func IsRequester(ctx context.Context, jobID, partnerID string) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	return db.Job.Query().
		Where(func(j *sql.Selector) {
			j.Where(sql.And(
				sql.EQ(j.C(entJob.FieldID), jobID),
				sql.EQ(j.C(entJob.RequesterColumn), partnerID),
			))
		}).
		Exist(ctx)
}
