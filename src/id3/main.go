package id3

import (
	"bytes"
	"fmt"
	"os"
)

const (
	tag_size  = 128
	tag_start = 3

	title_end   = 30
	artist_end  = 60
	album_end   = 90
	year_end    = 94
	comment_end = 124
)

var (
	content []byte
)

type id3 struct {
	filename string
}

func Read(filename string) *id3 {
	i := new(id3)
	i.filename = filename

	err := readV1(filename)
	if err != "" {
		err = fmt.Sprintf("Unable to get V1 [%s]\n", err)
	}

	err = readV2(filename)
	if err != "" {
		err = fmt.Sprintf("Unable to get V2 [%s]\n", err)
	}

	return i
}

func readV1(filename string) string {
	buffer := make([]byte, tag_size)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return "No good for V1"
	}

	file.Seek(-tag_size, 2)
	file.Read(buffer)

	if getString(buffer[0:tag_start]) != "TAG" {
		return "No V1 tag found"
	}

	title := getString(buffer[tag_start:title_end])

	_ = title
	return ""
}

func readV2(filename string) string {
	//

	return ""
}

func getString(buf []byte) string {
	p := bytes.IndexByte(buf, 0)

	if p == -1 {
		p = len(buf)
	}

	return string(buf[0:p])
}
