// Package song gives all methods and data for an individual Song.
package song

import (
	"fmt"
	"os"

	"github.com/cloudcloud/go-id3/id3"
)

// Song is the data container and method provider for working with a Song.
type Song struct {
	ID       string `json:"id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Album    string `json:"album"`
	Length   int    `json:"length"`
	Genre    string `json:"genre"`
	Filename string `json:"filename"`
	Filesize int    `json:"filesize"`
	Track    int    `json:"track"`
	Comment  string `json:"comment"`
	Year     int    `json:"year"`
}

// New will provision a fresh instance of Song.
func New() *Song {
	s := new(Song)

	return s
}

// FromFile takes some FileInfo and parses for Nottify.
func (s *Song) FromFile(f os.FileInfo, filename string) error {
	s.Filename = filename
	s.Filesize = int(f.Size())

	// parse the id3 and push to database
	i, err := id3.New(filename)
	if err != nil {
		return err
	}

	i = i.Process()
	fmt.Printf("%s - %s [%s] (%v)\n", i.GetArtist(), i.GetTitle(), i.GetAlbum(), i.GetTrackNumber())

	s.Artist = i.GetArtist()
	s.Title = i.GetTitle()
	s.Album = i.GetAlbum()
	s.Length = i.GetLength()
	s.Genre = i.GetGenre()
	s.Track = i.GetTrackNumber()
	s.Comment = i.GetComment()
	s.Year = i.GetReleaseYear()

	return nil
}

// LoadByFilename will retrieve the song from the database if it exists.
func (s *Song) LoadByFilename(filename string) (*Song, error) {
	return nil, nil
}
