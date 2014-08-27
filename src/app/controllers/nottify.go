package controllers

import (
	"github.com/revel/revel"
	"strconv"
)

type Nottify struct {
	*revel.Controller
}

func (c Nottify) Home() revel.Result {
	actualPin, _ := strconv.Atoi(c.Session["pin"])

	if actualPin != revel.Config.IntDefault("nottify.pin_code", 55555) {
		return c.Redirect(App.Index)
	}

	return c.Render()
}

