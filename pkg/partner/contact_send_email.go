package partner

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/util/log"
	"roofix/template"
	"time"
)

func ContactResetPasswordEmail(ctx context.Context, userID string) error {
	db := ent.GetClient()
	defer db.CloseClient()

	u, err := db.User.Get(ctx, userID)
	if err != nil {
		log.Error(err)
		return err
	}

	var token string
	token, err = account.NewToken(ctx, userID, account.TokenSetPassword, time.Hour*6)
	if err != nil {
		log.Error(err)
		return err
	}

	// send email
	mailer.Send(ctx, &mailer.Message{
		To: []string{
			fmt.Sprintf("%s %s<%s>", u.FirstName, u.LastName, u.Email),
		},
		Subject: "RFX Reset Password",
		Tpl:     template.EmailSetPassword,
		Modal: map[string]interface{}{
			"pathname": account.PathAccSetPwd + token,
		},
	})

	return nil
}
