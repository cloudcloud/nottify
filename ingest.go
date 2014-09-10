package main

import (
	"fmt"
	"go/build"
	"path"

	"github.com/cloudcloud/nottify/src/nottify"
	"github.com/revel/revel"

	"code.google.com/p/go-sqlite/go1/sqlite3"
)

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

	dsn, confErr := config.String("nottify.sqlite_path")
	if confErr == false || dsn == "" {
		errorf("No database details have been defined")
	}

	db, err := sqlite3.Open(dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	setupData(db)

	not_obj = nottify.Build(config, db)
	//not_obj.LoadDir("/media/files/music/")
}

func setupData(db *sqlite3.Conn) {
	var tables []string
	sql := "SELECT name FROM sqlite_master WHERE type='table'"

	for s, e := db.Query(sql); e == nil; e = s.Next() {
		var name string
		s.Scan(&name)
		tables = append(tables, name)
	}
	fmt.Println(tables)

	if len(tables) < 2 {
		sql = "DROP TABLE song"
		db.Exec(sql)

		sql = "CREATE TABLE IF NOT EXISTS song(id TEXT(52), title TEXT, artist TEXT, album TEXT, length INTEGER, genre TEXT, filename TEXT)"
		db.Exec(sql)

		//sql = "CREATE TABLE IF NOT EXISTS playlist"
		//db.Exec(sql)

		//sql = "CREATE TABLE IF NOT EXISTS history"
		//db.Exec(sql)

		//sql = "CREATE TABLE IF NOT EXISTS errors"
		//db.Exec(sql)

		//sql = "CREATE TABLE IF NOT EXISTS "
		//db.Exec(sql)
	}
}
