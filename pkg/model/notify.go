package model

import (
	"entgo.io/contrib/entgql"
	"roofix/pkg/enum"
	"time"
)

type NotifyMessageConnection struct {
	TotalCount int                      `json:"totalCount"`
	PageInfo   *entgql.PageInfo[string] `json:"pageInfo"`
	Edges      []*NotifyMessageEdge     `json:"edges,omitempty"`
}

type NotifyMessageEdge struct {
	Cursor entgql.Cursor[string] `json:"cursor"`
	Node   *NotifyMessage        `json:"node"`
}

type Notify struct {
	ID           string `json:"id"`
	Topic        string `json:"topic"`
	ReceiveEmail bool   `json:"receiveEmail"`
	ReceiveSms   bool   `json:"receiveSMS"`
}

type NotifyMessage struct {
	ID        string       `json:"id"`
	Channel   enum.Channel `json:"channel"`
	Topic     enum.Topic   `json:"topic"`
	RefID     *string      `json:"refID,omitempty"`
	Title     string       `json:"title"`
	Message   string       `json:"message"`
	From      string       `json:"from"`
	Unread    bool         `json:"unread"`
	CreatedAt time.Time    `json:"createdAt"`
}
