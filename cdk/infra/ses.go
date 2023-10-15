package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSesConfigSet(stack constructs.Construct, name string, feedbackTopicArn *string) {
	awsses.NewConfigurationSet(stack, &name, &awsses.ConfigurationSetProps{
		ConfigurationSetName: &name,
		ReputationMetrics:    jsii.Bool(true),
		SendingEnabled:       jsii.Bool(true),
		TlsPolicy:            awsses.ConfigurationSetTlsPolicy_REQUIRE,
	})

	n := name + "event"
	awsses.NewCfnConfigurationSetEventDestination(stack, &n, &awsses.CfnConfigurationSetEventDestinationProps{
		ConfigurationSetName: &name,
		EventDestination: &awsses.CfnConfigurationSetEventDestination_EventDestinationProperty{
			Name:    &n,
			Enabled: jsii.Bool(true),
			MatchingEventTypes: &[]*string{
				jsii.String("send"),
				jsii.String("reject"),
				jsii.String("bounce"),
				jsii.String("complaint"),
				jsii.String("delivery"),
			},
			SnsDestination: &awsses.CfnConfigurationSetEventDestination_SnsDestinationProperty{
				TopicArn: feedbackTopicArn,
			},
		},
	})

}
