package album

import "github.com/cloudcloud/nottify/nottify/song"

type Album struct {
	Album  string      `json:"album"`
	Image  string      `json:"image"`
	Songs  []song.Song `json:"songs"`
	Year   int         `json:"year"`
	Artist string      `json:"artist"`
}

func New(artist, album string) *Album {
	a := new(Album)
	a.Artist = artist
	a.Album = album

	a.Image = "http://cdn.last.fm/flatness/catalogue/noimage/2/default_album_large.png"

	return a
}

func (a *Album) GetImage() string {
	return a.Image
}

func (a *Album) AddSong(s *song.Song) {
	a.Songs = append(a.Songs[:], *s)
}
