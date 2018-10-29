package core

import (
	"fmt"
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
	if len(p[0]) < 1 {
		p = make([]string, len(n.Config.GetDirs()))
		copy(p, n.Config.GetDirs())
	}

	for _, x := range p {
		if n.Config.KnownDir(x) {
			go n.runDir(x, dirCh)
		} else if len(x) > 0 {
			n.Config.D(fmt.Sprintf("Unknown dir [%s]", x))
		}
	}

	for i := 0; i < len(p); i++ {
		<-dirCh
	}

	close(dirCh)
}

func (n *Nottify) runDir(d string, ch chan string) {
	//
	n.Config.D(fmt.Sprintf("Within runDir [%s]", d))

	ch <- "Complete"

	n.Config.D("After complete chan")
}
