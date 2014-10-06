package nottify

import (
	"fmt"
	"os"
	"strconv"

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
	filesize int
	track    int
	comment  string
	year     int
}

func (oSong Song) GetTitle() string {
	return oSong.title
}

func (oSong Song) GetArtist() string {
	return oSong.artist
}

func (o Song) LoadDatabase(data *sqlite3.Conn, uuid string) *Song {
	db = data

	args := sqlite3.NamedArgs{"$id": uuid}
	row := make(sqlite3.RowMap)

	for s, e := db.Query("SELECT * FROM song WHERE id=$id", args); e == nil; e = s.Next() {
		e = s.Scan(row)

		o.id = fmt.Sprintf("%v", row["id"])
		o.title = fmt.Sprintf("%v", row["title"])
		o.artist = fmt.Sprintf("%v", row["artist"])
		o.album = fmt.Sprintf("%v", row["album"])
		o.length, _ = strconv.Atoi(fmt.Sprintf("%v", row["length"]))
		o.genre = fmt.Sprintf("%v", row["genre"])
		o.track, _ = strconv.Atoi(fmt.Sprintf("%v", row["track"]))
		o.year, _ = strconv.Atoi(fmt.Sprintf("%v", row["year"]))
		o.filename = fmt.Sprintf("%v", row["filename"])
		o.filesize, _ = strconv.Atoi(fmt.Sprintf("%v", row["filesize"]))
		o.comment = fmt.Sprintf("%v", row["comment"])
	}

	return &o
}

func (oSong Song) ProcessSong(file os.FileInfo) string {
	if len(oSong.title) > 0 && len(oSong.artist) > 0 {
		// @todo: Perform comparison for updating
		return ""
	}

	oSong.filesize = int(file.Size())
	tags := id3.Read(oSong.filename)

	insert_sql := "INSERT INTO song (id, title, artist, album, length, genre, track, year, filename, filesize, comment) VALUES ($id, $title, $artist, $album, $length, $genre, $track, $year, $filename, $filesize, $comment)"
	error_sql := "INSERT INTO errors (filename, found) VALUES ($filename, $found)"

	if len(tags.GetTitle()) < 1 || len(tags.GetArtist()) < 1 {
		args := sqlite3.NamedArgs{"$filename": oSong.filename, "$found": fmt.Sprintf("%s", tags)}
		err := db.Exec(error_sql, args)

		if err != nil {
			fmt.Printf("[%s][%s]\n", oSong.id, err)
		}

	} else {
		args := sqlite3.NamedArgs{"$id": oSong.id, "$title": tags.GetTitle(), "$artist": tags.GetArtist(), "$album": tags.GetAlbum(), "$length": tags.GetLength(), "$genre": tags.GetGenre(), "$track": tags.GetTrack(), "$year": tags.GetYear(), "$filename": oSong.filename, "$filesize": oSong.filesize, "$comment": tags.GetComment()}
		err := db.Exec(insert_sql, args)

		if err != nil {
			fmt.Printf("[%s][%s]\n", oSong.id, err)
		}
	}

	return ""
}

func (o Song) LoadFromResponse(r sqlite3.RowMap) *Song {
	o.id = fmt.Sprintf("%v", r["id"])
	o.title = fmt.Sprintf("%v", r["title"])
	o.artist = fmt.Sprintf("%v", r["artist"])
	o.album = fmt.Sprintf("%v", r["album"])
	o.length, _ = strconv.Atoi(fmt.Sprintf("%v", r["length"]))
	o.genre = fmt.Sprintf("%v", r["genre"])
	o.track, _ = strconv.Atoi(fmt.Sprintf("%v", r["track"]))
	o.year, _ = strconv.Atoi(fmt.Sprintf("%v", r["year"]))
	o.filename = fmt.Sprintf("%v", r["filename"])
	o.filesize, _ = strconv.Atoi(fmt.Sprintf("%v", r["filesize"]))
	o.comment = fmt.Sprintf("%v", r["comment"])

	return &o
}

func (o Song) GetMap() sqlite3.RowMap {
	r := make(sqlite3.RowMap)

	r["id"] = o.id
	r["title"] = o.title
	r["artist"] = o.artist
	r["album"] = o.album
	r["length"] = o.length
	r["genre"] = o.genre
	r["track"] = o.track
	r["year"] = o.year
	r["filename"] = o.filename
	r["filesize"] = o.filesize
	r["comment"] = o.comment

	return r
}
