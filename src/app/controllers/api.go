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
	if !nottify.CheckPin(actualPin) {
		return a.Redirect(App.Index)
	}

	uuid := a.Params.Get("uuid")
	nott := nottify.LoadConnection()

	filename := nott.GetFilename(uuid)
	file, err := os.Open(filename)

	if err != nil {
		return a.Redirect(Nottify.Home)
	}

	return a.RenderFile(file, "inline")
}

func (a Api) SongList() revel.Result {
	actualPin, _ := strconv.Atoi(a.Session["pin"])
	if !nottify.CheckPin(actualPin) {
		return a.Redirect(App.Index)
	}

	limit, err := strconv.Atoi(a.Params.Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	artist := a.Params.Get("artist")
	title := a.Params.Get("title")
	album := a.Params.Get("album")

	nott := nottify.LoadConnection()
	songs := nott.LoadRandom(limit)

	_ = artist
	_ = title
	_ = album
	_ = songs

	return a.RenderJson(songs)
}

func (a Api) SongMeta() revel.Result {
	actualPin, _ := strconv.Atoi(a.Session["pin"])
	if !nottify.CheckPin(actualPin) {
		return a.Redirect(App.Index)
	}

	uuid := a.Params.Get("uuid")
	nott := nottify.LoadConnection()

	details := nott.GetSongMeta(uuid)
	return a.RenderJson(details)
}
