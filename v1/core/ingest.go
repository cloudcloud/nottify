package core

import (
	"os"
)

type entry struct {
	file *os.File
}

var (
	count  = 0
	dirCh  = make(chan string)
	fileCh = make(chan entry, 10)
	mainCh = make(chan string)
)

// Use a channel to manage each directory
// Within each directory, walk the file system
// For each file, push into the file chan

func (n *Nottify) Ingest(p []string) {
	dirCount := 0
	if len(p) > 0 {
		for _, x := range p {
			// check if x is in n.Config.GetDirs()
			if n.Config.KnownDir(x) {
				n.runDir(x)
				dirCount++
			}
		}
	} else {
		for _, x := range n.Config.GetDirs() {
			n.runDir(x)
			dirCount++
		}
	}

	//
	for i := 0; i < dirCount; i++ {
		<-dirCh
	}
}

func (n *Nottify) runDir(d string) {
}
