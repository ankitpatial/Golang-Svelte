/* Copyright (c) 2022. Ankit Patial.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 */

package ws

import (
	"context"
	"encoding/json"
	"roofix/ent"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
	"time"
)

type message struct {
	Channel enum.Channel           `json:"channel"`
	Topic   enum.Topic             `json:"topic"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func MsgFromClient(ctx context.Context, connId string, msg string) error {
	var m message
	if err := json.Unmarshal([]byte(msg), &m); err != nil {
		log.Error(err)
		return err
	}

	switch m.Channel {
	case enum.ChannelPing:
		ping(ctx, connId)
	}
	return nil
}

// ping is to set user's ID and Pathname info
func ping(ctx context.Context, connID string) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.UserSessionSocket.UpdateOneID(connID).SetUpdatedAt(time.Now().UTC())
	if _, err := qry.Save(ctx); err != nil {
		if !ent.IsNotFound(err) {
			log.Error(err)
		}
		return
	}
}
