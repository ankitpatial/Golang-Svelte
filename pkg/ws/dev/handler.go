/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package dev

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"roofix/pkg/util/uid"
)

type WsData struct {
	ConnID string                 `json:"connectionId"`
	Data   map[string]interface{} `json:"data"`
}

func (d WsData) Bind(r *http.Request) error {
	return nil
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		q := r.URL.Query()
		client := &Client{
			id:    uid.ULID(),
			hub:   hub,
			conn:  conn,
			send:  make(chan []byte, 256),
			token: q.Get("token"),
		}

		client.hub.Register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		go client.readPump()

		render.Status(r, http.StatusOK)
	}
}

func HandleWsMsg(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body WsData
		if err := render.Bind(r, &body); err != nil {
			render.JSON(w, r, map[string]interface{}{
				"error": err,
			})
		}

		for cl := range hub.Clients {
			if cl.id == body.ConnID {
				d, _ := json.Marshal(body.Data)
				cl.send <- d
			}
		}

		render.Status(r, http.StatusOK)
	}
}
