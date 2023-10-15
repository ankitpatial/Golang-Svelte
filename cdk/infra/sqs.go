package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewQueue(stack constructs.Construct, name string, retentionDays uint8, visibilitySec float64) awssqs.IQueue {
	deadQName := jsii.String(name + "_DEAD")

	// DEAD queue
	deadQ := awssqs.NewQueue(
		stack,
		deadQName,
		&awssqs.QueueProps{
			QueueName:         deadQName,
			VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(visibilitySec + 5)),
			RetentionPeriod:   awscdk.Duration_Days(jsii.Number(float64(retentionDays))),
			Encryption:        awssqs.QueueEncryption_KMS_MANAGED,
		})

	// queue
	return awssqs.NewQueue(stack, &name, &awssqs.QueueProps{
		QueueName:         &name,
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(visibilitySec)),
		RetentionPeriod:   awscdk.Duration_Days(jsii.Number(float64(retentionDays))),
		Encryption:        awssqs.QueueEncryption_KMS_MANAGED,
		DeadLetterQueue: &awssqs.DeadLetterQueue{
			MaxReceiveCount: jsii.Number(2),
			Queue:           deadQ,
		},
	})
}

func NewFIFOQueue(stack constructs.Construct, name string, retentionDays uint8, visibilitySec float64) awssqs.IQueue {
	deadQName := jsii.String(name + "_DEAD.fifo")
	// DEAD Queue
	deadQ := awssqs.NewQueue(
		stack,
		deadQName,
		&awssqs.QueueProps{
			Encryption:        awssqs.QueueEncryption_KMS_MANAGED,
			Fifo:              jsii.Bool(true),
			QueueName:         deadQName,
			RetentionPeriod:   awscdk.Duration_Days(jsii.Number(float64(retentionDays))),
			VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(visibilitySec + 5)),
		})

	// job queue
	name += ".fifo"
	return awssqs.NewQueue(stack, &name, &awssqs.QueueProps{
		DeadLetterQueue: &awssqs.DeadLetterQueue{
			MaxReceiveCount: jsii.Number(2),
			Queue:           deadQ,
		},
		Encryption:        awssqs.QueueEncryption_KMS_MANAGED,
		Fifo:              jsii.Bool(true),
		QueueName:         &name,
		RetentionPeriod:   awscdk.Duration_Days(jsii.Number(float64(retentionDays))),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(visibilitySec)),
	})
}
