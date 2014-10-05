package controllers

import (
	"os"
	"strconv"

	"github.com/cloudcloud/nottify/src/nottify"
	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

func (a Api) Song() revel.Result {
	actualPin, _ := strconv.Atoi(a.Session["pin"])
	uuid := a.Params.Get("uuid")

	if actualPin != revel.Config.IntDefault("nottify.pin_code", 55555) {
		return a.Redirect(App.Index)
	}

	nott := nottify.LoadConnection()

	filename := nott.GetFilename(uuid)
	file, err := os.Open(filename)

	if err != nil {
		return a.Redirect(Nottify.Home)
	}

	return a.RenderFile(file, "inline")
}
