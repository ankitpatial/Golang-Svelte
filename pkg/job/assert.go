package job

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/job"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
)

// AssertAccess for context user
// Admin will have access by default
// Site_User (Linked Partner) must have job assigned
func AssertAccess(ctx context.Context, jobID string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Job.Query().Where(job.ID(jobID))
	u := account.CtxUser(ctx)

	// none ADMIN check
	if !u.IsAdmin {
		qry.Where(func(j *sql.Selector) {
			// creator filter
			pred := []*sql.Predicate{
				sql.EQ(j.C(job.CreatorColumn), u.ID),
				sql.EQ(j.C(job.SalesRepColumn), u.ID),
			}
			// linked partner acc info
			p := u.Partner
			if p != nil {
				// assigned to filter
				if p.Type == enum.PartnerRoofing {
					pred = append(pred, sql.EQ(j.C(job.RoofingPartnerColumn), p.ID))
				}
				//
				// partner-admin filter
				if u.IsCompanyAdmin {
					pred = append(pred, sql.EQ(j.C(job.RequesterColumn), p.ID))
				}
			}

			j.Where(sql.Or(pred...))
		})
	}

	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.JobNotAssigned)
	}

	return nil
}

// AssertAdminOrCompanyAdmin job access by checking context user role.
// will assert context-users job access
func AssertAdminOrCompanyAdmin(ctx context.Context, jobID string) error {
	u := account.CtxUser(ctx)
	if !u.IsAdmin && !u.IsCompanyAdmin {
		return msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Job.Query().Where(job.ID(jobID))

	// company-ADMIN
	if u.IsCompanyAdmin {
		qry.Where(func(j *sql.Selector) {
			j.Where(sql.EQ(j.C(job.RequesterColumn), u.Partner.ID))
		})
	}

	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.JobNotAssigned)
	}

	return nil
}

func AssertAdminOrAssigned(ctx context.Context, jobID string) error {
	u := account.CtxUser(ctx)
	if u == nil || (!u.IsAdmin && !u.IsCompanyAdmin) {
		return msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Job.Query().Where(job.ID(jobID))
	if u.IsCompanyAdmin { // assigned partner filter
		qry.Where(func(j *sql.Selector) {
			j.Where(sql.EQ(j.C(job.RoofingPartnerColumn), u.Partner.ID))
		})
	}

	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.JobNotAssigned)
	}

	return nil
}

func AssertAssigned(ctx context.Context, jobID string) error {
	u := account.CtxUser(ctx)
	if u == nil || !u.IsCompanyAdmin {
		return msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()
	// assigned partner filter
	qry := db.Job.Query().
		Where(job.ID(jobID)).
		Where(func(j *sql.Selector) {
			j.Where(sql.EQ(j.C(job.RoofingPartnerColumn), u.Partner.ID))
		})
	exist, err := qry.Exist(ctx)
	if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	} else if !exist {
		return msg.AsError(msg.JobNotAssigned)
	}

	return nil
}
