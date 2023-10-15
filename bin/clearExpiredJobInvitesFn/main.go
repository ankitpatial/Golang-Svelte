/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/pkg/job/unassign"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) error {
	config.PrintBuildInfo()
	return unassign.StaleJobs(ctx)
}
