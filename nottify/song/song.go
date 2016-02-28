package song

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cloudcloud/nottify/src/id3"

	"code.google.com/p/go-sqlite/go1/sqlite3"
)

var (
	db *sqlite3.Conn
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

func New(d *sqlite3.Conn, r sqlite3.RowMap) *Song {
	db = d
	s := new(Song)

	if r != nil {
		s.Id = fmt.Sprintf("%s", r["id"])
		s.Title = fmt.Sprintf("%v", r["title"])
		s.Artist = fmt.Sprintf("%v", r["artist"])
		s.Album = fmt.Sprintf("%v", r["album"])
		s.Length, _ = strconv.Atoi(fmt.Sprintf("%v", r["length"]))
		s.Genre = fmt.Sprintf("%v", r["genre"])
		s.Track, _ = strconv.Atoi(fmt.Sprintf("%v", r["track"]))
		s.Year, _ = strconv.Atoi(fmt.Sprintf("%v", r["year"]))
		s.Filename = fmt.Sprintf("%v", r["filename"])
		s.Filesize, _ = strconv.Atoi(fmt.Sprintf("%v", r["filesize"]))
		s.Comment = fmt.Sprintf("%v", r["comment"])
	}

	return s
}

func NewFile(f os.FileInfo, filename string) *Song {
	s := new(Song)
	s.Filename = filename
	s.Filesize = int(f.Size())

	r := s.loadTags()
	s.store(r)

	return s
}

func (s *Song) GetAlbum() string {
	return s.Album
}

func (s *Song) loadTags() bool {
	t := id3.Read(s.Filename)
	if len(t.GetTitle()) < 1 || len(t.GetArtist()) < 1 {
		return false
	}

	s.Title = t.GetTitle()
	s.Artist = t.GetArtist()
	s.Album = t.GetAlbum()
	s.Length = t.GetLength()
	s.Genre = t.GetGenre()
	s.Track = t.GetTrack()
	s.Year = t.GetYear()
	s.Comment = t.GetComment()

	return true
}

func (s *Song) store(b bool) {
	sql := ""
	args := sqlite3.NamedArgs{}

	if b {
		sql = ""
	} else {
		sql = ""
	}

	err := db.Exec(sql, args)
	if err != nil {
		fmt.Printf("[%s][%s]\n", s.Id, err)
	}
}
