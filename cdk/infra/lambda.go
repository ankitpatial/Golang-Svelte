package infra

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
	"os/exec"
	"path/filepath"
)

func NewLayerVersion(
	stack constructs.Construct, name, version string, arch []awslambda.Architecture, runtime []awslambda.Runtime,
) awslambda.LayerVersion {
	wd, _ := os.Getwd()
	assetPath := filepath.Join(wd, "bin", name)

	props := &awslambda.LayerVersionProps{
		LayerVersionName:        aws.String(name + "-" + version),
		License:                 aws.String("private"),
		CompatibleArchitectures: &arch,
		CompatibleRuntimes:      &runtime,
		Code:                    awslambda.Code_FromAsset(aws.String(assetPath), nil),
	}

	return awslambda.NewLayerVersion(stack, &name, props)
}

func Lambda(
	stack constructs.Construct,
	appEnv,
	fnName,
	fnPath string,
	arch awslambda.Architecture,
	memory float64,
	timout awscdk.Duration,
	role awsiam.IRole,
	vpc awsec2.Vpc,
	sg *[]awsec2.ISecurityGroup,
	envVars map[string]*string,
) awslambda.Function {
	return NewGoFn(stack, appEnv, SuffixIt(appEnv, fnName), fnPath, arch, memory, timout, role, vpc, sg, envVars)
}

func NewGoFn(
	stack constructs.Construct,
	appEnv,
	fnName,
	fnPath string,
	arch awslambda.Architecture,
	memory float64,
	timout awscdk.Duration,
	role awsiam.IRole,
	vpc awsec2.Vpc,
	sg *[]awsec2.ISecurityGroup,
	envVars map[string]*string,
) awslambda.Function {
	// build go code
	assetPath := buildLambda(appEnv, fnName, fnPath, arch)
	// fn props
	props := &awslambda.FunctionProps{
		Architecture:   arch,
		Environment:    &envVars,
		FunctionName:   jsii.String(fnName),
		LogRetention:   awslogs.RetentionDays_ONE_MONTH,
		MemorySize:     jsii.Number(memory),
		Role:           role,
		SecurityGroups: sg,
		Timeout:        timout,
		Vpc:            vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
		Code:    awslambda.Code_FromAsset(jsii.String(assetPath), nil),
		Handler: jsii.String("bootstrap"),
		Runtime: awslambda.Runtime_PROVIDED_AL2(), // https://docs.amazonaws.cn/en_us/lambda/latest/dg/lambda-golang.html
	}
	// fn
	return awslambda.NewFunction(stack, jsii.String(fnName), props)
}

func NewPythonFn(
	stack constructs.Construct,
	fnName,
	fnPath string,
	arch awslambda.Architecture,
	memory float64,
	timout awscdk.Duration,
	role awsiam.IRole,
	vpc awsec2.Vpc,
	sg *[]awsec2.ISecurityGroup,
) awslambda.Function {
	wd, _ := os.Getwd()
	assetPath := filepath.Join(wd, "bin", fnPath)
	// fn props
	props := &awslambda.FunctionProps{
		FunctionName: jsii.String(fnName),
		Runtime:      awslambda.Runtime_PYTHON_3_8(), // https://docs.amazonaws.cn/en_us/lambda/latest/dg/lambda-golang.html
		Architecture: arch,
		LogRetention: awslogs.RetentionDays_ONE_MONTH,
		Code:         awslambda.Code_FromAsset(jsii.String(assetPath), nil),
		Handler:      jsii.String("main.lambda_handler"),
		Role:         role,
		MemorySize:   jsii.Number(memory),
		Timeout:      timout,
		Vpc:          vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
		},
		SecurityGroups: sg,
	}
	// fn
	return awslambda.NewFunction(stack, jsii.String(fnName), props)
}

func buildLambda(appEnv string, name, path string, arch awslambda.Architecture) string {
	wd, _ := os.Getwd()
	arc := "arm64"
	if arch == awslambda.Architecture_X86_64() {
		arc = "amd64"
	}
	// buildCmd is dependent on Makefile action "build-lambda" defined in root level Makefile
	buildCmd := fmt.Sprintf("make env=%s arch=%s name=%s path=%s build-lambda", appEnv, arc, name, path)
	//fmt.Println("   " + buildCmd)
	ExecCmd(buildCmd)
	// out dir: ./dist
	return filepath.Join(wd, "cdk.out", "build", name)
}

func ExecCmd(buildCmd string) {
	cmd := exec.Command("/bin/sh", "-c", buildCmd)
	out, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println("   " + scanner.Text())
	}
	_ = cmd.Wait()
}
