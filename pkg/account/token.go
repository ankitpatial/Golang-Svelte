/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package account

import (
	"context"
	"errors"
	"roofix/ent"
	"roofix/pkg/util/log"
	"roofix/pkg/util/uid"
	"time"
)

type TokenType string

const (
	TokenSetPassword    TokenType = "SetPassword"
	TokenChangePassword TokenType = "ChangePassword"
)

func (t TokenType) String() string {
	return string(t)
}

func NewToken(ctx context.Context, userID string, action TokenType, d time.Duration) (string, error) {
	tokenID := uid.ULID()
	token, err := signToken(map[string]interface{}{"id": tokenID}, d)
	if err != nil {
		return "", err
	}

	// save token to db
	db := ent.GetClient()
	defer db.CloseClient()
	err = db.Token.Create().
		SetID(tokenID).
		SetAction(action.String()).
		SetData(map[string]interface{}{
			"userID": userID,
		}).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return "", errors.New("failed to create token")
	}

	return token, nil
}

func getToken(ctx context.Context, id string) (*ent.Token, error) {
	db := ent.GetClient()
	defer db.CloseClient()
	return db.Token.Get(ctx, id)
}

func setTokenUsed(ctx context.Context, id string) {
	db := ent.GetClient()
	defer db.CloseClient()
	_ = db.Token.UpdateOneID(id).SetConfirmedAt(time.Now()).Exec(ctx)
}
