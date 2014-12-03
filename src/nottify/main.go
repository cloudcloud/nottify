package nottify

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"go/build"

	"code.google.com/p/go-sqlite/go1/sqlite3"
	"github.com/cloudcloud/nottify/src/nottify/artist"
	"github.com/cloudcloud/nottify/src/nottify/song"
	"github.com/revel/revel"
)

var (
	database *sqlite3.Conn
	config   *revel.MergedConfig
)

type Nottify struct {
	songList []song.Song
}

func (n *Nottify) LoadDir(directory string) {
	filepath.Walk(directory, n.walk)

	fmt.Printf("Total songs: [%d]\n", len(n.songList))
}

func (n *Nottify) loadFile(path string, file os.FileInfo) {
	uuid := n.genUUID(file)
	_ = uuid

	/* a revisit is a must
	s := song.New()
	s.Filename = path
	s.Id = uuid

	s = s.LoadDatabase(database, uuid)
	if s == nil {
		fmt.Printf("Unable to load %s\n", uuid)
		return
	}

	err := s.ProcessSong(file)
	if err != "" {
		fmt.Printf("Failed to process %s\n", uuid)
		return
	}

	n.songList = append(n.songList, *s)
	*/
}

func (n *Nottify) walk(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	if filepath.Ext(f.Name()) == ".mp3" {
		n.loadFile(path, f)
	}

	return nil
}

func (n *Nottify) genUUID(f os.FileInfo) string {
	hasher := sha1.New()
	hasher.Write([]byte(f.Name()))

	str := hasher.Sum(nil)
	ret := fmt.Sprintf("%x", str)

	size := fmt.Sprintf("%v", f.Size())
	for i := len(size); i < 10; i++ {
		size = strings.Join([]string{size, "a"}, "")
	}

	ret = fmt.Sprintf("%s-%s", ret, size)
	return ret
}

func Build(conf *revel.MergedConfig, db *sqlite3.Conn) *Nottify {
	m := new(Nottify)

	database = db
	config = conf

	return m
}

func CheckPin(pin int) bool {
	check := revel.Config.IntDefault("nottify.pin_code", 55555)

	if check != pin {
		return false
	}

	return true
}

func LoadConnection() *Nottify {
	gopath := build.Default.GOPATH
	revel.ConfPaths = []string{path.Join(gopath, "src/github.com/cloudcloud/nottify/src/conf")}
	config, err := revel.LoadConfig("app.conf")

	if err != nil || config == nil {
		panic("Failed to Config")
	}

	dsn, db_err := config.String("nottify.sqlite_path")
	if db_err == false || dsn == "" {
		panic("No database details have been defined")
	}

	db, err := sqlite3.Open(dsn)
	if err != nil {
		panic(err.Error())
	}

	m := Build(config, db)
	return m
}

func (n *Nottify) LoadRandom(limit int) map[string]song.Song {
	result := make(map[string]song.Song)
	count := 0
	sql := "select * from song order by random() limit " + strconv.Itoa(limit) + ";"

	for s, e := database.Query(sql); e == nil; e = s.Next() {
		row := make(sqlite3.RowMap)
		s.Scan(row)

		song := song.New(database, row)
		result[strconv.Itoa(count)] = *song
		count += 1
	}

	return result
}

func (n *Nottify) LoadFromArtist(a string) *artist.Artist {
	art := artist.New(a)
	art.Load(database)

	return art
}

func (n *Nottify) GetFilename(uuid string) string {
	var filename string
	sql := "select filename from song where id=$uuid"
	args := sqlite3.NamedArgs{"$uuid": uuid}

	for s, e := database.Query(sql, args); e == nil; e = s.Next() {
		row := make(sqlite3.RowMap)
		s.Scan(row)

		filename = fmt.Sprintf("%s", row["filename"])
	}

	return filename
}

func (n *Nottify) GetSongMeta(uuid string) sqlite3.RowMap {
	song := new(Song)
	sql := "select * from song where id=$uuid"
	args := sqlite3.NamedArgs{"$uuid": uuid}

	_ = song
	for s, e := database.Query(sql, args); e == nil; e = s.Next() {
		row := make(sqlite3.RowMap)
		s.Scan(row)

		//song = song.LoadFromResponse(row)
	}

	return sqlite3.RowMap{}
}

func makeSeo(s string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, "-"))

	return s
}

func contains(e []string, c string) bool {
	for _, s := range e {
		if s == c {
			return true
		}
	}
	return false
}
