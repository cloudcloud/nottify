// Package song gives all methods and data for an individual Song.
package song

import (
	"fmt"
	"os"

	id3 "github.com/cloudcloud/go-id3"
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
	i := &id3.File{Debug: false}
	h, err := os.Open(s.Filename)
	if err != nil {
		return fmt.Errorf("File bad")
	}
	defer h.Close()

	i.Process(h)

	return nil
}

// LoadByFilename will retrieve the song from the database if it exists.
func (s *Song) LoadByFilename(filename string) (*Song, error) {
	return nil, nil
}
