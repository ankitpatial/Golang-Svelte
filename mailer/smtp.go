/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package mailer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"os"
	"path/filepath"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/pkg/util/open"
	"roofix/pkg/util/uid"
)

// smtp interface to send email
type smtp interface {
	Send(ctx context.Context, to []string, subject string, htmlBody string) error
}

// DevMail is to simulate send email at development level
type fileDump struct {
}

func (fileDump) Send(_ context.Context, _ []string, _ string, htmlBody string) error {
	// write tpl to tmp folder
	dir := os.TempDir()
	id := uid.ULID()
	file := filepath.Join(dir, id+".html")
	if err := os.WriteFile(file, []byte(htmlBody), 0440); err != nil {
		return err
	}

	return open.WithDefaultApp(file)
}

// SesMail will AWS:SES as email smtp
type awsSES struct {
}

func (awsSES) Send(ctx context.Context, to []string, subject string, htmlBody string) error {
	sender := config.Read().Mail.Sender
	log.Info("sending email to: %s, subject: %s, sender: %s", to, subject, sender)

	cl := ses.NewFromConfig(config.AwsConfig())
	charSet := "UTF-8"

	_, err := cl.SendEmail(ctx, &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: to,
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data:    aws.String(subject),
				Charset: aws.String(charSet),
			},
			Body: &types.Body{
				Html: &types.Content{
					Data:    &htmlBody,
					Charset: aws.String(charSet),
				},
			},
		},
		Source: aws.String(sender),
	})

	if err != nil {
		log.Error(err)
	}

	return err
}
