package sm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"roofix/cdk/infra"
)

const Name = "html-to-pdf"

func HtmlToPdf(stack constructs.Construct, appEnv string, fnRole awsiam.IRole, vpc awsec2.Vpc, backendSG awsec2.ISecurityGroup) {
	weasyPrint := infra.NewLayerVersion(stack, "weasyPrintLayer", "v1",
		[]awslambda.Architecture{
			awslambda.Architecture_X86_64(),
		},
		[]awslambda.Runtime{
			awslambda.Runtime_PYTHON_3_8(),
		})

	dir := "htmltopdfSM"
	smName := infra.SuffixIt(appEnv, Name)
	arm := awslambda.Architecture_ARM_64()
	x86 := awslambda.Architecture_X86_64()
	sg := []awsec2.ISecurityGroup{
		backendSG,
	}
	d := awscdk.Duration_Minutes(jsii.Number(5))

	infra.LogInfo("=> state machine: " + smName)

	s1 := infra.SuffixIt(appEnv, Name+"-step01")
	s1Path := dir + "/step01"
	step1 := awsstepfunctionstasks.NewLambdaInvoke(stack, jsii.String("render-html"), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: infra.NewGoFn(stack, appEnv, s1, s1Path, arm, 256, d, fnRole, vpc, &sg, nil),
	})

	s2 := infra.SuffixIt(appEnv, Name+"-step02")
	s2Path := dir + "/step02"
	fn := infra.NewPythonFn(stack, s2, s2Path, x86, 256, d, fnRole, vpc, &sg)
	fn.AddLayers(weasyPrint)
	step2 := awsstepfunctionstasks.NewLambdaInvoke(stack, jsii.String("render-pdf"), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: fn,
	})

	s3 := infra.SuffixIt(appEnv, Name+"-step03")
	s3Path := dir + "/step03"
	step3 := awsstepfunctionstasks.NewLambdaInvoke(stack, jsii.String("save-document"), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: infra.NewGoFn(stack, appEnv, s3, s3Path, arm, 256, d, fnRole, vpc, &sg, nil),
	})

	definition := step1.Next(step2).Next(step3)

	stateMachine := awsstepfunctions.NewStateMachine(stack, &smName, &awsstepfunctions.StateMachineProps{
		Definition:       definition,
		StateMachineName: &smName,
		Timeout:          awscdk.Duration_Minutes(jsii.Number(30)),
	})

	awscdk.NewCfnOutput(stack, jsii.String(smName+"-cfn"), &awscdk.CfnOutputProps{
		Value: stateMachine.StateMachineArn(),
	})
}
