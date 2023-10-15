/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/pkg/eagleview"
	"roofix/pkg/estimate"
	"roofix/pkg/task"
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
	queueUrl, _ = queue.GetUrl(config.TaskQueueName())

	log.Info("queue: " + queueUrl)
	log.Info("message count: %d", len(evt.Records))

	var wg sync.WaitGroup
	for _, rec := range evt.Records {
		wg.Add(1)
		go process(ctx, rec, &wg)
	}
	wg.Wait()

	m := "Done!"
	log.Info(m)
	return m, nil
}

func process(ctx context.Context, sqsMsg events.SQSMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	var m task.Message
	var err error
	if ctx, m, err = queue.Unmarshal[task.Message](ctx, sqsMsg.Body); err != nil {
		log.Error(err)
		return
	}

	log.Info("=== task %s ===\n", m.Name)
	t := taskByName(m)
	if fmt.Sprintf("%v", t) == "<nil>" {
		log.Warn("WARN worker for task '%s' not found\n", m.Name)
		queue.Remove(ctx, queueUrl, sqsMsg.ReceiptHandle)
		return
	}

	if err = task.Process(ctx, t, &m); err != nil {
		log.Error(err)
		return
	} else {
		queue.Remove(ctx, queueUrl, sqsMsg.ReceiptHandle)
	}
}

func taskByName(m task.Message) task.Task {
	switch m.Name {
	case task.Estimate:
		var t estimate.BGTask
		_ = json.Unmarshal([]byte(m.Data), &t)
		return t
	//case task.RfxEstimate:
	//	var t rfx.TaskRfxEstimate
	//	_ = json.Unmarshal([]byte(m.Data), &t)
	//	return t
	case task.EagleViewEstimate:
		var t eagleview.TaskEstimate
		_ = json.Unmarshal([]byte(m.Data), &t)
		return t
	case task.EagleViewGetReport:
		var t eagleview.TaskGetReport
		_ = json.Unmarshal([]byte(m.Data), &t)
		return t

	default:
		return nil
	}
}
