/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package dev

import (
	"context"
	service "roofix/pkg/ws"
)

type Msg struct {
	ConnId string
	Data   []byte
}

// Hub maintains the set of active Clients and broadcasts messages to the
// Clients.
type Hub struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan []byte

	// Direct message to client
	Message chan *Msg

	// Register requests from the Clients.
	Register chan *Client

	// Unregister requests from Clients.
	UnRegister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			service.Register(context.Background(), client.id, client.token)
			h.Clients[client] = true
		case client := <-h.UnRegister:
			service.UnRegister(context.Background(), client.id)
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
