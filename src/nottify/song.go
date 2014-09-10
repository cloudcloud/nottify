package nottify

import (
	"fmt"
	"os"

	"code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/cloudcloud/nottify/src/id3"
)

var (
	db *sqlite3.Conn
)

type Song struct {
	id       string
	title    string
	artist   string
	album    string
	length   int
	genre    string
	filename string
}

func (oSong Song) GetTitle() string {
	return oSong.title
}

func (oSong Song) GetArtist() string {
	return oSong.artist
}

func (oSong Song) LoadDatabase(data *sqlite3.Conn, uuid string) bool {
	db = data

	args := sqlite3.NamedArgs{"$id": uuid}
	for s, e := db.Query("SELECT filename, artist FROM song WHERE id=$id", args); e == nil; e = s.Next() {
		e = s.Scan(oSong.filename, oSong.artist)
		break
	}

	return true
}

func (oSong Song) ProcessSong(file os.FileInfo) string {
	filesize := fmt.Sprintf("%d", file.Size())

	tags := id3.Read(oSong.filename)

	_ = filesize
	_ = tags
	return ""
}
