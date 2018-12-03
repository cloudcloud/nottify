package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudcloud/nottify/v1/config"
)

type entry struct {
	file os.FileInfo
	path string
}

// Ingest is the gateway for executing an ingestion process
// within Nottify.
func (n *Nottify) Ingest(p []string) {
	n.m = &sync.Map{}
	n.fc = make(chan entry, 40)
	defer close(n.fc)

	if len(p[0]) < 1 {
		p = make([]string, len(n.Config.GetDirs()))
		copy(p, n.Config.GetDirs())
	}

	go n.handleFile()
	for _, x := range p {
		n.processDir(x)
	}

	n.wg.Wait()

	count := 0
	n.m.Range(func(k, v interface{}) bool {
		count++
		return true
	})

	n.Config.O(config.Info, fmt.Sprintf("Completed processing all dirs! Found %v songs!", count))
}

func (n *Nottify) checkFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if isGoodFile(info) {
		n.wg.Add(1)
		n.fc <- entry{file: info, path: path}
	}

	return nil
}

func (n *Nottify) handleFile() {
	for {
		s, more := <-n.fc

		if !more {
			return
		} else {
			n.Config.D(fmt.Sprintf("Found audio file [%s]", s.path))
			n.m.Store(s.path, s)
			n.wg.Done()
		}
	}
}

func (n *Nottify) processDir(d string) {
	if n.Config.KnownDir(d) {
		n.Config.O(config.Info, fmt.Sprintf("Processing %s", d))

		filepath.Walk(d, n.checkFile)
	} else if len(d) > 0 {
		n.Config.D(fmt.Sprintf("Unknown dir [%s]", d))
	}
}

func isGoodFile(f os.FileInfo) bool {
	if f != nil && !f.IsDir() && filepath.Ext(f.Name()) == ".mp3" {
		return true
	}

	return false
}
