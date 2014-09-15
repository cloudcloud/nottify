package id3

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	version_length = 2
	flag_length    = 1
	size_length    = 4
	frame_id       = 4
	frame_size     = 4
	frame_flags    = 2
)

var (
	content []byte
)

type id3 struct {
	filename string

	title      string
	artist     string
	album      string
	year       int
	track      int
	comment    string
	genre_code int
	genre      string
	length     int
}

func Read(filename string) *id3 {
	i := new(id3)
	i.filename = filename

	err := i.readV2(filename)
	if err != "" {
		err = fmt.Sprintf("Unable to get V2 [%s]\n", err)
	}

	if len(i.title) < 1 || len(i.artist) < 1 {
		err = i.readV1(filename)
		if err != "" {
			err = fmt.Sprintf("Unable to get V1 [%s]\n", err)
		}
	}

	return i
}

func (i *id3) GetTitle() string {
	return i.title
}

func (i *id3) GetArtist() string {
	return i.artist
}

func (i *id3) GetAlbum() string {
	return i.album
}

func (i *id3) GetLength() int {
	return i.length
}

func (i *id3) GetGenre() string {
	return i.genre
}

func (i *id3) GetTrack() int {
	return i.track
}

func (i *id3) GetYear() int {
	return i.year
}

func (i *id3) GetComment() string {
	return i.comment
}

func (i *id3) Print() {
	fmt.Printf("[%s] - [%s]\n", i.artist, i.title)
}

func getString(buf []byte) string {
	p := bytes.IndexByte(buf, 0)

	if p == -1 {
		p = len(buf)
	}

	return strings.TrimSpace(cleanUTF8(string(buf[0:p])))
}

func getInt(buf []byte) int {
	p := bytes.IndexByte(buf, 0)

	if p == -1 {
		p = len(buf)
	}

	response, _ := strconv.Atoi(cleanUTF8(string(buf[0:p])))
	return response
}

func cleanUTF8(s string) string {
	if !utf8.ValidString(s) {
		v := make([]rune, 0, len(s))
		for i, r := range s {
			if r == utf8.RuneError {
				_, size := utf8.DecodeRuneInString(s[i:])
				if size == 1 {
					continue
				}
			}
			v = append(v, r)
		}
		s = string(v)
	}

	return s
}
