package infra

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"roofix/config"
	"roofix/pkg/enum"
)

func LambdaFnRole(stack constructs.Construct, roleName string, conf *config.Infra) awsiam.IRole {
	return awsiam.NewRole(stack, &roleName, &awsiam.RoleProps{
		RoleName:    &roleName,
		AssumedBy:   awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
		Description: jsii.String("lambda fn basic access rules"),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"SecretAccess": policySecretRead(),
			"SendMail":     policySendMail(),
			"CloudFront":   policyCFN(),
			"InvokeStepFn": policyInvokeStepFn(),
			"DumpErr":      policyDumpErr(conf),
		},
		ManagedPolicies: &[]awsiam.IManagedPolicy{
			awsiam.ManagedPolicy_FromManagedPolicyArn(
				stack,
				jsii.String("vpc-exec-access"),
				jsii.String("arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"),
			),
		},
	})
}

func policySecretRead() awsiam.PolicyDocument {
	return newPolicyDocument([]awsiam.PolicyStatement{
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Actions: &([]*string{
				jsii.String("secretsmanager:DescribeSecret"),
				jsii.String("secretsmanager:GetSecretValue"),
			}),
			Effect: awsiam.Effect_ALLOW,
			Resources: &([]*string{
				jsii.String("arn:aws:secretsmanager:*:*:secret:*"),
			}),
		}),
	})
}

func policySendMail() awsiam.PolicyDocument {
	return newPolicyDocument([]awsiam.PolicyStatement{
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &([]*string{
				jsii.String("ses:SendEmail"),
				jsii.String("ses:SendRawEmail"),
			}),
			Resources: &([]*string{
				jsii.String("arn:aws:ses:*:*:identity/*"),
				jsii.String("arn:aws:ses:*:*:configuration-set/*"),
			}),
		}),
	})
}

func policyDumpErr(conf *config.Infra) awsiam.PolicyDocument {
	return newPolicyDocument([]awsiam.PolicyStatement{
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &([]*string{
				jsii.String("s3:PutObject"),
			}),
			Resources: &([]*string{
				jsii.String(fmt.Sprintf("arn:aws:s3:::%s/%s/*", conf.LogBucket, enum.DirErrorLogs)),
			}),
		}),
	})
}

func policyInvokeStepFn() awsiam.PolicyDocument {
	return newPolicyDocument([]awsiam.PolicyStatement{
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &([]*string{
				jsii.String("states:StartExecution"),
				jsii.String("states:DescribeExecution"),
			}),
			Resources: &([]*string{
				jsii.String("arn:aws:states:*:*:stateMachine:*"),
			}),
		}),
	})
}

func policyCFN() awsiam.PolicyDocument {
	return newPolicyDocument([]awsiam.PolicyStatement{
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &([]*string{
				jsii.String("cloudfront:CreateInvalidation"),
			}),
			Resources: &([]*string{
				jsii.String("arn:aws:cloudfront:*:*:distribution/*"),
			}),
		}),
	})
}

func PolicyStmtApiExec() awsiam.PolicyStatement {
	return awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &([]*string{
			jsii.String("execute-api:ManageConnections"),
		}),
		Resources: &([]*string{
			jsii.String("arn:aws:execute-api:*:*:*/*/*/*"),
		}),
	})
}

func newPolicyDocument(stmts []awsiam.PolicyStatement) awsiam.PolicyDocument {
	return awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
		AssignSids: jsii.Bool(true),
		Statements: &stmts,
	})
}
