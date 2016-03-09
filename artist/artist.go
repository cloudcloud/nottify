// Package artist will give interactions for Artist related queries.
package artist

import (
	"github.com/cloudcloud/nottify/album"
	"github.com/cloudcloud/nottify/song"
)

// Artist is the base of the object for the data and methods upon an Artist
type Artist struct {
	Artist string        `json:"artist"`
	Albums []album.Album `json:"albums"`
}

// New will provision a new instance of the Artist object
func New(artist string) *Artist {
	a := new(Artist)
	a.Artist = artist

	return a
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
