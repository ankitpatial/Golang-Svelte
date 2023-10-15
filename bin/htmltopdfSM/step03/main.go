/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 *
 * Description: lambda to save document info to db
 */

package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/bin/htmltopdfSM"
	"roofix/config"
	"roofix/pkg/document"
	"roofix/pkg/util/log"
)

type Input struct {
	Payload string `json:"Payload"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, e Input) (string, error) {
	config.PrintBuildInfo()

	var p htmltopdfSM.OutputRenderPDF
	if err := json.Unmarshal([]byte(e.Payload), &p); err != nil {
		return "ERROR", err
	}

	// create a document entry
	inp := &document.Input{
		Folder:      p.Folder,
		Dir:         p.Dir,
		Section:     p.Section,
		Name:        p.FileName,
		FileName:    p.FileName,
		ContentType: &p.ContentType,
		ContentSize: p.ContentSize,
	}

	// save document to db
	_, err := document.Save(ctx, p.Bucket, p.Key, inp, true)
	if err != nil {
		log.Info("FAILED, document.Save")
		return "FAILED", err
	}

	return "SUCCESS", nil
}
