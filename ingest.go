package main

import (
	"database/sql"
	"go/build"
	"path"

	"github.com/cloudcloud/nottify/src/nottify"
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
	gopath  string
	not_obj *nottify.Nottify
)

func init() {
	cmdIngest.Run = ingestRun
}

func ingestRun(args []string) {
	gopath := build.Default.GOPATH

	revel.ConfPaths = []string{path.Join(gopath, "src/github.com/cloudcloud/nottify/src/conf")}
	config, err := revel.LoadConfig("app.conf")
	if err != nil || config == nil {
		errorf("Failed to Config")
	}

	mysqlDsn, confErr := config.String("nottify.mysql_dsn")
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

	not_obj = nottify.Build(config, db)
	not_obj.LoadDir("/media/files/music/")
}
