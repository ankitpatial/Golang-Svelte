package htmltopdfSM

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/jsii-runtime-go"
	"roofix/cdk/infra"
	"roofix/config"
	"roofix/pkg/util/log"
	"roofix/template"
	"time"
)

const Name = "html-to-pdf"

func Exec(ctx context.Context, dest Destination, tplName template.Name, tplTitle string, tplData map[string]interface{}) {
	if config.IsDevEnv() {
		log.Info("dev env has no way to execute this State machine")
		return
	}

	log.InfoBullet("Generate Price PDF For JobID: %s", dest.Dir)

	s := config.Read().Infra.Stack
	client := sfn.NewFromConfig(config.AwsConfig())
	arn := fmt.Sprintf("arn:aws:states:%s:%s:stateMachine:%s", s.Region, s.Account, infra.SuffixIt(config.AppEnv, Name))
	data := InputRenderHtml{
		Tpl:   tplName,
		Title: tplTitle,
		Data:  tplData,
		Dest:  dest,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
	}

	name := fmt.Sprintf("%s-%s", dest.Dir, time.Now().UTC().Format("200601021504"))
	response, err := client.StartExecution(ctx, &sfn.StartExecutionInput{
		Name:            &name,
		StateMachineArn: &arn,
		Input:           jsii.String(string(b)),
	})
	if err != nil {
		log.Error(err)
	} else {
		log.InfoBullet("State Machine ExecutionArn: %s", response.ExecutionArn)
	}
}
