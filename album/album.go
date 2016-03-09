// Package album provides the interactions for all Album related data. Retrieval, Storage, and Seeking.
package album

import "github.com/cloudcloud/nottify/song"

// Album is the object base upon which data is managed and actions defined.
type Album struct {
	Album  string      `json:"album"`
	Image  string      `json:"image"`
	Songs  []song.Song `json:"songs"`
	Year   int         `json:"year"`
	Artist string      `json:"artist"`
}

// New will provision a new instance of Album
func New(artist, album string) *Album {
	a := new(Album)
	a.Artist = artist
	a.Album = album

	a.Image = "http://cdn.last.fm/flatness/catalogue/noimage/2/default_album_large.png"

	return a
}

// GetImage simply returns the Image URL of the Album
func (a *Album) GetImage() string {
	return a.Image
}

// AddSong will append a Song to the Album
func (a *Album) AddSong(s *song.Song) {
	a.Songs = append(a.Songs[:], *s)
}
