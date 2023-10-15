package partner

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/user"
	"roofix/ent/usersession"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
)

func SetActiveStatus(ctx context.Context, partnerID string, isActive bool) error {
	db := ent.GetClient()
	defer db.CloseClient()

	var action audit.Action
	var status enum.PartnerStatus
	if isActive {
		status = enum.PartnerStatusActive
		action = audit.PartnerActive
	} else {
		status = enum.PartnerStatusInActive
		action = audit.PartnerInActive
	}

	// start tx
	p, err := db.Partner.Query().
		Where(partner.ID(partnerID)).
		WithContacts(func(u *ent.UserQuery) {
			u.Select(user.FieldID)
		}).
		Only(ctx)
	if err != nil {
		return err
	}

	if p.Status == enum.PartnerStatusOnboarding {
		return errors.New("partner status is 'Onboarding'")
	}

	// set Partner active = value
	err = db.Partner.
		UpdateOneID(partnerID).
		SetStatus(status).
		Exec(ctx)
	if err != nil {
		return err
	}

	// if account is InActive then let's kick out all user sessions
	if !isActive {
		_, err = db.UserSession.Delete().Where(func(us *sql.Selector) {
			us.Where(sql.EQ(us.C(usersession.PartnerColumn), partnerID))
		}).Exec(ctx)
		if err != nil {
			log.Error(err)
		}
	}

	audit.OpSuccess(ctx, action, fmt.Sprintf("partnerID: %s is %s", partnerID, status))
	return nil
}
