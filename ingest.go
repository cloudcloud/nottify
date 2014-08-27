package main

import (
	"database/sql"
	"fmt"
	"go/build"
	"path"

	"github.com/revel/revel"
)

import _ "github.com/go-sql-driver/mysql"

var cmdIngest = &Command{
	UsageLine: "ingest",
	Short:     "Complete an ingestion of Media",
	Long: `
Ingest will complete a file system trawl for media files.

For example:
	nottify ingest
`,
}

var (
	gopath string
	Config *revel.MergedConfig
)

func init() {
	cmdIngest.Run = ingestRun
}

func ingestRun(args []string) {
	gopath := build.Default.GOPATH

	revel.ConfPaths = []string{path.Join(gopath, "src/github.com/cloudcloud/nottify/src/conf")}
	Config, err := revel.LoadConfig("app.conf")
	if err != nil || Config == nil {
		errorf("Failed to Config")
	}

	mysqlDsn, confErr := Config.String("nottify.mysql_dsn")
	if confErr == false || mysqlDsn == "" {
		errorf("No database details have been defined")
	}

	db, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	queryResult, err := db.Query("SHOW databases")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(queryResult)
}
