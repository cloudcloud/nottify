package controllers

import (
    "github.com/revel/revel"
    "regexp"
	"strconv"
)

type App struct {
    *revel.Controller
}

func (c App) Index() revel.Result {
    return c.Render()
}

func (c App) Home() revel.Result {
	return c.Render()
}

func (c App) Login(pin string) revel.Result {
    c.Validation.Required(pin).Message("The PIN is required.")
    c.Validation.MaxSize(pin, 5).Message("Pin should be 5 characters long")
    c.Validation.MinSize(pin, 5).Message("Pin should be 5 characters long")
    c.Validation.Match(pin, regexp.MustCompile("^\\d{5}$")).Message("The Pin needs to be exactly 4 digits")

    actualPin := revel.Config.IntDefault("nottify.pin_code", 55555)
    foundPin, _ := strconv.Atoi(pin)

    c.Validation.Range(foundPin, actualPin, actualPin).Message("Provided Pin is incorrect")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

	c.Session["pin"] = strconv.Itoa(foundPin)
    return c.Redirect(Nottify.Home)
}

