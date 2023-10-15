/*
 * Copyright (c) 2022. SimSaw Software Pvt. Ltd. All Right Reserved.
 * Author: Ankit Patial
 * Date: 4/1/22, 12:02 PM
 */

package main

import (
	"context"
	"flag"
	"roofix/bin/migrate/schema"
	"roofix/bin/migrate/seed"
	"roofix/config"
	"roofix/ent"
	"roofix/pkg/util/log"
)

var (
	adminUsers = []ent.User{
		{
			Email:     "admin@simsaw.com",
			FirstName: "Admin",
			LastName:  "User",
			Pwd:       "password",
		},
	}
)

func main() {
	config.PrintBuildInfo()

	var withAdmin, withPostal, withPricing bool
	flag.BoolVar(&withAdmin, "with-admin", false, "import default admins mentioned in code")
	flag.BoolVar(&withPostal, "with-postal", false, "import default postal data")
	flag.BoolVar(&withPricing, "with-pricing", false, "import default pricing data")
	flag.Parse()

	// Run migration.
	ctx := context.Background()

	log.Info("migrating schema...")
	if err := schema.Migrate(ctx); err != nil {
		log.Fatal("migrate schema", err)
	} else {
		log.Info("migrating schema, DONE!")
	}

	// quickScript(ctx)
	if withAdmin {
		log.Info("migrating admin users...")
		if err := schema.SaveAdminUsers(ctx, adminUsers); err != nil {
			log.Error(err)
		}
		log.Info("migrating admin users, DONE!")
	}

	if withPostal {
		seed.USPostal(ctx)
	}

	if withPricing {
		seed.USPricing(ctx)
	}

	//schema.MigrateEstimate(ctx)

	seed.OptionList(ctx)
	seed.Partners(ctx)
}

func quickScript(ctx context.Context) {
	log.Info("=== UpdateScript ===>")
	// From ent.Client.
	db := ent.SqlDB()
	qry := ``
	if _, err := db.ExecContext(ctx, qry); err != nil {
		log.Error(err)
	}
}

func estChange(ctx context.Context) {
	log.Info("=== UpdateScript ===>")
	// From ent.Client.
	db := ent.SqlDB()
	qry := `
alter table estimates change provider estimator varchar(50) not null;
alter table estimates change order_id estimator_order_id bigint unsigned default '0' null;
alter table estimates change report_id estimator_report_id bigint unsigned default '0' null;
drop table job_status;
alter table jobs change partner_price work_order_price decimal(10, 2) default 0.00 not null;
`
	if _, err := db.ExecContext(ctx, qry); err != nil {
		log.Error(err)
	}
}
