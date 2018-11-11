package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type entry struct {
	file os.FileInfo
	path string
}

var (
	count  = 0
	fileCh = make(chan entry, 10)
)

// Ingest is the gateway for executing an ingestion process
// within Nottify.
func (n *Nottify) Ingest(p []string) {
	dirCh := make(chan string)
	defer close(dirCh)

	if len(p[0]) < 1 {
		p = make([]string, len(n.Config.GetDirs()))
		copy(p, n.Config.GetDirs())
	}

	for _, x := range p {
		if n.Config.KnownDir(x) {
			go filepath.Walk(x, n.checkFile)
		} else if len(x) > 0 {
			n.Config.D(fmt.Sprintf("Unknown dir [%s]", x))
		}
	}

	go n.handleFile(fileCh, dirCh)

	for i := 0; i < len(p); i++ {
		<-dirCh
	}

	n.Config.D("Completed processing all dirs!")

	close(dirCh)
	close(fileCh)
}

func (n *Nottify) handleFile(ch chan entry, cl chan string) {
	for {
		s, more := <-ch
		if !more {
			n.Config.D("Completed files")

			cl <- "Completed"
		} else {
			n.Config.D(fmt.Sprintf("Found audio file [%s]", s.path))
		}
	}
}

func (n *Nottify) checkFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if isGoodFile(info) {
		fileCh <- entry{file: info, path: path}
	}

	return err
}

func isGoodFile(f os.FileInfo) bool {
	if strings.HasSuffix(f.Name(), ".mp3") && !f.IsDir() {
		return true
	}

	return false
}
