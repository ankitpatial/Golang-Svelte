/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcData struct {
	VPC          awsec2.Vpc
	SgHttpAccess awsec2.SecurityGroup
	SgBackend    awsec2.SecurityGroup
	SgSSH        awsec2.SecurityGroup
	SgRDS        awsec2.SecurityGroup
}

func NewVpc(stack constructs.Construct) *VpcData {
	// create VPC
	vpc := awsec2.NewVpc(stack, jsii.String("vpc"), &awsec2.VpcProps{
		Cidr:                   jsii.String("10.0.0.0/21"),
		DefaultInstanceTenancy: awsec2.DefaultInstanceTenancy_DEFAULT,
		EnableDnsHostnames:     nil,
		EnableDnsSupport:       nil,
		FlowLogs:               nil,
		GatewayEndpoints: &(map[string]*awsec2.GatewayVpcEndpointOptions{
			"S3": {
				Service: awsec2.GatewayVpcEndpointAwsService_S3(),
			},
		}),
		MaxAzs:      jsii.Number(3),
		NatGateways: jsii.Number(1),
		SubnetConfiguration: &([]*awsec2.SubnetConfiguration{
			{
				Name:       jsii.String("public"),
				SubnetType: awsec2.SubnetType_PUBLIC,
				CidrMask:   jsii.Number(24),
			},
			{
				Name:       jsii.String("private"),
				SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
				CidrMask:   jsii.Number(24),
			},
		}),
	})

	// vpc security Groups
	sgHttpAccess := awsec2.NewSecurityGroup(stack, jsii.String("http-access"), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String("http-access"),
		AllowAllOutbound:  jsii.Bool(false),
		Description:       jsii.String("Allow Http & Https traffic"),
	})

	sgHttpAccess.AddIngressRule(
		awsec2.Peer_Ipv4(jsii.String("0.0.0.0/0")),
		awsec2.Port_Tcp(jsii.Number(80)),
		jsii.String("HTTP traffic"),
		nil)

	sgHttpAccess.AddIngressRule(
		awsec2.Peer_Ipv6(jsii.String("::/0")),
		awsec2.Port_Tcp(jsii.Number(443)),
		jsii.String("HTTPS traffic"),
		nil)

	// Backend access
	sgBackend := awsec2.NewSecurityGroup(stack, jsii.String("backend"), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String("backend"),
		AllowAllOutbound:  jsii.Bool(true),
		Description:       jsii.String("Backend servers, must be in restrictive env."),
	})

	// SSH access
	sgSshAccess := awsec2.NewSecurityGroup(stack, jsii.String("ssh-access"), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String("ssh-access"),
		AllowAllOutbound:  jsii.Bool(false),
		Description:       jsii.String("SSH access rules"),
	})

	// RDS access
	sgRdsAccess := awsec2.NewSecurityGroup(stack, jsii.String("rds-access"), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String("rds-access"),
		AllowAllOutbound:  jsii.Bool(false),
		Description:       jsii.String("RDS access rules"),
	})

	// RdsAccess:allow port 3306 for self
	sgRdsAccess.AddIngressRule(
		awsec2.Peer_SecurityGroupId(sgBackend.SecurityGroupId(), nil),
		awsec2.Port_Tcp(jsii.Number(3306)),
		jsii.String("Backend db access"),
		nil)

	return &VpcData{
		VPC:          vpc,
		SgHttpAccess: sgHttpAccess,
		SgBackend:    sgBackend,
		SgSSH:        sgSshAccess,
		SgRDS:        sgRdsAccess,
	}
}
