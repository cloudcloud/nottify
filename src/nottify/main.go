package nottify

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/revel/revel"
)

var (
	database *sql.DB
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
	if !s.LoadDatabase(database, uuid) {
		fmt.Printf("Unable to load %s\n", uuid)
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

func Build(conf *revel.MergedConfig, db *sql.DB) *Nottify {
	m := new(Nottify)

	database = db
	config = conf

	return m
}