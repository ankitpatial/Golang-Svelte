/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package apiaccess

import (
	"context"
	"roofix/ent"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/validate"
	"time"
)

// Get api access info
// must be used at server level only.
// do not pull this info to client level.
func Get(ctx context.Context, name string) (*ent.ApiAccess, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	api, err := db.ApiAccess.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	if api.Password != "" {
		api.Password, err = crypt.Decrypt(api.Password)
		if err != nil {
			return nil, err
		}
	}

	if api.Key != "" {
		api.Key, err = crypt.Decrypt(api.Key)
		if err != nil {
			return nil, err
		}
	}

	if api.Secret != "" {
		api.Secret, err = crypt.Decrypt(api.Secret)
		if err != nil {
			return nil, err
		}
	}

	return api, nil
}

// GetKey api access info
// must be used at server level only.
// do not pull this info to client level.
func GetKey(ctx context.Context, name string) (*ent.ApiAccess, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	api, err := db.ApiAccess.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	if api.Key != "" {
		api.Key, err = crypt.Decrypt(api.Key)
		if err != nil {
			return nil, err
		}
	}

	return api, nil
}

func Save(ctx context.Context, input *Input) (bool, error) {
	if input.ID == "" {
		_, err := create(ctx, input)
		return err != nil, err
	}

	return update(ctx, input)
}

func UpdateApiAccessKey(ctx context.Context, id, key string) error {
	v, err := crypt.Encrypt(key)
	if err != nil {
		return err
	}

	db := ent.GetClient()
	defer db.CloseClient()
	return db.ApiAccess.UpdateOneID(id).SetKey(v).Exec(ctx)
}

func UpdateApiAccessSecret(ctx context.Context, id, secret string) error {
	v, err := crypt.Encrypt(secret)
	if err != nil {
		return err
	}

	db := ent.GetClient()
	defer db.CloseClient()
	return db.ApiAccess.UpdateOneID(id).SetSecret(v).Exec(ctx)
}

func UpdateToken(ctx context.Context, apiAccessID, token, refreshToken string, lifeInSec time.Duration) error {
	db := ent.GetClient()
	defer db.CloseClient()

	return db.ApiAccess.UpdateOneID(apiAccessID).
		SetAccessToken(token).
		SetRefreshToken(refreshToken).
		SetExpiresAt(time.Now().Add(lifeInSec)).
		Exec(ctx)
}

func create(ctx context.Context, input *Input) (string, error) {
	if err := validate.Struct(input); err != nil {
		return "", err
	}

	db := ent.GetClient()
	defer db.CloseClient()

	pwd := input.Password
	key := input.Key
	secret := input.Secret

	var err error
	if pwd != "" {
		pwd, err = crypt.Encrypt(pwd)
		if err != nil {
			return "", err
		}
	}

	if key != "" {
		key, err = crypt.Encrypt(key)
		if err != nil {
			return "", err
		}
	}

	if secret != "" {
		secret, err = crypt.Encrypt(secret)
		if err != nil {
			return "", err
		}
	}

	qry := db.ApiAccess.Create().
		SetID(input.Name).
		SetURL(input.URL).
		SetUsername(input.Username).
		SetPassword(pwd).
		SetKey(key).
		SetSecret(secret)

	o, err := qry.Save(ctx)
	if err != nil {
		return "", err
	}

	return o.ID, nil
}

func update(ctx context.Context, input *Input) (bool, error) {

	return true, nil
}
