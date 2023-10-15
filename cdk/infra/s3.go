package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewBucket(stack constructs.Construct, name string, cors *[]*awss3.CorsRule) awss3.IBucket {
	// bucket to hold app data
	bucket := awss3.NewBucket(stack, jsii.String("bucket-"+name), &awss3.BucketProps{
		BucketName:        jsii.String(name),
		Encryption:        awss3.BucketEncryption_S3_MANAGED,
		EnforceSSL:        jsii.Bool(true),
		BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
		RemovalPolicy:     awscdk.RemovalPolicy_RETAIN,
		Cors:              cors,
	})

	// CFN
	awscdk.NewCfnOutput(stack, jsii.String(name+"-cfn"), &awscdk.CfnOutputProps{
		Value: bucket.BucketName(),
	})

	return bucket
}

func NewBucketDeployment(stack constructs.Construct, bucket awss3.IBucket, assetsPath string) {
	awss3deployment.NewBucketDeployment(
		stack,
		jsii.String("assets-deployment"),
		&awss3deployment.BucketDeploymentProps{
			DestinationBucket: bucket,
			Sources: &[]awss3deployment.ISource{
				awss3deployment.Source_Asset(jsii.String(assetsPath), nil),
			},
		})
}
