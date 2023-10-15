/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package task

import (
	"context"
	"encoding/json"
	"roofix/config"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/pkg/util/app"
	"roofix/pkg/util/log"
	"roofix/pkg/util/queue"
	"roofix/pkg/ws"
)

const (
	NearMapEstimate    Name = "nearmap-estimate"
	EagleViewEstimate  Name = "eagle-view-estimate"
	EagleViewGetReport Name = "eagle-view-getReport"
	Estimate           Name = "estimate"
)

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) Humanize() string {
	switch n {
	case NearMapEstimate, EagleViewEstimate:
		return "Estimate Request"
	case EagleViewGetReport:
		return "Get Estimate Report"
	default:
		return ""
	}
}

type Task interface {
	Name() Name
	Process(ctx context.Context, progress chan string, done chan error)
}

type Message struct {
	Name         Name     `json:"name"`
	Data         string   `json:"data"`
	NotifyUsers  []string `json:"users"`
	NotifyRoutes []string `json:"paths"`
}

func Submit(ctx context.Context, t Task) error {
	log.Info("submit task: %s", t.Name())

	// marshal task message
	raw, _ := json.Marshal(t)
	m := Message{
		Name: t.Name(),
		Data: string(raw),
		NotifyUsers: []string{
			account.CtxUserID(ctx),
		},
	}

	// in DEV mode
	if config.IsDevEnv() {
		log.Info("dev mode, process task: %s", t.Name())
		go func() {
			c := app.SetCtx(context.Background(), app.ReadCtx(ctx))
			if err := Process(c, t, &m); err != nil {
				log.PrintError(err)
			}
		}()

		return nil
	}

	// production submit to queue
	q := config.TaskQueueName()
	return queue.Send(ctx, q, m)
}

func Process(ctx context.Context, t Task, m *Message) error {
	if m == nil {
		log.Warn("nil Message was passed on")
		return nil
	}

	log.Info("task %s\n", t.Name())

	p := make(chan string)
	d := make(chan error)

	userSockets, er := ws.GetUserSocket(ctx, m.NotifyUsers)
	if er != nil {
		return er
	}

	hasUserSoc := len(userSockets) > 0

	go t.Process(ctx, p, d)
	for {
		select {
		case pm := <-p:
			log.Info("task %s: %s", t.Name(), pm)
			if hasUserSoc {
				ws.Send(ctx, userSockets, ws.Message{
					Channel: enum.ChannelTask,
					Topic:   enum.TopicProgress,
					Title:   t.Name().Humanize(),
					Message: pm,
					Data: map[string]interface{}{
						"task": t.Name(),
						"done": false,
					},
				})
			}

		case err := <-d:
			msg := "Done!"
			done := true
			if err != nil {
				msg = "Failed: " + err.Error()
				done = false
			}

			if hasUserSoc {
				ws.Send(ctx, userSockets, ws.Message{
					Channel: enum.ChannelTask,
					Topic:   enum.TopicProgress,
					Title:   t.Name().Humanize(),
					Message: msg,
					Data: map[string]interface{}{
						"task": t.Name(),
						"done": done,
					},
				})
			}
			return err
		}
	}
}
