package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"roofix/config"
)

func HostedZone(stack constructs.Construct, setting *config.Infra) awsroute53.IHostedZone {
	if setting.HostedZoneID == "" {
		return nil
	}

	return awsroute53.HostedZone_FromHostedZoneAttributes(
		stack,
		jsii.String(setting.Stack.Name+"-hosted-zone"),
		&awsroute53.HostedZoneAttributes{
			HostedZoneId: jsii.String(setting.HostedZoneID),
			ZoneName:     jsii.String(setting.HostedZone),
		},
	)
}
