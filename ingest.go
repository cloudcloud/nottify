package main

import (
	"database/sql"
	"fmt"

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

func init() {
	cmdIngest.Run = ingestRun
}

func ingestRun(args []string) {
	mysqlDsn := revel.Config.StringDefault("nottify.mysql_dsn", "derp:@/nottify")

	_, err := sql.Open("mysql", mysqlDsn)
	fmt.Println(err)

	//queryResult, err := db.Query("SELECT * FROM songs")
	//fmt.Println(queryResult)

	//fmt.Println(err)
}
