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
	Id       string `"id" json:"id"`
	Title    string `json:"title" "title"`
	Artist   string `json:"artist" "artist"`
	Album    string `json:"album" "album"`
	Length   int    `json:"length" "length"`
	Genre    string `json:"genre" "genre"`
	Filename string `json:"filename" "filename"`
	Filesize int    `json:"filesize" "filesize"`
	Track    int    `json:"track" "track"`
	Comment  string `json:"comment" "comment"`
	Year     int    `json:"year" "year"`
}

func (oSong Song) GetTitle() string {
	return oSong.Title
}

func (oSong Song) GetArtist() string {
	return oSong.Artist
}

func (o Song) LoadDatabase(data *sqlite3.Conn, uuid string) *Song {
	db = data

	args := sqlite3.NamedArgs{"$id": uuid}
	row := make(sqlite3.RowMap)

	for s, e := db.Query("SELECT * FROM song WHERE id=$id", args); e == nil; e = s.Next() {
		e = s.Scan(row)

		o.Id = fmt.Sprintf("%v", row["id"])
		o.Title = fmt.Sprintf("%v", row["title"])
		o.Artist = fmt.Sprintf("%v", row["artist"])
		o.Album = fmt.Sprintf("%v", row["album"])
		o.Length, _ = strconv.Atoi(fmt.Sprintf("%v", row["length"]))
		o.Genre = fmt.Sprintf("%v", row["genre"])
		o.Track, _ = strconv.Atoi(fmt.Sprintf("%v", row["track"]))
		o.Year, _ = strconv.Atoi(fmt.Sprintf("%v", row["year"]))
		o.Filename = fmt.Sprintf("%v", row["filename"])
		o.Filesize, _ = strconv.Atoi(fmt.Sprintf("%v", row["filesize"]))
		o.Comment = fmt.Sprintf("%v", row["comment"])
	}

	return &o
}

func (oSong Song) ProcessSong(file os.FileInfo) string {
	if len(oSong.Title) > 0 && len(oSong.Artist) > 0 {
		// @todo: Perform comparison for updating
		return ""
	}

	oSong.Filesize = int(file.Size())
	tags := id3.Read(oSong.Filename)

	insert_sql := "INSERT INTO song (id, title, artist, album, length, genre, track, year, filename, filesize, comment) VALUES ($id, $title, $artist, $album, $length, $genre, $track, $year, $filename, $filesize, $comment)"
	error_sql := "INSERT INTO errors (filename, found) VALUES ($filename, $found)"

	if len(tags.GetTitle()) < 1 || len(tags.GetArtist()) < 1 {
		args := sqlite3.NamedArgs{"$filename": oSong.Filename, "$found": fmt.Sprintf("%s", tags)}
		err := db.Exec(error_sql, args)

		if err != nil {
			fmt.Printf("[%s][%s]\n", oSong.Id, err)
		}

	} else {
		args := sqlite3.NamedArgs{"$id": oSong.Id, "$title": tags.GetTitle(), "$artist": tags.GetArtist(), "$album": tags.GetAlbum(), "$length": tags.GetLength(), "$genre": tags.GetGenre(), "$track": tags.GetTrack(), "$year": tags.GetYear(), "$filename": oSong.Filename, "$filesize": oSong.Filesize, "$comment": tags.GetComment()}
		err := db.Exec(insert_sql, args)

		if err != nil {
			fmt.Printf("[%s][%s]\n", oSong.Id, err)
		}
	}

	return ""
}

func (o Song) LoadFromResponse(r sqlite3.RowMap) *Song {
	o.Id = fmt.Sprintf("%v", r["id"])
	o.Title = fmt.Sprintf("%v", r["title"])
	o.Artist = fmt.Sprintf("%v", r["artist"])
	o.Album = fmt.Sprintf("%v", r["album"])
	o.Length, _ = strconv.Atoi(fmt.Sprintf("%v", r["length"]))
	o.Genre = fmt.Sprintf("%v", r["genre"])
	o.Track, _ = strconv.Atoi(fmt.Sprintf("%v", r["track"]))
	o.Year, _ = strconv.Atoi(fmt.Sprintf("%v", r["year"]))
	o.Filename = fmt.Sprintf("%v", r["filename"])
	o.Filesize, _ = strconv.Atoi(fmt.Sprintf("%v", r["filesize"]))
	o.Comment = fmt.Sprintf("%v", r["comment"])

	return &o
}

func (o Song) GetMap() sqlite3.RowMap {
	r := make(sqlite3.RowMap)

	r["id"] = o.Id
	r["title"] = o.Title
	r["artist"] = o.Artist
	r["album"] = o.Album
	r["length"] = o.Length
	r["genre"] = o.Genre
	r["track"] = o.Track
	r["year"] = o.Year
	r["filename"] = o.Filename
	r["filesize"] = o.Filesize
	r["comment"] = o.Comment

	return r
}
