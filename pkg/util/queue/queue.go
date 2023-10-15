package queue

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"os"
	"roofix/config"
	"roofix/pkg/util/app"
	"roofix/pkg/util/log"
)

type queueMsg struct {
	CtxData app.CtxData `json:"ctxData"`
	Data    any         `json:"data"`
}

func Send[T any](ctx context.Context, name string, msg T) error {
	var err error
	url := os.Getenv(name + "_URL")
	if url == "" {
		// get queue url
		url, err = GetUrl(name)
		if err != nil {
			return err
		}
	}

	body, err := json.Marshal(queueMsg{
		CtxData: app.ReadCtx(ctx),
		Data:    msg,
	})
	if err != nil {
		return err
	}
	cl := client()

	res, err := cl.SendMessage(
		ctx,
		&sqs.SendMessageInput{
			MessageBody: aws.String(string(body)),
			QueueUrl:    &url,
		})

	if res != nil {
		log.Info("message to queue %s, messageId :%s", url, *res.MessageId)
	}

	return err
}

// Unmarshal queue message with app context data loading
func Unmarshal[T any](ctx context.Context, body string) (context.Context, T, error) {
	var msg queueMsg
	var res T
	if err := json.Unmarshal([]byte(body), &msg); err != nil {
		return ctx, res, err
	}

	newCtx := app.SetCtx(ctx, msg.CtxData)

	d, err := json.Marshal(msg.Data)
	if err != nil {
		return newCtx, res, err
	}

	err = json.Unmarshal(d, &res)
	return newCtx, res, err
}

func GetUrl(name string) (string, error) {
	ctx := context.Background()
	cl := client()
	res, err := cl.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: &name})
	if err != nil {
		return "", err
	}

	url := *res.QueueUrl
	_ = os.Setenv(name+"_URL", url)
	return url, nil
}

func client() *sqs.Client {
	return sqs.NewFromConfig(config.AwsConfig())
}
