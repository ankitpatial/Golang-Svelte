package partner

import (
	"context"
	"errors"
	"fmt"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/notification"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/pkg/util/uid"
	"roofix/pkg/util/validate"
	"roofix/template"
	"strings"
	"time"
)

// SaveContacts is step 2 for partner on-boarding
func SaveContacts(ctx context.Context, partnerID string, contacts []*ContactUserInput) ([]*ent.PartnerContact, error) {
	if partnerID == "" {
		return nil, msg.AsError(msg.ParamMissing, "partnerID")
	}

	if len(contacts) == 0 {
		log.Warn("SaveContacts was called with empty data")
		return nil, nil
	}

	db := ent.GetClient()
	defer db.CloseClient()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// var userIDs []string
	for _, inp := range contacts {
		// validate struct
		if e := validate.Struct(inp); e != nil {
			return nil, e
		}

		// save user
		var userID string
		if userID, err = saveUser(ctx, tx, partnerID, inp); err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		// save partner contact
		if err = saveContact(ctx, tx, partnerID, userID, inp); err != nil {
			_ = tx.Rollback()
			return nil, err
		}

		// add userID to list
		//userIDs = append(userIDs, userID)
	}

	// update partner steps completed
	err = tx.Partner.
		Update().
		Where(partner.ID(partnerID), partner.SetupStepsCompletedLT(2)).
		SetSetupStepsCompleted(2).
		Exec(ctx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	audit.OpSuccess(ctx, audit.PartnerSaveContacts, fmt.Sprintf("save partnerID: %s", partnerID))
	return db.PartnerContact.Query().
		WithUser().
		Where(partnercontact.PartnerID(partnerID)).
		Order(ent.Asc(partnercontact.FieldCreatedAt)).
		All(ctx)
}

func SaveContact(ctx context.Context, partnerID string, contact *ContactUserInput) error {
	if partnerID == "" {
		return msg.AsError(msg.ParamMissing, "partnerID")
	}

	// validate struct
	if err := validate.Struct(contact); err != nil {
		return err
	}

	db := ent.GetClient()
	defer db.CloseClient()
	// begin a TX
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return err
	}

	// save user
	userID, err := saveUser(ctx, tx, partnerID, contact)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// save partner contact
	if err = saveContact(ctx, tx, partnerID, userID, contact); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		return err
	}

	// new user account created
	// send notification
	if contact.UserID == "" {
		// email ==>
		var pName, token string
		if p, _ := db.Partner.Get(ctx, partnerID); p != nil {
			pName = p.Name
		}
		token, _ = account.NewToken(ctx, userID, account.TokenSetPassword, time.Hour*6)
		if pName != "" && token != "" {
			mailer.Send(ctx, &mailer.Message{
				To: []string{
					fmt.Sprintf("%s %s<%s>", contact.FirstName, contact.LastName, contact.Email),
				},
				Subject: "RFX Account Created",
				Tpl:     template.EmailPartnerNewUser,
				Modal: map[string]interface{}{
					"partner":  pName,
					"role":     contact.Role.Humanize(),
					"username": contact.Email,
					"pathname": account.PathAccSetPwd + token,
				},
			})
		}

		// notification ==>
		notification.SendAndSave(
			ctx,
			&notification.Message{
				Channel: enum.ChannelPartnerUser,
				Topic:   enum.TopicCreated,
				Title:   "Added New User",
				RefID:   partnerID,
				Message: fmt.Sprintf("A new user %s %s is added as %s", contact.FirstName, contact.LastName, contact.Role.Humanize()),
				Data: map[string]interface{}{
					partnerID: partnerID,
					userID:    userID,
				},
				Audience: notification.Audience{},
			},
			notification.ToUser(account.CtxUserID(ctx)),
			notification.ToCompanyAdmins(partnerID),
			notification.ToAdmins(),
		)
	}

	return nil
}

// saveUser info fpr partner contact
func saveUser(
	ctx context.Context, tx *ent.Tx, partnerID string, inp *ContactUserInput,
) (string, error) {
	if tx == nil || inp == nil {
		return "", errors.New("missing required fields")
	}

	userID := inp.UserID
	fName := strings.TrimSpace(inp.FirstName)
	lName := strings.TrimSpace(inp.LastName)
	email := strings.ToLower(strings.TrimSpace(inp.Email))
	ph := strings.TrimSpace(inp.Phone)

	// check user email available
	uEmailQry := tx.User.Query().Where(user.EmailEqualFold(email))
	if userID != "" {
		uEmailQry.Where(user.IDNEQ(userID))
	}
	if exist, _ := uEmailQry.Exist(ctx); exist {
		return "", errors.New(fmt.Sprintf("%s email address is not avialbale", inp.Type))
	}

	if userID != "" { // update USER
		// user must be linked to given partnerID
		c, err := tx.PartnerContact.
			Query().
			Where(partnercontact.UserID(userID), partnercontact.PartnerID(partnerID)).
			Count(ctx)
		if err != nil {
			return "", err
		} else if c == 0 {
			return "", errors.New(fmt.Sprintf("email '%s' is associated to another account", inp.Email))
		}

		qry := tx.User.
			UpdateOneID(userID).
			SetEmail(email).
			SetFirstName(fName).
			SetLastName(lName).
			SetPhone(ph)
		// change account status
		if inp.AccountStatus != nil {
			qry.SetStatus(*inp.AccountStatus)
		}
		// save
		if err = qry.Exec(ctx); err != nil {
			return "", err
		}
	} else { // create USER
		randomPwd := crypt.Hash(uid.ULID())
		qry := tx.User.
			Create().
			SetEmail(email).
			SetPwd(randomPwd).
			SetFirstName(fName).
			SetLastName(lName).
			SetPhone(ph).
			SetRole(enum.RoleSiteUser).
			SetStatus(enum.AccountStatusPending)

		if u, err := qry.Save(ctx); err != nil {
			return "", err
		} else {
			userID = u.ID
		}
	}

	return userID, nil
}

// saveContact detail of partner
func saveContact(ctx context.Context, tx *ent.Tx, partnerID, userID string, inp *ContactUserInput) error {
	conType := inp.Type
	id := str.Val(inp.ID)

	// create
	if id == "" {
		qry := tx.PartnerContact.
			Create().
			SetType(conType).
			SetPartnerID(partnerID).
			SetUserID(userID).
			SetNillableTitle(inp.Title).
			SetNillableDescription(inp.Description)

		if conType == enum.PartnerContactAccounting && inp.OtherEmail != "" {
			qry.SetInvoicingEmail(inp.OtherEmail)
		}

		if inp.Role != nil {
			qry.SetRole(*inp.Role)
		} else {
			switch inp.Type {
			case enum.PartnerContactPrimary,
				enum.PartnerContactOperations,
				enum.PartnerContactAccounting,
				enum.PartnerContactInvoicing,
				enum.PartnerContactCustomerService:
				qry.SetRole(enum.PartnerContactRoleAdmin)
			default:
				qry.SetRole(enum.PartnerContactRoleSalesRep)
			}
		}

		return qry.Exec(ctx)
	}

	// update
	qry := tx.PartnerContact.
		UpdateOneID(id).
		SetType(conType).
		SetNillableTitle(inp.Title).
		SetNillableDescription(inp.Description)

	if conType == enum.PartnerContactAccounting && inp.OtherEmail != "" {
		qry.SetInvoicingEmail(inp.OtherEmail)
	}

	return qry.Exec(ctx)
}
