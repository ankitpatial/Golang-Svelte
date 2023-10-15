package partner

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
	"roofix/pkg/util/uid"
)

// QuickAddSalesRep to a partner account
// will not be used in normal flow, but if external estimate request comes that has sales rep info that do not exist yet
// then we need to create the sales rep user account and link it to the partner
func QuickAddSalesRep(ctx context.Context, db *ent.Client, partnerID, email, firstName, lastName, phone string) string {
	// create a user with role = SITE_USER
	u, saveErr := db.User.Create().
		SetEmail(email).
		SetPwd(uid.ULID()).
		SetFirstName(firstName).
		SetLastName(lastName).
		SetPhone(phone).
		SetRole(enum.RoleSiteUser).
		SetStatus(enum.AccountStatusPending).
		Save(ctx)
	if saveErr != nil {
		log.Warn("failed to create user for sales rep %s", email)
		return ""
	}

	uID := u.ID
	// link user to partner
	saveErr = db.PartnerContact.Create().
		SetType(enum.PartnerContactCustom).
		SetRole(enum.PartnerContactRoleSalesRep).
		SetPartnerID(partnerID).
		SetUserID(uID).
		Exec(ctx)
	if saveErr != nil {
		log.Warn("failed to create partner contact relation for sales rep %s", email)
	}

	return uID
}

func QuickAddSalesRepTX(ctx context.Context, tx *ent.Tx, partnerID, email, firstName, lastName, phone string) string {
	// create a user with role = SITE_USER
	u, saveErr := tx.User.Create().
		SetEmail(email).
		SetPwd(uid.ULID()).
		SetFirstName(firstName).
		SetLastName(lastName).
		SetPhone(phone).
		SetRole(enum.RoleSiteUser).
		SetStatus(enum.AccountStatusPending).
		Save(ctx)
	if saveErr != nil {
		log.Warn("failed to create user for sales rep %s", email)
		return ""
	}

	uID := u.ID
	// link user to partner
	saveErr = tx.PartnerContact.Create().
		SetType(enum.PartnerContactCustom).
		SetRole(enum.PartnerContactRoleSalesRep).
		SetPartnerID(partnerID).
		SetUserID(uID).
		Exec(ctx)
	if saveErr != nil {
		log.Warn("failed to create partner contact relation for sales rep %s", email)
	}

	return uID
}
