package main

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
	"path/filepath"
	"roofix/cdk/infra"
	"roofix/cdk/sm"
	"roofix/config"
	"roofix/pkg/enum"
	"strings"
)

var suffix string

func main() {
	defer jsii.Close()
	env := os.Getenv("STACK_ENV")
	env = strings.Trim(env, " ")

	appEnv := "stage"
	if "prod" == strings.ToLower(env) || "production" == strings.ToLower(env) {
		appEnv = "production"
		suffix = "-prod"
	}

	app := awscdk.NewApp(nil)
	buildStack(app, appEnv)
	app.Synth(nil)
}

// buildStack for given env
func buildStack(app constructs.Construct, appEnv string) {
	conf := config.ReadByEnv(appEnv).Infra
	stackEnv := &awscdk.Environment{
		Account: jsii.String(conf.Stack.Account),
		Region:  jsii.String(conf.Stack.Region),
	}

	stack := awscdk.NewStack(app, &conf.Stack.Name, &awscdk.StackProps{
		Env:       stackEnv,
		StackName: jsii.String(conf.Stack.Name),
	})

	// VPC
	vpc := infra.NewVpc(stack)
	backend := []awsec2.ISecurityGroup{vpc.SgBackend}
	backendWithRDS := []awsec2.ISecurityGroup{vpc.SgBackend, vpc.SgRDS}

	// RDS
	insCls := awsec2.InstanceClass_BURSTABLE3
	insSize := awsec2.InstanceSize_MEDIUM
	cluster, proxyEP := infra.NewAuroraCluster(stack, conf, appEnv, vpc.VPC, vpc.SgRDS, insCls, insSize)
	cluster.AddRotationSingleUser(&awsrds.RotationSingleUserOptions{
		AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(90)),
		ExcludeCharacters:  jsii.String("!@#$%^&*"),
	})

	// SQS
	taskQ := infra.NewQueue(stack, conf.TaskQueue, 14, 60*15)
	mailQ := infra.NewQueue(stack, conf.MailQueue, 14, 60)
	notificationQ := infra.NewQueue(stack, conf.NotificationQueue, 14, 60)
	//dataSyncQ := infra.NewFIFOQueue(stack, conf.DataSyncQueue, 14, 60)

	// SNS
	feedbackTopic := infra.NewTopic(stack, "mail-feedback"+suffix, "Mail Feedback")

	// SES
	infra.NewSesConfigSet(stack, "mail-config"+suffix, feedbackTopic.TopicArn())

	// watch event rules
	//evtFiveMin := infra.NewRateRule(stack, "five-minutes"+suffix, "repeat after 5 minutes", awscdk.Duration_Minutes(jsii.Number(5)))
	//threeHrs := infra.NewRateRule(stack, "three-hrs"+suffix, "repeat after 3 hours", awscdk.Duration_Hours(jsii.Number(3)))
	evtDawnDusk := infra.NewCronRule(stack, "cst-dawn-dusk"+suffix, "early morning and late evening", awsevents.CronOptions{
		Hour:   jsii.String("2,10"), // 2(9pm CST), 10(5am CST),
		Minute: jsii.String("0"),
	})

	evtMidNight := infra.NewCronRule(stack, "cst-midnight"+suffix, "midnight", awsevents.CronOptions{
		Hour:   jsii.String("5"), // 5(12pm CST)
		Minute: jsii.String("0"),
	})

	// LAMBDA, variables
	oneMin := awscdk.Duration_Minutes(jsii.Number(1))
	fifteenMin := awscdk.Duration_Minutes(jsii.Number(15))
	hostedZone := infra.HostedZone(stack, conf)

	// Certs
	siteCert := infra.NewCertificate(stack, "website-cert", conf.WebsiteDomain, hostedZone)
	wsCert := infra.NewCertificate(stack, "ws-cert", conf.WsDomain, hostedZone)
	assetCert := infra.NewCertificate(stack, "assets-cert", conf.AssetsDomain, hostedZone)

	// LAMBDA functions
	fnRole := infra.LambdaFnRole(stack, "lambda-fn-role"+suffix, conf)
	arm := awslambda.Architecture_ARM_64()
	// clearExpiredJobInvitesFn
	clearExpiredJobInvitesFn := infra.Lambda(
		stack, appEnv, "clearExpiredJobInvites", "clearExpiredJobInvitesFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, nil,
	)
	evtDawnDusk.AddTarget(awseventstargets.NewLambdaFunction(clearExpiredJobInvitesFn, nil))

	// clearExpiredSessionsFn
	clearExpiredSessions := infra.Lambda(
		stack, appEnv, "clearExpiredSessions", "clearExpiredSessionsFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, nil,
	)
	evtMidNight.AddTarget(awseventstargets.NewLambdaFunction(clearExpiredSessions, nil))

	// dataSyncFn
	dataSyncFn := infra.Lambda(
		stack, appEnv, "dataSync", "dataSyncFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, nil,
	)
	//dataSyncQ.GrantConsumeMessages(dataSyncFn)
	dataSyncFn.AddEventSource(awslambdaeventsources.NewSqsEventSource(mailQ, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize:               jsii.Number(10),
		Enabled:                 jsii.Bool(true),
		ReportBatchItemFailures: nil,
	}))

	// docUploadFn
	docUploadFn := infra.Lambda(
		stack, appEnv, "docUpload", "docUploadFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backend, nil,
	)

	// mailerFn
	mailerFn := infra.Lambda(
		stack, appEnv, "mailer", "mailerFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backend, map[string]*string{
			fmt.Sprintf("%s_URL", conf.MailQueue): mailQ.QueueUrl(),
		},
	)
	mailQ.GrantConsumeMessages(mailerFn)
	mailerFn.AddEventSource(awslambdaeventsources.NewSqsEventSource(mailQ, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize:               jsii.Number(10),
		Enabled:                 jsii.Bool(true),
		ReportBatchItemFailures: nil,
	}))

	// mailFeedbackFn
	mailFeedbackFn := infra.Lambda(
		stack, appEnv, "mailFeedback", "mailFeedbackFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backend, nil,
	)
	feedbackTopic.AddSubscription(awssnssubscriptions.NewLambdaSubscription(
		mailFeedbackFn,
		&awssnssubscriptions.LambdaSubscriptionProps{
			DeadLetterQueue: nil,
			FilterPolicy:    nil,
		},
	))

	// migrateFn
	infra.Lambda(
		stack, appEnv, "migrate", "migrateFn", arm, 1024, awscdk.Duration_Minutes(jsii.Number(5)),
		fnRole, vpc.VPC, &backendWithRDS, map[string]*string{
			"DB_PROXY": proxyEP,
		},
	)

	// notificationFn
	notificationFn := infra.Lambda(
		stack, appEnv, "notification", "notificationFn", arm, 128, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, map[string]*string{
			fmt.Sprintf("%s_URL", conf.NotificationQueue): notificationQ.QueueUrl(),
			"DB_PROXY": proxyEP,
		},
	)
	notificationQ.GrantConsumeMessages(notificationFn)
	notificationFn.AddEventSource(awslambdaeventsources.NewSqsEventSource(notificationQ, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize:               jsii.Number(5),
		Enabled:                 jsii.Bool(true),
		ReportBatchItemFailures: nil,
	}))

	// taskFn
	taskFn := infra.Lambda(
		stack, appEnv, "task", "taskFn", arm, 256, fifteenMin,
		fnRole, vpc.VPC, &backendWithRDS, map[string]*string{
			fmt.Sprintf("%s_URL", conf.TaskQueue):         taskQ.QueueUrl(),
			fmt.Sprintf("%s_URL", conf.MailQueue):         mailQ.QueueUrl(),
			fmt.Sprintf("%s_URL", conf.NotificationQueue): notificationQ.QueueUrl(),
			"DB_PROXY": proxyEP,
		},
	)
	taskFn.AddToRolePolicy(infra.PolicyStmtApiExec())
	taskQ.GrantSendMessages(taskFn)
	taskQ.GrantConsumeMessages(taskFn)
	taskFn.AddEventSource(awslambdaeventsources.NewSqsEventSource(taskQ, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize:               jsii.Number(5),
		Enabled:                 jsii.Bool(true),
		ReportBatchItemFailures: nil,
	}))
	mailQ.GrantSendMessages(taskFn)
	notificationQ.GrantSendMessages(taskFn)

	// websiteFn
	infra.ExecCmd(fmt.Sprintf("make domain=%s mode=%s build-assets", conf.AssetsDomain, appEnv))
	websiteFn := infra.Lambda(
		stack, appEnv, "website", "websiteFn", arm, 256, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, map[string]*string{
			fmt.Sprintf("%s_URL", conf.TaskQueue):         taskQ.QueueUrl(),
			fmt.Sprintf("%s_URL", conf.MailQueue):         mailQ.QueueUrl(),
			fmt.Sprintf("%s_URL", conf.NotificationQueue): notificationQ.QueueUrl(),
			//fmt.Sprintf("%s_URL", conf.DataSyncQueue):     dataSyncQ.QueueUrl(),
			"DB_PROXY": proxyEP,
		},
	)
	infra.RestApi(stack, "website", conf.WebsiteDomain, websiteFn, siteCert, hostedZone)
	taskQ.GrantSendMessages(websiteFn)         // task Queue Access
	mailQ.GrantSendMessages(websiteFn)         // mail Queue Access
	notificationQ.GrantSendMessages(websiteFn) // notification Queue Access
	//dataSyncQ.GrantSendMessages(websiteFn)     // notification Queue Access

	// wsApiFn
	wsApiFn := infra.Lambda(
		stack, appEnv, "wsApi", "wsApiFn", arm, 256, oneMin,
		fnRole, vpc.VPC, &backendWithRDS, map[string]*string{
			"DB_PROXY": proxyEP,
		},
	)
	infra.NewWsApi(stack, stackEnv, "ws"+suffix, conf.WsDomain, wsApiFn, wsCert, hostedZone)

	// STEP-Functions
	sm.HtmlToPdf(stack, appEnv, fnRole, vpc.VPC, vpc.SgBackend)

	// S3, data bucket
	dataBucket := infra.NewBucket(
		stack,
		conf.DataBucket,
		&[]*awss3.CorsRule{
			{
				AllowedOrigins: &[]*string{aws.String(fmt.Sprintf("https://%s", conf.WebsiteDomain))},
				AllowedMethods: &[]awss3.HttpMethods{awss3.HttpMethods_PUT},
				AllowedHeaders: &[]*string{aws.String("*")},
				MaxAge:         aws.Float64(3000),
			},
		})
	// secure files
	secrets := fmt.Sprintf("%s/*", enum.DirSecrets)
	dataBucket.GrantRead(taskFn, secrets)
	dataBucket.GrantRead(websiteFn, secrets)
	dataBucket.GrantPut(websiteFn, secrets)
	// email templates
	mails := fmt.Sprintf("%s/*", enum.DirEmailTemplates)
	dataBucket.GrantRead(mailerFn, mails)
	dataBucket.GrantRead(websiteFn, mails)
	// pdf templates
	pdfs := fmt.Sprintf("%s/*", enum.DirPdfTemplates)
	dataBucket.GrantRead(taskFn, pdfs)

	// data folders
	dataBucket.GrantReadWrite(websiteFn, fmt.Sprintf("%s/*", enum.FolderEstimates))
	dataBucket.GrantReadWrite(websiteFn, fmt.Sprintf("%s/*", enum.FolderJobDocs))
	dataBucket.GrantReadWrite(websiteFn, fmt.Sprintf("%s/*", enum.FolderPartnerDocs))
	dataBucket.GrantReadWrite(websiteFn, fmt.Sprintf("%s/*", enum.FolderTrainingVideos))
	dataBucket.GrantDelete(websiteFn, fmt.Sprintf("%s/*", enum.FolderTrainingVideos))
	dataBucket.GrantPut(websiteFn, fmt.Sprintf("%s/*", enum.FolderPublicData))
	dataBucket.GrantDelete(websiteFn, fmt.Sprintf("%s/*", enum.FolderPublicData))
	// data upload notification
	dataBucket.AddEventNotification(awss3.EventType_OBJECT_CREATED_PUT, awss3notifications.NewLambdaDestination(docUploadFn))

	// S3, assets bucket and deployment
	assetBucket := infra.NewBucket(stack, conf.AssetsBucket, nil)
	wd, _ := os.Getwd()
	assetsDir := filepath.Join(wd, "website", "build")
	infra.NewBucketDeployment(stack, assetBucket, assetsDir)
	// NOTE: data bucket permission policy for CDN is only on "PublicData/*", need change it manually
	infra.NewDistributionMultiOrigin(
		stack, "assets"+suffix, conf.AssetsDomain, jsii.String(conf.Stack.Region),
		assetBucket, dataBucket, fmt.Sprintf("%s/*", enum.FolderPublicData), assetCert, hostedZone,
	)
	assetBucket.GrantRead(websiteFn, "*")

	// S3, log bucket
	logBucket := infra.NewBucket(stack, conf.LogBucket, nil)
	logBucket.GrantPut(mailFeedbackFn, "mail/*")
	logBucket.GrantRead(mailerFn, "mail/*")
	// dir: error-logs
	errLogs := fmt.Sprintf("%s/*", enum.DirErrorLogs)
	logBucket.GrantRead(websiteFn, errLogs)

}
