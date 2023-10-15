/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package mailer

import (
	"context"
	"roofix/template"
)

type Message struct {
	To      []string               `json:"to,omitempty"`
	Subject string                 `json:"subject,omitempty"`
	Tpl     template.Name          `json:"tpl,omitempty"`
	Modal   map[string]interface{} `json:"modal,omitempty"`
}

func (msg *Message) bodyHTML(ctx context.Context) (string, error) {
	return template.Exec(ctx, template.KindEmail, msg.Tpl, msg.Modal)
}
