package infra

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func RestApi(stack constructs.Construct, name, domain string, fn awslambda.IFunction, cert awscertificatemanager.ICertificate, zone awsroute53.IHostedZone) awsapigateway.LambdaRestApi {
	apiName := jsii.String(name)
	apiDomain := jsii.String(domain)
	api := awsapigateway.NewLambdaRestApi(stack, apiName, &awsapigateway.LambdaRestApiProps{
		DomainName: &awsapigateway.DomainNameOptions{
			Certificate: cert,
			DomainName:  apiDomain,
		},
		EndpointConfiguration: &awsapigateway.EndpointConfiguration{
			Types: &[]awsapigateway.EndpointType{
				awsapigateway.EndpointType_EDGE,
			},
		},
		Handler:     fn,
		RestApiName: apiName,
		Description: jsii.String(name + " api hosted on serverless"),
		Proxy:       jsii.Bool(true),
	})

	if zone != nil {
		awsroute53.NewARecord(stack, jsii.String(name+"-record"), &awsroute53.ARecordProps{
			Zone:       zone,
			RecordName: apiDomain,
			Ttl:        awscdk.Duration_Seconds(jsii.Number(300)),
			Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewApiGateway(api)),
		})
	}
	// output API Host address
	awscdk.NewCfnOutput(stack, jsii.String(name+"-cfn"), &awscdk.CfnOutputProps{
		Value: jsii.String(fmt.Sprintf("https://%s", *apiDomain)),
	})

	return api
}

func NewWsApi(
	stack constructs.Construct,
	env *awscdk.Environment,
	name,
	domain string,
	apiFn awslambda.IFunction,
	cert awscertificatemanager.ICertificate,
	zone awsroute53.IHostedZone,
) {
	apiName := jsii.String(name)
	apiDomain := jsii.String(domain)

	api := awsapigatewayv2.NewCfnApi(stack, apiName, &awsapigatewayv2.CfnApiProps{
		Name:                     apiName,
		Description:              jsii.String("websocket api"),
		ProtocolType:             jsii.String("WEBSOCKET"),
		RouteSelectionExpression: jsii.String("${request.body.action}"),
	})

	stageName := "prod"
	stage := awsapigatewayv2.NewCfnStage(stack, jsii.String(name+"-stage"), &awsapigatewayv2.CfnStageProps{
		ApiId:      api.Ref(),
		StageName:  jsii.String(stageName),
		AutoDeploy: jsii.Bool(true),
	})

	integration := awsapigatewayv2.NewCfnIntegration(stack, jsii.String(name+"-integration"), &awsapigatewayv2.CfnIntegrationProps{
		ApiId:             api.Ref(),
		IntegrationType:   jsii.String("AWS_PROXY"),
		IntegrationMethod: jsii.String("POST"),
		IntegrationUri: jsii.String(fmt.Sprintf(
			"arn:aws:apigateway:%s:lambda:path/2015-03-31/functions/%s/invocations",
			*env.Region,
			*apiFn.FunctionArn(),
		)),
		CredentialsArn: awsiam.NewRole(stack, jsii.String(name+"-role"), &awsiam.RoleProps{
			RoleName:    jsii.String(name + "-role"),
			AssumedBy:   awsiam.NewServicePrincipal(jsii.String("apigateway.amazonaws.com"), nil),
			Description: jsii.String("websocket api role"),
			InlinePolicies: &(map[string]awsiam.PolicyDocument{
				"LambdaAccess": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
					Statements: &([]awsiam.PolicyStatement{
						awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
							Sid:    jsii.String("01"),
							Effect: awsiam.Effect_ALLOW,
							Actions: &[]*string{
								jsii.String("lambda:InvokeFunction"),
							},
							Resources: &[]*string{
								jsii.String(fmt.Sprintf("arn:aws:lambda:%s:%s:function:*", *env.Region, *env.Account)),
							},
						}),
					}),
				}),
			}),
		}).RoleArn(),
	})

	// $connect route
	stage.AddDependsOn(awsapigatewayv2.NewCfnRoute(stack, jsii.String("ConnectRoute"), &awsapigatewayv2.CfnRouteProps{
		ApiId:    api.Ref(),
		RouteKey: jsii.String("$connect"),
		Target: jsii.String(fmt.Sprintf(
			"integrations/%s",
			*integration.Ref(),
		)),
	}))

	// $disconnect route
	stage.AddDependsOn(awsapigatewayv2.NewCfnRoute(stack, jsii.String("DisconnectRoute"), &awsapigatewayv2.CfnRouteProps{
		ApiId:    api.Ref(),
		RouteKey: jsii.String("$disconnect"),
		Target: jsii.String(fmt.Sprintf(
			"integrations/%s",
			*integration.Ref(),
		)),
	}))

	// $default route
	stage.AddDependsOn(awsapigatewayv2.NewCfnRoute(stack, jsii.String("DefaultRoute"), &awsapigatewayv2.CfnRouteProps{
		ApiId:    api.Ref(),
		RouteKey: jsii.String("$default"),
		Target: jsii.String(fmt.Sprintf(
			"integrations/%s",
			*integration.Ref(),
		)),
	}))

	// deployment
	awsapigatewayv2.NewCfnDeployment(stack, jsii.String("DeployWsApi"), &awsapigatewayv2.CfnDeploymentProps{
		ApiId:       api.Ref(),
		Description: jsii.String("ws api deployment"),
		StageName:   stage.Ref(),
	})

	// custom API domain
	dn := awsapigateway.NewDomainName(stack, jsii.String(name+"-domain"), &awsapigateway.DomainNameProps{
		Certificate:  cert,
		DomainName:   apiDomain,
		EndpointType: awsapigateway.EndpointType_REGIONAL,
	})

	awsapigatewayv2.NewCfnApiMapping(stack, jsii.String(name+"-domain-mapping"), &awsapigatewayv2.CfnApiMappingProps{
		ApiId:      api.Ref(),
		DomainName: dn.DomainName(),
		Stage:      jsii.String(stageName),
	})

	// create Route53 A record
	if zone != nil {
		awsroute53.NewARecord(stack, jsii.String("WsDomainRecordSet"), &awsroute53.ARecordProps{
			Zone:       zone,
			RecordName: apiDomain,
			Ttl:        awscdk.Duration_Seconds(jsii.Number(300)),
			Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewApiGatewayDomain(dn)),
		})
	}
}
