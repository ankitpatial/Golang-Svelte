/* Copyright (c) 2022. Ankit Patial.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 */

package ws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"net/url"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/usersession"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
)

var apiClient *apigatewaymanagementapi.Client

type wsApi struct {
	domain string
}

func (ws wsApi) ResolveEndpoint(region string, _ apigatewaymanagementapi.EndpointResolverOptions) (aws.Endpoint, error) {
	var u url.URL
	u.Host = ws.domain
	u.Scheme = "https"
	ep := aws.Endpoint{
		SigningRegion: region,
		URL:           u.String(),
	}
	return ep, nil
}

func client() *apigatewaymanagementapi.Client {
	if apiClient != nil {
		return apiClient
	}

	cfg := config.AwsConfig()
	api := wsApi{domain: config.WsDomain()}
	apiClient = apigatewaymanagementapi.NewFromConfig(cfg, apigatewaymanagementapi.WithEndpointResolver(api))
	return apiClient
}

func Register(ctx context.Context, connID, token string) {
	if token == "" {
		log.Warn("TOKEN MISSING, ws register %v - token", connID)
		return
	}

	log.Info("ws register %s - token", connID)

	var sessID string
	var err error
	if sessID, err = crypt.Decrypt(token); err != nil {
		log.Error(err)
		return
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// pull user-session
	live, err := db.UserSession.Query().Where(usersession.ID(sessID)).Exist(ctx)
	if err != nil { // unexpected error
		log.Error(err)
		return
	}

	if live {
		if err = db.UserSessionSocket.Create().SetID(connID).SetSessionID(sessID).Exec(ctx); err != nil {
			log.Error(err)
			return
		}
		log.Info("** new socket connID %s", connID)
	}
}

func UnRegister(ctx context.Context, connId string) {
	log.Info("ws unregister %s", connId)
	removeConnection(ctx, connId)
}

func removeConnection(ctx context.Context, connID string) {
	db := ent.GetClient()
	defer db.CloseClient()

	if err := db.UserSessionSocket.DeleteOneID(connID).Exec(ctx); err != nil {
		if !ent.IsNotFound(err) {
			log.Error(err)
		}
	}
}
