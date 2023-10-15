/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 */

package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/bin/migrate/schema"
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/util/log"
)

func main() {
	lambda.Start(Handler)
}

func Handler(_ context.Context, _ events.APIGatewayProxyRequest) (string, error) {
	config.PrintBuildInfo()

	ctx := context.Background()

	//beforeScript(ctx)

	if err := schema.Migrate(ctx); err != nil {
		return "", err
	} else {
		log.Info("migrating schema, DONE!")
	}

	//seed.USPostal(ctx)
	//seed.USPricing(ctx)
	//seed.OptionList(ctx)
	//seed.Partners(ctx)

	//schema.MigrateEstimate(ctx)

	// afterScript(ctx)
	return "Done!", nil
}

func beforeScript(ctx context.Context) {
	log.Info("=== Before Script ===>")
	qry := ``
	if _, err := ent.SqlDB().ExecContext(ctx, qry); err != nil {
		log.Error(err)
	}
}

func afterScript(ctx context.Context) {
	log.Info("=== After Script ===>")
	qry := ``
	if _, err := ent.SqlDB().ExecContext(ctx, qry); err != nil {
		log.Error(err)
	}
}
