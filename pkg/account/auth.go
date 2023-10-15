package account

import (
	"context"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/validate"
	"strings"
	"time"
)

func Authenticate(ctx context.Context, inp *LoginInput) (*ent.User, error) {
	// validate
	if err := validate.Struct(inp); err != nil {
		return nil, err
	}

	db := ent.GetClient()
	// get user by email
	u, err := db.User.
		Query().
		Where(user.EmailEQ(strings.TrimSpace(inp.Email))).
		Select(append(sessionUserField, user.FieldPwd, user.FieldWrongAttempts, user.FieldLockedUntil)...).
		Only(ctx)
	if ent.IsNotFound(err) {
		log.Warn("login, wrong email %s", inp.Email)
		return nil, msg.AsError(msg.WrongLoginCred)
	} else if err != nil {
		log.Error(err)
		return nil, msg.AsError(msg.FailedToFind, "user")
	}

	// check if account is locked
	now := time.Now().UTC()
	if u.LockedUntil != nil && now.Before(u.LockedUntil.UTC()) {
		return nil, msg.AsError(msg.AccountLockedUntil, config.TimeStrAppTZ(*u.LockedUntil))
	}

	// compare password
	if !crypt.CompareHash(u.Pwd, inp.Password) {
		u.WrongAttempts += 1
		until := time.Now().Add(AccLocDuration).UTC()
		locked := u.WrongAttempts >= MaxWrongAttempts
		q := db.User.UpdateOneID(u.ID)
		if locked {
			q.SetLockedUntil(until)
		}

		if err := q.SetWrongAttempts(u.WrongAttempts).SetWrongAttemptAt(now).Exec(ctx); err != nil {
			log.Error(err)
		} else if locked {
			// TODO: raise account locked email alert

			return nil, msg.AsError(msg.AccountLockedUntil, config.TimeStrAppTZ(until))
		}

		log.Warn("login, wrong password for email %s", u.Email)
		return nil, msg.AsError(msg.WrongLoginCred)
	} else if u.WrongAttempts > 0 {
		err := db.User.UpdateOneID(u.ID).
			ClearWrongAttempts().
			ClearWrongAttemptAt().
			ClearLockedUntil().
			Exec(ctx)
		if err != nil {
			log.Error(err)
		}
	}

	// check account is active
	if u.Status != enum.AccountStatusActive {
		return nil, msg.AsError(msg.AccountNotActive)
	}

	return u, nil
}
