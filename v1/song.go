package v1

import (
	"os"
	"strings"

	id3 "github.com/cloudcloud/go-id3"
)

// Song is a behaviour to work with individual Songs.
type Song interface {
	GetArtist() string
	GetTitle() string
}

// BaseSong is a structure to contain data for a Song.
type BaseSong struct {
	FilePath string `json:"file_path"`

	Artist string
	Album  string
	Title  string
}

// GetArtist .
func (b *BaseSong) GetArtist() string {
	return b.Artist
}

// GetTitle .
func (b *BaseSong) GetTitle() string {
	return b.Title
}

// FromFile will take a specific filename and turn it into a Song.
func FromFile(f string) (*BaseSong, error) {
	s := &BaseSong{
		FilePath: f,
	}

	if err := s.parseFile(); err != nil {
		// something something
		return nil, err
	}

	return s, nil
}

func (s *BaseSong) parseFile() error {
	f, err := os.Open(s.FilePath)
	if err != nil {
		return err
	}

	i := &id3.File{Filename: s.FilePath}
	i.Process(f)

	t := " \t\n\r\u0000\x00"

	s.Artist = strings.Trim(i.GetArtist(), t)
	s.Album = strings.Trim(i.GetAlbum(), t)
	s.Title = strings.Trim(i.GetTitle(), t)

	return nil
}
