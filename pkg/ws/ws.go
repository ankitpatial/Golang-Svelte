/* Copyright (c) 2022. Ankit Patial.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 */

package ws

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/types"
	"net/http"
	"roofix/config"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
)

type Option int

type Message struct {
	Channel enum.Channel           `json:"channel"`
	Topic   enum.Topic             `json:"topic"`
	Title   string                 `json:"title"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Send will post data to given socketID list
func Send(ctx context.Context, connIDs []string, msg Message) {
	if len(connIDs) == 0 {
		log.Warn("connIDs list is empty")
		return
	}

	isDev := config.IsDevEnv()
	// Send jobs to the job channel
	for _, connID := range connIDs {
		post(ctx, connID, &msg, isDev)
	}
}

func post(ctx context.Context, connID string, msg *Message, isDev bool) {
	log.Info("processing %s by workedID %d", connID)
	// DEV env
	if isDev {
		b, _ := json.Marshal(map[string]interface{}{
			"connectionId": connID,
			"data":         msg,
		})

		ep := config.Read().Website.Url + "/ws-msg"
		if _, err := http.Post(ep, "application/json", bytes.NewBuffer(b)); err != nil {
			log.Error(err)
		}
		log.Info("posted message to connId %s %s", connID, ep)

		return
	}

	// PROD env where AWS WS Api is in place]
	data, _ := json.Marshal(msg)
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connID),
		Data:         data,
	}
	if _, err := client().PostToConnection(ctx, input); err != nil {
		if errors.Is(err, &types.GoneException{}) {
			// delete
			log.Info("socket client is gone, time to remove the connection")
			removeConnection(ctx, connID)
		} else {
			log.Error(err)
		}
	} else {
		log.Info("posted message to connId %s", connID)
	}
}
