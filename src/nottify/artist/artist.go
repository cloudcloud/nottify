package artist

import (
	"code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/cloudcloud/nottify/src/nottify/album"
	"github.com/cloudcloud/nottify/src/nottify/song"
)

var (
	db *sqlite3.Conn
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

func (a *Artist) Load(d *sqlite3.Conn) {
	db = d
	sql := "select * from song where artist=? order by filename asc"

	for s, e := db.Query(sql, a.Artist); e == nil; e = s.Next() {
		row := make(sqlite3.RowMap)
		s.Scan(row)

		song := song.New(db, row)

		a.addSong(song)
	}
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
