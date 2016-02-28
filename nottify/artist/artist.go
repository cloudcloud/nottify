package artist

import (
	"github.com/cloudcloud/nottify/nottify/album"
	"github.com/cloudcloud/nottify/nottify/song"
)

type Artist struct {
	Artist string        `json:"artist"`
	Albums []album.Album `json:"albums"`
}

func New(artist string) *Artist {
	a := new(Artist)
	a.Artist = artist

	return a
}

func (a *Artist) Load() {
	sql := "select * from song where artist=? order by filename asc"
	_ = sql
}

func (a *Artist) addSong(s *song.Song) {
	i := false

	for art := range a.Albums {
		if s.Album == a.Albums[art].Album {
			a.Albums[art].AddSong(s)
			i = true
		}
	}

	if !i {
		alb := album.New(a.Artist, s.Album)
		alb.AddSong(s)
		a.Albums = append(a.Albums[:], *alb)
	}
}
