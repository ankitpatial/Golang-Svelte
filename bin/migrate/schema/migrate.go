package schema

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"roofix/ent"
	"roofix/ent/job"
	"roofix/ent/user"
	"roofix/pkg/enum"
	pkgJob "roofix/pkg/job"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
)

func Migrate(ctx context.Context) error {
	cl := ent.GetClient()
	return cl.Schema.Create(ctx,
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	)
}

func SaveAdminUsers(ctx context.Context, admins []ent.User) error {
	cl := ent.GetClient()
	adminRole := enum.RoleAdmin
	activeStatus := enum.AccountStatusActive

	for _, u := range admins {
		log.Info("==  %s ==", u.Email)
		oldUser, err := cl.User.Query().Where(user.EmailEQ(u.Email)).First(ctx)
		if err != nil && !ent.IsNotFound(err) { // unknown error
			log.Error(err)
			continue
		}

		if oldUser != nil { // existing user
			log.Info("user with email %s already exists", oldUser.Email)
			if oldUser.Role != adminRole {
				_, err := oldUser.Update().
					SetRole(adminRole).
					SetEmailVerified(false).
					SetStatus(activeStatus).
					Save(ctx)
				if err != nil {
					log.Error(err)
				}
			}
			continue
		}

		// no existing user, create new
		_, err = cl.User.Create().
			SetEmail(u.Email).
			SetPwd(crypt.Hash(u.Pwd)).
			SetEmailVerified(true).
			SetFirstName(u.FirstName).
			SetLastName(u.LastName).
			SetRole(adminRole).
			SetStatus(activeStatus).
			Save(ctx)

		if err != nil {
			log.Error(err)
			continue
		}

		log.Info("created user %s", u.Email)
	}

	return nil
}

func MigrateJobPrice(ctx context.Context) {
	log.Info("migrating job price")
	db := ent.GetClient()
	defer db.CloseClient()

	all, err := db.Job.Query().Where(job.PriceGT(0)).All(ctx)
	if err != nil {
		log.PrintError(err)
	}

	for _, j := range all {
		err = j.Update().
			SetContractPrice(j.Price).
			SetWorkOrderPrice(j.Price * pkgJob.PartnerPriceMultiplyFactor).
			Exec(ctx)
		if err != nil {
			log.PrintError(err)
		}

	}
}
