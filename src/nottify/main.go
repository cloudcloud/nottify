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
	"github.com/revel/revel"
)

var (
	database *sqlite3.Conn
	config   *revel.MergedConfig
)

type Nottify struct {
	songList []Song
}

func (n *Nottify) LoadDir(directory string) {
	filepath.Walk(directory, n.walk)

	fmt.Printf("Total songs: [%d]\n", len(n.songList))
}

func (n *Nottify) loadFile(path string, file os.FileInfo) {
	uuid := n.genUUID(file)

	s := new(Song)
	s.filename = path
	s.id = uuid

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

func (n *Nottify) LoadRandom(limit int) map[int]sqlite3.RowMap {
	result := make(map[int]sqlite3.RowMap)
	count := 0
	sql := "select * from song order by random() limit " + strconv.Itoa(limit) + ";"

	for s, e := database.Query(sql); e == nil; e = s.Next() {
		row := make(sqlite3.RowMap)
		s.Scan(row)

		row["artist_seo"] = makeSeo(fmt.Sprintf("%s", row["artist"]))
		row["title_seo"] = makeSeo(fmt.Sprintf("%s", row["title"]))

		result[count] = row
		count += 1
	}

	return result
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

func makeSeo(s string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, "-"))

	return s
}
