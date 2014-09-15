package nottify

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
