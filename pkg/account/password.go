package account

import (
	"context"
	"errors"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/user"
	"roofix/mailer"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/pkg/util/validate"
	"roofix/template"
	"time"
)

const PathAccSetPwd = "/account/set-password?token="

// SetMyPassword for logged-in user's account
func SetMyPassword(ctx context.Context, ctxUserID, oldPwd, newPwd, confirmPwd string) error {
	if ctxUserID == "" || oldPwd == "" || newPwd == "" || confirmPwd == "" {
		return msg.AsError(msg.OneOfParamMissing)
	}

	if newPwd != confirmPwd {
		return errors.New("confirm password does not match")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// pull user
	u, err := db.User.Get(ctx, ctxUserID)
	if err != nil && ent.IsNotFound(err) {
		return msg.AsError(msg.FailedToFind, "user")
	} else if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}

	// check old pwd
	if !crypt.CompareHash(u.Pwd, oldPwd) {
		return errors.New("old password is not correct")
	}

	// set new password
	_, err = u.Update().SetPwd(crypt.Hash(newPwd)).Save(ctx)

	// TODO: trigger notification
	return err
}

// SetUserPassword is supposed to be called by ADMIN to set none admin user password
// and will not be used to change self or other ADMIN user password
func SetUserPassword(ctx context.Context, ctxUserID, userID, newPwd, confirmPwd string) error {
	if ctxUserID == "" || userID == "" || newPwd == "" || confirmPwd == "" {
		return msg.AsError(msg.OneOfParamMissing)
	}

	if newPwd != confirmPwd {
		return errors.New("confirm password does not match")
	}

	if ctxUserID == userID { // session user must not use it set its own password
		return errors.New("not be used to set self password")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// pull user
	u, err := db.User.Get(ctx, userID)
	if err != nil && ent.IsNotFound(err) {
		return msg.AsError(msg.FailedToFind, "user")
	} else if err != nil {
		log.Error(err)
		return msg.AsError(msg.ServerError)
	}

	if u.Role == enum.RoleAdmin { // ADMIN account password can't be changed in this func
		return errors.New("to set admin account password please try forgot password functionality")
	}

	// set new password
	_, err = u.Update().SetPwd(crypt.Hash(newPwd)).Save(ctx)

	// TODO: trigger notification
	return err
}

// ChangePasswordRequest will send reset password email with a link to set new password
func ChangePasswordRequest(ctx context.Context, email string) error {
	if email == "" {
		return msg.AsError(msg.NotFound, "email")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	e := str.TrimAndToLower(email)
	u, err := db.User.Query().
		Where(user.Email(e)).
		Select(user.FieldEmail, user.FieldStatus).
		First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return msg.AsError(msg.NotFound, "email")
	} else if err != nil {
		return err
	}

	if u.Status != enum.AccountStatusActive {
		return msg.AsError(msg.AccountNotActive)
	}

	token, err := NewToken(ctx, u.ID, TokenChangePassword, time.Hour*6)
	if err != nil {
		return err
	}

	// send reset password email
	mailer.Send(ctx, &mailer.Message{
		To: []string{
			u.Email,
		},
		Subject: "Your Roofix account password",
		Tpl:     template.EmailSetPassword,
		Modal: map[string]interface{}{
			"name":     u.Name(),
			"pathname": PathAccSetPwd + token,
		},
	})

	return nil
}

// SetPassword to new account or old account
func SetPassword(ctx context.Context, inp *SetPasswordInput) (*ent.User, error) {
	if err := validate.Struct(inp); err != nil {
		return nil, err
	}

	if inp.Password != inp.ConfirmPassword {
		return nil, msg.AsError(msg.ConfirmPwdMismatch)
	}

	// parse token
	c, err := parseToken(inp.Token)
	if err != nil {
		log.Error(err)
		return nil, errors.New("invalid token")
	}

	// pull from db
	t, err := getToken(ctx, c["id"].(string))
	if err != nil {
		return nil, err
	}

	if t.ConfirmedAt != nil {
		return nil, errors.New("token is already used")
	}

	userID := t.Data["userID"].(string)

	db := ent.GetClient()
	defer db.CloseClient()

	u, err := db.User.Query().Where(user.ID(userID)).First(ctx)
	if err != nil && ent.IsNotFound(err) {
		return nil, msg.AsError(msg.NotFound, "email")
	} else if err != nil {
		return nil, err
	}

	switch t.Action {
	case "SetPassword": // new account set password
		_, err = u.Update().
			SetPwd(crypt.Hash(inp.Password)).
			SetEmailVerified(true).
			SetStatus(enum.AccountStatusActive).
			Save(ctx)
	case "ChangePassword":
		_, err = u.Update().
			SetPwd(crypt.Hash(inp.Password)).
			ClearWrongAttempts().
			ClearWrongAttemptAt().
			ClearLockedUntil().
			Save(ctx)
	}

	if err != nil {
		return nil, err
	}

	setTokenUsed(ctx, t.ID)

	// sign in
	return u, nil
}

// SetPasswordLink is url to where user can set a new account password
func SetPasswordLink(ctx context.Context, uID string) string {
	if token, err := NewToken(ctx, uID, TokenSetPassword, time.Hour*48); err != nil {
		log.Error(err)
		return ""
	} else {
		return config.Read().Website.Url + PathAccSetPwd + token
	}
}
