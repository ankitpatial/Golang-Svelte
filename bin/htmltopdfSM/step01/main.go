/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 * Description: lambda to render html to pdf
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/bin/htmltopdfSM"
	"roofix/config"
	"roofix/template"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, inp htmltopdfSM.InputRenderHtml) (*htmltopdfSM.OutputRenderHtml, error) {
	config.PrintBuildInfo()

	// render html with data to print
	inp.Data["title"] = inp.Title
	html, err := template.Exec(ctx, template.KindPDF, inp.Tpl, inp.Data)
	if err != nil {
		return nil, err
	}

	return &htmltopdfSM.OutputRenderHtml{
		Bucket: inp.Dest.Bucket,
		Key:    inp.Dest.ObjectKey(),
		Dest:   inp.Dest,
		Html:   html,
	}, err
}
