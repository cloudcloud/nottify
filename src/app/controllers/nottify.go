package controllers

import (
	"strconv"

	"code.google.com/p/go-sqlite/go1/sqlite3"
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

	//nott := nottify.LoadConnection()

	return c.Render()
}
