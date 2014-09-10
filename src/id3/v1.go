package id3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	tag_size  = 128
	tag_start = 3

	title_end   = 33
	artist_end  = 63
	album_end   = 90
	year_end    = 94
	comment_end = 124
)

func (i *id3) readV1(filename string) string {
	buffer := make([]byte, tag_size)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return "Unable to open for V1"
	}

	file.Seek(-tag_size, 2)
	file.Read(buffer)

	if getString(buffer[0:tag_start]) != "TAG" {
		return "No V1 tag found"
	}

	i.title = strings.TrimSpace(getString(buffer[tag_start:title_end]))
	i.artist = getString(buffer[title_end:artist_end])
	i.album = getString(buffer[artist_end:album_end])
	i.year = getInt(buffer[album_end:year_end])
	i.comment = getString(buffer[year_end:comment_end])

	if buffer[comment_end-2] == 0 {
		i.track, _ = strconv.Atoi(fmt.Sprintf("%d", buffer[comment_end-1]))
	}

	i.genre_code, _ = strconv.Atoi(fmt.Sprintf("%d", buffer[comment_end]))

	// empty return is gooood!
	return ""
}
