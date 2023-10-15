/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/pkg/notification"
	"roofix/pkg/util/log"
	"roofix/pkg/util/queue"
	"sync"
)

var queueUrl string

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, evt events.SQSEvent) (string, error) {
	config.PrintBuildInfo()
	var err error
	queueUrl, err = queue.GetUrl(config.NotificationQueueName())
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.InfoBullet("queue: " + queueUrl)
	log.Info("message count: %d\n", len(evt.Records))

	var wg sync.WaitGroup
	wg.Add(len(evt.Records))
	for _, e := range evt.Records {
		go process(ctx, &wg, e)
	}
	wg.Wait()

	m := "Done!"
	log.Info(m)
	return m, nil
}

func process(ctx context.Context, wg *sync.WaitGroup, e events.SQSMessage) {
	defer wg.Done()

	var msg notification.Message
	var err error
	log.Info("%s", e.Body)
	if ctx, msg, err = queue.Unmarshal[notification.Message](ctx, e.Body); err != nil {
		log.Error(err)
		return
	}

	if err = notification.Process(ctx, &msg); err != nil {
		log.Error(err)
		return
	}

	queue.Remove(ctx, queueUrl, e.ReceiptHandle)
}
