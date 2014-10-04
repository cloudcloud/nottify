package controllers

import (
	"go/build"
	"path"
	"strconv"

	"code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/cloudcloud/nottify/src/nottify"
	"github.com/revel/revel"
)

var (
	conf *revel.MergedConfig
	db   *sqlite3.Conn
)

type Nottify struct {
	*revel.Controller
}

func (c Nottify) Home() revel.Result {
	actualPin, _ := strconv.Atoi(c.Session["pin"])

	if actualPin != revel.Config.IntDefault("nottify.pin_code", 55555) {
		return c.Redirect(App.Index)
	}

	conf, db := c.loadConnection()
	nott := nottify.Build(conf, db)

	row := nott.LoadRandom(10)
	return c.Render(row)
}

func (c Nottify) loadConnection() (*revel.MergedConfig, *sqlite3.Conn) {
	gopath := build.Default.GOPATH
	revel.ConfPaths = []string{path.Join(gopath, "src/github.com/cloudcloud/nottify/src/conf")}
	config, err := revel.LoadConfig("app.conf")

	if err != nil || config == nil {
		panic("Failed to Config")
	}

	dsn, confErr := config.String("nottify.sqlite_path")
	if confErr == false || dsn == "" {
		panic("No database details have been defined")
	}

	db, err := sqlite3.Open(dsn)
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	return config, db
}
