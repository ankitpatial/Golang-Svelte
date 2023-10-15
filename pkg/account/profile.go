package account

import (
	"context"
	"errors"
	"roofix/ent"
	"roofix/ent/user"
	"roofix/pkg/model"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"strings"
)

func UpdateProfile(ctx context.Context, u *User, inp *model.InputUserProfile) error {
	if u == nil {
		return msg.AsError(msg.ParamMissing, "'user'")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.User.UpdateOneID(u.ID)
	// First Name
	if inp.FirstName != nil {
		qry.SetFirstName(*inp.FirstName)
	}

	// Last Name
	if inp.LastName != nil {
		qry.SetLastName(*inp.LastName)
	}

	// Phone
	if inp.Phone != nil {
		qry.SetPhone(*inp.Phone)
	}

	// Profile Picture
	if inp.Picture != nil {
		qry.SetPicture(*inp.Picture)
	}

	// reset password
	oldPwd := strings.TrimSpace(str.Val(inp.OldPwd))
	if oldPwd != "" {
		newPwd := strings.TrimSpace(str.Val(inp.NewPwd))
		confirmPwd := strings.TrimSpace(str.Val(inp.ConfirmNewPwd))
		// validate
		if newPwd == "" || confirmPwd == "" {
			return errors.New("to change password:  Old, New & Confirm passwords are required")
		}

		// compare
		if newPwd != confirmPwd {
			return errors.New("confirm password does not match")
		}

		// pull use with old password
		u, err := db.User.Query().Where(user.ID(u.ID)).Select(user.FieldPwd).Only(ctx)
		if err != nil {
			log.Error(err)
			return msg.AsError(msg.ServerError)
		}

		// compare old pwd
		if !crypt.CompareHash(u.Pwd, oldPwd) {
			return errors.New("old password is not correct")
		}

		// set new pwd
		qry.SetPwd(crypt.Hash(newPwd))
	}

	// ** save user data **
	if err := qry.Exec(ctx); err != nil {
		log.Error(err)
		return errors.New("failed to save profile info")
	}

	return nil
}
