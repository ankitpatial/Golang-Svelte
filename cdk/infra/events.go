/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewRateRule(stack constructs.Construct, name, desc string, d awscdk.Duration) awsevents.Rule {
	return awsevents.NewRule(stack, &name, &awsevents.RuleProps{
		RuleName:    &name,
		Description: &desc,
		Enabled:     jsii.Bool(true),
		Schedule:    awsevents.Schedule_Rate(d),
	})
}

func NewCronRule(stack constructs.Construct, name, desc string, opt awsevents.CronOptions) awsevents.Rule {
	return awsevents.NewRule(stack, &name, &awsevents.RuleProps{
		RuleName:    &name,
		Description: &desc,
		Enabled:     jsii.Bool(true),
		Schedule:    awsevents.Schedule_Cron(&opt),
	})
}
