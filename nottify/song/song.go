package song

import (
	"os"
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

func NewFile(f os.FileInfo, filename string) *Song {
	s := new(Song)
	s.Filename = filename
	s.Filesize = int(f.Size())

	return s
}
