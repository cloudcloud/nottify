// Package nottify provides the base interface to the main Nottify shared functionality.
package nottify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Nottify provides the shared object state.
type Nottify struct {
	//songList []song.Song
	config *Config
	db     *Db
}

// ChanMsg defines details about found files during ingestion.
type ChanMsg struct {
	file string
	info os.FileInfo
}

var (
	mainChannel chan string
	fileProc    chan ChanMsg
	songCount   int
)

// New will instantiate a Nottify for general work.
func New(c *Config) *Nottify {
	n := new(Nottify)
	n.config = c

	return n
}

// Ingest will complete a comprehensive scan of Config.Dirs
func (n *Nottify) Ingest() (string, error) {
	// channel to track crawl
	mainChannel = make(chan string)
	// channel for individual file push
	fileProc = make(chan ChanMsg, 10)

	// get the directories
	d, err := n.config.Get([]string{"dirs"})
	if err != nil {
		panic(err)
	}

	// allow for parallel file processing
	go func() {
		for {
			j, more := <-fileProc
			if !more {
				mainChannel <- fmt.Sprintf("\nFinished Processing. Total of [%v] tracks ingested.\n", songCount)
			} else {
				_ = j
				//song := song.New()

				//err = song.FromFile(j.info, j.file)
				//if err != nil {
				// pad out - these could just be already known tracks
				//} else {
				// increment result set
				//songCount++
				//}
			}
		}
	}()

	// use them individually
	dirs := strings.Split(d, ", ")
	for _, d := range dirs {
		// push each dir into a goroutine
		filepath.Walk(d, n.walk)
	}
	close(fileProc)

	return <-mainChannel, nil
}

func (n *Nottify) walk(path string, f os.FileInfo, err error) error {
	// only want specific files
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if f.IsDir() {
		return nil
	}

	if filepath.Ext(f.Name()) == ".mp3" {
		fileProc <- ChanMsg{path, f}
	}

	return nil
}
