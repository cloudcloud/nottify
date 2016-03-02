package song

import (
	"fmt"
	"os"

	"github.com/cloudcloud/go-id3/id3"
)

type Song struct {
	Id       string `json:"id"`
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
	fmt.Printf("%+v\n", i)

	// need to add some getters on i
	for _, v := range i.ID3V2.Items {
		fmt.Printf("%+v\n", v)
	}

	return nil
}
