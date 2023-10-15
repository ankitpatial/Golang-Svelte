/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/mailer/feedback"
	"roofix/pkg/util/log"
	"sync"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, evt events.SNSEvent) (string, error) {
	config.PrintBuildInfo()
	var wg sync.WaitGroup
	for _, r := range evt.Records {
		go process(ctx, &wg, r)
		wg.Add(1)
	}
	wg.Wait()
	return "", nil
}

func process(ctx context.Context, wg *sync.WaitGroup, r events.SNSEventRecord) {
	defer wg.Done()
	msg := r.SNS.Message
	fmt.Println(msg)

	var evt feedback.Event
	if err := json.Unmarshal([]byte(msg), &evt); err != nil {
		log.Error(err)
		return
	}

	switch evt.Type {
	case "Send":
		feedback.Send(ctx, &evt.Mail)
	case "Delivery":
		feedback.Delivery(ctx, &evt.Mail, evt.Delivery)
	case "Bounce":
		feedback.Bounce(ctx, &evt.Mail, evt.Bounce)
	default:
		fmt.Printf("Event Current: %s is not handled", evt.Type)
	}
}
