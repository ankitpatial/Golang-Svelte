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
	"roofix/mailer"
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
	queueUrl, err = queue.GetUrl(config.MailQueueName())
	if err != nil {
		log.Error(err)
		return "", err
	}

	fmt.Println("queue: " + queueUrl)
	fmt.Printf("message count: %d\n", len(evt.Records))

	var wg sync.WaitGroup
	for _, e := range evt.Records {
		wg.Add(1)
		go process(ctx, &wg, e)
	}
	wg.Wait()

	m := "Done!"
	fmt.Println(m)
	return m, nil
}

func process(ctx context.Context, wg *sync.WaitGroup, e events.SQSMessage) {
	defer wg.Done()

	// parse message
	var msg mailer.Message
	var err error
	if ctx, msg, err = queue.Unmarshal[mailer.Message](ctx, e.Body); err != nil {
		log.Error(err)
		return
	}

	if err = mailer.Process(ctx, &msg); err != nil {
		log.Error(err)
		return
	} else {
		queue.Remove(ctx, queueUrl, e.ReceiptHandle)
	}

}
