package nottify

import "database/sql"

var (
	db *sql.DB
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

func (oSong Song) SetID(uuid string) {
	oSong.id = uuid
}

func (oSong Song) LoadDatabase(data *sql.DB, uuid string) bool {
	db = data

	rows, err := db.Query("SELECT filename, artist FROM song WHERE id=?", uuid)
	if err != nil {
		return false
	}

	for rows.Next() {
		err = rows.Scan(oSong.filename, oSong.artist)
		break
	}

	return true
}
