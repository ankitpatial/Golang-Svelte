/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/pkg/util/queue"
)

var queueUrl string

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, evt events.SQSEvent) (string, error) {
	config.PrintBuildInfo()
	var err error
	queueUrl, err = queue.GetUrl(config.MailQueueName())
	if err != nil {
		log.Error(err)
		return "", err
	}

	fmt.Println("queue: " + queueUrl)
	fmt.Printf("message count: %d\n", len(evt.Records))

	for _, e := range evt.Records {
		process(ctx, e)
	}

	m := "Done!"
	fmt.Println(m)
	return m, nil
}

func process(ctx context.Context, e events.SQSMessage) {
	// parse message
	//var msg mailer.Message
	//var err error
	//if ctx, msg, err = queue.Unmarshal[mailer.Message](ctx, e.Body); err != nil {
	//	log.Error(err)
	//	return
	//}

	log.Info("pending to implement")

	queue.Remove(ctx, queueUrl, e.ReceiptHandle)
}
