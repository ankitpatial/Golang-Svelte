package installation

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/installationjob"
	"roofix/pkg/account"
	"roofix/pkg/msg"
	"roofix/pkg/util/log"
)

// AssertAdminOrCompanyAdmin job access by checking context user role.
// will assert context-users job access
func AssertAdminOrCompanyAdmin(ctx context.Context, jobID string) error {
	u := account.CtxUser(ctx)
	if !u.IsAdmin && !u.IsCompanyAdmin {
		return msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.InstallationJob.Query().Where(installationjob.ID(jobID))
	// company-ADMIN
	if u.IsCompanyAdmin {
		qry.Where(func(j *sql.Selector) {
			j.Where(sql.EQ(j.C(installationjob.RequestingPartnerColumn), u.Partner.ID))
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
