/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"roofix/config"
)

// NewAuroraCluster will create Aurora RDS cluster
func NewAuroraCluster(
	stack constructs.Construct,
	conf *config.Infra,
	appEnv string,
	vpc awsec2.IVpc,
	sg awsec2.ISecurityGroup,
	class awsec2.InstanceClass,
	size awsec2.InstanceSize,
) (awsrds.DatabaseCluster, *string) {
	user := jsii.String(conf.DB.Username)
	sgGroups := []awsec2.ISecurityGroup{
		sg,
	}
	// instance
	instanceType := awsec2.InstanceType_Of(class, size)
	secret := awsrds.Credentials_FromGeneratedSecret(user, &awsrds.CredentialsBaseOptions{
		SecretName: jsii.String(conf.DB.SecretName),
	})

	// db cluster
	cluster := awsrds.NewDatabaseCluster(stack, jsii.String("database"), &awsrds.DatabaseClusterProps{
		Engine: awsrds.DatabaseClusterEngine_AuroraMysql(&awsrds.AuroraMysqlClusterEngineProps{
			Version: awsrds.AuroraMysqlEngineVersion_VER_3_02_2(),
		}),
		Backup: &awsrds.BackupProps{
			Retention:       awscdk.Duration_Days(jsii.Number(30)),
			PreferredWindow: jsii.String("05:00-06:00"), // cst midnight
		},
		Credentials:         secret,
		DefaultDatabaseName: jsii.String(conf.DB.Name),
		DeletionProtection:  jsii.Bool(true),
		RemovalPolicy:       awscdk.RemovalPolicy_RETAIN,
		SecurityGroups:      &sgGroups,
		StorageEncrypted:    jsii.Bool(true),
		Vpc:                 vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
		Writer: awsrds.ClusterInstance_Provisioned(jsii.String("writer"), &awsrds.ProvisionedClusterInstanceProps{
			AutoMinorVersionUpgrade: jsii.Bool(true),
			//EnablePerformanceInsights:   jsii.Bool(true),
			//PerformanceInsightRetention: awsrds.PerformanceInsightRetention_MONTHS_3,
			PubliclyAccessible: jsii.Bool(false),
			InstanceType:       instanceType,
		}),
		Readers: nil,
	})

	secrets := []awssecretsmanager.ISecret{
		cluster.Secret(),
	}

	proxy := awsrds.NewDatabaseProxy(stack, jsii.String(SuffixIt(appEnv, "proxy")), &awsrds.DatabaseProxyProps{
		Secrets:                   &secrets,
		Vpc:                       vpc,
		IdleClientTimeout:         awscdk.Duration_Seconds(jsii.Number(3)),
		MaxConnectionsPercent:     jsii.Number(80),
		MaxIdleConnectionsPercent: jsii.Number(20),
		SecurityGroups:            &sgGroups,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
		ProxyTarget: awsrds.ProxyTarget_FromCluster(cluster),
	})

	role := awsiam.NewRole(stack, jsii.String(SuffixIt(appEnv, "proxy-role")), &awsiam.RoleProps{
		AssumedBy: awsiam.NewAccountPrincipal(conf.Stack.Account),
	})
	proxy.GrantConnect(role, user)
	return cluster, proxy.Endpoint()
}
