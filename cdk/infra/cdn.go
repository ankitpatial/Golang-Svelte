/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  17/05/22, 2:45 PM
 */

package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewDistribution(
	stack constructs.Construct,
	name string,
	domain string,
	region *string,
	bucket awss3.IBucket,
	cert awscertificatemanager.ICertificate,
	zone awsroute53.IHostedZone,
) {
	// origin access identity
	oai := awscloudfront.NewOriginAccessIdentity(
		stack,
		jsii.String(name+"-OAI"),
		&awscloudfront.OriginAccessIdentityProps{
			Comment: jsii.String("assets bucket access identity"),
		},
	)

	// response cache policy
	resPolicy := responseHeadersPolicy(stack, name)

	// cloud front distribution
	cdn := distribution(stack, name, oai, resPolicy, domain, region, bucket, cert)

	// create Route53 A record
	if zone != nil {
		awsroute53.NewARecord(stack, jsii.String(name+"RecordSet"), &awsroute53.ARecordProps{
			Zone:       zone,
			RecordName: jsii.String(domain),
			Ttl:        awscdk.Duration_Seconds(jsii.Number(300)),
			Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewCloudFrontTarget(cdn)),
		})
	}
}

func NewDistributionMultiOrigin(
	stack constructs.Construct,
	name string,
	domain string,
	region *string,
	bucket1 awss3.IBucket,
	bucket2 awss3.IBucket,
	bucket2pattern string,
	cert awscertificatemanager.ICertificate,
	zone awsroute53.IHostedZone,
) {
	// origin access identity
	oai := awscloudfront.NewOriginAccessIdentity(
		stack,
		jsii.String(name+"-OAI"),
		&awscloudfront.OriginAccessIdentityProps{
			Comment: jsii.String("assets bucket access identity"),
		},
	)

	// response cache policy
	resPolicy := responseHeadersPolicy(stack, name)

	// cloud front distribution
	dist := distribution(stack, name, oai, resPolicy, domain, region, bucket1, cert)

	origin2 := awscloudfrontorigins.NewS3Origin(bucket2, &awscloudfrontorigins.S3OriginProps{
		OriginShieldRegion:   region,
		OriginAccessIdentity: oai,
	})

	dist.AddBehavior(jsii.String(bucket2pattern), origin2, &awscloudfront.AddBehaviorOptions{
		AllowedMethods:        awscloudfront.AllowedMethods_ALLOW_GET_HEAD(),
		Compress:              jsii.Bool(true),
		CachePolicy:           awscloudfront.CachePolicy_CACHING_OPTIMIZED(),
		ResponseHeadersPolicy: resPolicy,
		SmoothStreaming:       jsii.Bool(true),
		ViewerProtocolPolicy:  awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
	})

	// create Route53 A record
	if zone != nil {
		awsroute53.NewARecord(stack, jsii.String(name+"RecordSet"), &awsroute53.ARecordProps{
			Zone:       zone,
			RecordName: jsii.String(domain),
			Ttl:        awscdk.Duration_Seconds(jsii.Number(300)),
			Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewCloudFrontTarget(dist)),
		})
	}
}

func responseHeadersPolicy(stack constructs.Construct, name string) awscloudfront.ResponseHeadersPolicy {
	return awscloudfront.NewResponseHeadersPolicy(stack, jsii.String(name+"-response-headers"), &awscloudfront.ResponseHeadersPolicyProps{
		ResponseHeadersPolicyName: jsii.String(name + "-response-policy"),
		// CorsBehavior: &awscloudfront.ResponseHeadersCorsBehavior{
		//	AccessControlAllowOrigins:     jsii.Strings(scopeProps.WebsiteDomainName),
		//	AccessControlAllowHeaders:     jsii.Strings("*"),
		//	AccessControlAllowMethods:     jsii.Strings("*"),
		//	AccessControlExposeHeaders:    jsii.Strings("*"),
		//	AccessControlMaxAge:           awscdk.Duration_Seconds(jsii.Number(600)),
		//	AccessControlAllowCredentials: jsii.Bool(false),
		//	OriginOverride:                jsii.Bool(false),
		//},
		CustomHeadersBehavior: &awscloudfront.ResponseCustomHeadersBehavior{
			CustomHeaders: &[]*awscloudfront.ResponseCustomHeader{
				{
					Header:   jsii.String("Cache-Control"),
					Value:    jsii.String("public, immutable, max-age=31536000"),
					Override: jsii.Bool(false),
				},
			},
		},
	})
}

func distribution(
	stack constructs.Construct,
	name string,
	oai awscloudfront.IOriginAccessIdentity,
	resPolicy awscloudfront.IResponseHeadersPolicy,
	domain string,
	region *string,
	bucket awss3.IBucket,
	cert awscertificatemanager.ICertificate,
) awscloudfront.Distribution {
	return awscloudfront.NewDistribution(stack, jsii.String(name+"-cdn"), &awscloudfront.DistributionProps{
		PriceClass:             awscloudfront.PriceClass_PRICE_CLASS_100,
		DomainNames:            jsii.Strings(domain),
		Certificate:            cert,
		Enabled:                jsii.Bool(true),
		HttpVersion:            awscloudfront.HttpVersion_HTTP2,
		MinimumProtocolVersion: awscloudfront.SecurityPolicyProtocol_TLS_V1_2_2021,
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			Origin: awscloudfrontorigins.NewS3Origin(bucket, &awscloudfrontorigins.S3OriginProps{
				OriginShieldRegion:   region,
				OriginAccessIdentity: oai,
			}),
			AllowedMethods:        awscloudfront.AllowedMethods_ALLOW_GET_HEAD(),
			Compress:              jsii.Bool(true),
			CachePolicy:           awscloudfront.CachePolicy_CACHING_OPTIMIZED(),
			ResponseHeadersPolicy: resPolicy,
			SmoothStreaming:       jsii.Bool(true),
			ViewerProtocolPolicy:  awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
	})
}
