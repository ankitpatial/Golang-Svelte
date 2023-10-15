package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

func NewTopic(stack constructs.Construct, name, displayName string) awssns.ITopic {
	return awssns.NewTopic(stack, &name, &awssns.TopicProps{
		TopicName:   &name,
		DisplayName: &displayName,
	})
}
