/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  05/04/22, 5:08 PM
 */

package mailer

import (
	"context"
	_ "embed"
	"fmt"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/pkg/util/queue"
	"roofix/pkg/util/storage"
	"roofix/pkg/util/str"
)

var mailer smtp

func init() {
	if config.Read().IsDevEnv() {
		mailer = fileDump{}
		return
	}

	mailer = awsSES{}
}

// Send mail message to queue for processing
func Send(ctx context.Context, msg *Message) {
	if msg == nil {
		log.Warn("mailer.Send(msg) was called with nil data")
		return
	}

	if len(msg.To) == 0 {
		log.Warn("WARN mailer.Send(msg) was called with no recipient")
		return
	}

	if config.IsDevEnv() {
		if body, err := msg.bodyHTML(ctx); err != nil {
			log.Error(err)
		} else if err = mailer.Send(ctx, msg.To, msg.Subject, body); err != nil {
			log.Error(err)
		}
		return
	}

	// submit to queue
	if err := queue.Send(ctx, config.MailQueueName(), msg); err != nil {
		log.Error(err)
	}

	log.Info("email is in queue: to: %v, subject: %s\n", msg.To, msg.Subject)
}

// Process mail message and deliver
func Process(ctx context.Context, msg *Message) error {
	var to []string
	if config.IsDevEnv() {
		to = msg.To
	} else { // check bounce history
		to = sanitize(ctx, msg.To)
		if to == nil || len(to) == 0 {
			log.Warn("no recipient found")
			return nil
		}
	}

	// render email template
	htmlBody, err := msg.bodyHTML(ctx)
	if err != nil {
		return err
	}

	// send mail
	return mailer.Send(ctx, to, msg.Subject, htmlBody)
}

// sanitize will check the mail address in bounce history
func sanitize(ctx context.Context, list []string) []string {
	var res []string
	bucket := config.LogBucket()

	for _, email := range list {
		key := fmt.Sprintf("mail/bounce/%s", str.TrimAndToLower(email))
		if inBounce := storage.Exists(ctx, bucket, key); inBounce {
			log.Info("email %s is in bounce list", email)
		} else {
			res = append(res, email)
		}
	}

	return res
}
