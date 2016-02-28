package nottify

import (
	"fmt"

	"github.com/cloudcloud/nottify/nottify/song"
)

type Nottify struct {
	songList []song.Song
}

func (n *Nottify) LoadDir(directory string) {
	fmt.Printf("Total songs: [%d]\n", len(n.songList))
}
