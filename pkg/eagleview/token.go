/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package eagleview

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"roofix/pkg/apiaccess"
	"roofix/pkg/util/http"
	"roofix/pkg/util/log"
	"time"
)

type TokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint   `json:"expires_in"`
}

type ErrResp struct {
	Message     string `json:"Message"`
	Err         string `json:"error"`
	Description string `json:"error_description"`
}

func (er ErrResp) Error() string {
	if er.Description == "" {
		return er.Err
	}

	return fmt.Sprintf("%s, %s", er.Err, er.Description)
}

func apiToken(ctx context.Context) (host, token string, err error) {
	api, err := apiaccess.Get(ctx, ProviderName)
	if err != nil {
		return host, token, err
	}

	host = api.URL

	// has token and it's still valid
	if api.AccessToken != "" && time.Now().Before(api.ExpiresAt) {
		log.Info("use token from DB")
		token = api.AccessToken
		return host, token, nil
	}

	log.Info("fetching token...")
	ep := fmt.Sprintf("%s/Token", api.URL)
	vals := url.Values{
		"grant_type": {"password"},
		"username":   {api.Username},
		"password":   {api.Password},
	}
	d, err := http.Post(ep, vals, http.WithAuthBasic(api.Key, api.Secret), http.WithContentTypeForm())
	if err != nil {
		return host, token, err
	}

	// parse response
	var res TokenResp
	if err := json.Unmarshal(d, &res); err != nil {
		return host, token, err
	}

	token = res.AccessToken

	log.Info("got token, going to save it in DB")
	life := time.Second * time.Duration(res.ExpiresIn)
	if err := apiaccess.UpdateToken(ctx, api.ID, token, res.RefreshToken, life); err != nil {
		log.Error(err)
	}

	return host, token, nil
}
