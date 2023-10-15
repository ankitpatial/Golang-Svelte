package account

import (
	"context"
	_ "embed"
	"roofix/ent"
	"roofix/mailer"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/parse"
	"roofix/pkg/util/str"
	"roofix/pkg/util/uid"
	"roofix/pkg/util/validate"
	"roofix/template"
	"strings"
)

func Create(ctx context.Context, inp *CreateUserInput) (string, error) {
	// validate input
	if err := validate.Struct(inp); err != nil {
		return "", err
	}

	e := str.TrimAndToLower(inp.Email)
	if checkEmailInUse(ctx, e, "") {
		return "", msg.AsError(msg.AlreadyExists, "user with same email")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	//  create user as pending with random password
	randomPwd := crypt.Hash(uid.ULID())
	u, err := db.User.Create().
		SetEmail(e).
		SetPwd(randomPwd).
		SetFirstName(strings.TrimSpace(inp.FirstName)).
		SetLastName(strings.TrimSpace(inp.LastName)).
		SetNillablePhone(inp.Phone).
		SetRole(inp.Role).
		SetNillableNote(inp.Note).
		Save(ctx)

	if err != nil {
		return "", err
	}

	link := SetPasswordLink(ctx, u.ID)
	if link == "" {
		return u.ID, err
	}

	// send email to confirm and set new account password
	data := parse.StructToMap(inp)
	data["link"] = link
	mailer.Send(ctx, &mailer.Message{
		To:      []string{inp.Email},
		Subject: "Roofix New User Account",
		Tpl:     template.EmailNewUserAccount,
		Modal:   data,
	})

	return u.ID, nil
}
