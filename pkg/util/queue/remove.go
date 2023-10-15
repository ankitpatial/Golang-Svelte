/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package queue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"roofix/pkg/util/log"
)

func Remove(ctx context.Context, queueUrl, receiptHandle string) {
	if queueUrl == "" {
		log.Warn("abort, Queue URL is empty")
		return
	}

	_, err := client().DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &queueUrl,
		ReceiptHandle: &receiptHandle,
	})

	if err != nil {
		log.Error(err)
	} else {
		log.Info("queue Remove: %s", queueUrl)
	}
}
