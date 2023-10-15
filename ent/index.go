package ent

import (
	"database/sql"
	ensql "entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	appConf "roofix/config"
	"roofix/pkg/util/log"
	"strings"
)

var dbConnStr, dbEngin string

func init() {
	cfg := appConf.Read().DB
	dbEngin = cfg.Engine
	dbConnStr = cfg.ConnStr()
	if dbConnStr == "" {
		panic(errors.New("unable to find database connection string"))
	}
}

func SqlDB() *sql.DB {
	drv, err := ensql.Open(dbEngin, dbConnStr)
	if err != nil {
		log.Error(err)
		log.Fatal("db connect failed with error", err)
	}
	return drv.DB()
}

func GetClient() *Client {
	// setup conn
	drv, err := ensql.Open(dbEngin, dbConnStr)
	if err != nil {
		log.Error(err)
		log.Fatal("db connect failed with error", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	//db.SetConnMaxLifetime(time.Second)

	// cache client
	return NewClient(Driver(drv))
}

// CloseClient will call close db connection
// https://go.dev/doc/database/manage-connections
// close connection call makes more sense when we are using dedicated connection
// e.g. "BeginTx, ExecContext, PingContext, PrepareContext, QueryContext, and QueryRowContext"
func (c *Client) CloseClient() {
	if err := c.driver.Close(); err != nil {
		log.Error(err)
	}
}

func (u *User) Name() string {
	return strings.Join(strings.Fields(fmt.Sprintf("%s %s ", u.FirstName, u.LastName)), " ")
}
