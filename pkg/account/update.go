/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  21/04/22, 4:44 PM
 */

package account

import (
	"context"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/validate"
	"strings"
)

type UpdateUserInput struct {
	ID        string             `json:"id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Phone     *string            `json:"phone"`
	Role      enum.Role          `json:"role"`
	Status    enum.AccountStatus `json:"status"`
	Note      *string            `json:"note"`
}

func UpdateUser(ctx context.Context, inp *UpdateUserInput, myRole enum.Role) error {
	// validate input
	if err := validate.Struct(inp); err != nil {
		return err
	}

	user, err := ent.GetClient().User.Get(ctx, inp.ID)
	if err != nil || user == nil {
		return msg.AsError(msg.FailedToFind, "user")
	}

	// validate, only Admin can update another Admin user's
	if !checkAdminSavingAdmin(myRole, inp.Role) {
		return msg.AsError(msg.NotAuthorized)
	}

	qry := user.Update().
		SetFirstName(strings.TrimSpace(inp.FirstName)).
		SetLastName(strings.TrimSpace(inp.LastName)).
		SetRole(inp.Role).
		SetStatus(inp.Status)

	if inp.Note == nil {
		qry.ClearNote()
	} else {
		qry.SetNillableNote(inp.Note)
	}

	if inp.Phone != nil {
		qry.SetPhone(strings.TrimSpace(*inp.Phone))
	}

	// update user
	_, err = qry.Save(ctx)
	return err
}
