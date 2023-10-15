/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/aws/aws-lambda-go/lambda"
	"roofix/config"
	"roofix/ent"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context) error {
	config.PrintBuildInfo()
	db := ent.GetClient()
	defer db.CloseClient()

	// clear socket connections older than 3 minutes
	_, err := db.UserSession.Delete().Where(func(s *sql.Selector) {
		s.Where(sql.ExprP("expires_at < UTC_TIMESTAMP"))
	}).Exec(ctx)
	return err
}
