package id3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	header_length = 10
)

var (
	header []byte
)

type Frame struct {
	name    string
	data    string
	size    int
	initial byte
}

func (i *id3) readV2(filename string) string {
	header = make([]byte, header_length)

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return "Unable to open V2"
	}

	file.Read(header)
	if getString(header[:3]) != "ID3" {
		return "No ID3v2 found"
	}

	// get frames, process
	for c := 0; c < 20; c++ {
		frame := getFrame(file)
		if frame.initial != byte('T') {
			continue
		}

		if frame.name == "" {
			break
		}

		trimmed := strings.TrimSpace(frame.data)
		if frame.name == "TIT2" {
			i.title = trimmed
		} else if frame.name == "TPE1" {
			i.artist = trimmed
		} else if frame.name == "TCOM" {
			if i.artist == "" {
				i.artist = trimmed
			}
		} else if frame.name == "TPE2" {
			if i.artist == "" {
				i.artist = trimmed
			}
		} else if frame.name == "TOPE" {
			if i.artist == "" {
				i.artist = trimmed
			}
		} else if frame.name == "TALB" {
			i.album = trimmed
		} else if frame.name == "TRCK" {
			i.track, _ = strconv.Atoi(trimmed)
		} else if frame.name == "TIT1" {
			i.genre = trimmed
		} else if frame.name == "TCON" {
			// this is sort of genre detail
			if i.genre == "" {
				i.genre = trimmed
			}
		} else if frame.name == "TYER" {
			i.year, _ = strconv.Atoi(trimmed)
		} else if frame.name == "TDAT" {
			if i.year < 1 {
				i.year, _ = strconv.Atoi(trimmed)
			}
		} else if len(frame.name) != 4 || isSkippable(frame.name) {
			continue
		} else {
			fmt.Printf("Found [%s]:[%s]\n", frame.name, trimmed)
		}
	}

	return ""
}

func getFrame(f *os.File) *Frame {
	frame := new(Frame)

	tag := make([]byte, 4)
	size := make([]byte, 4)
	padding := make([]byte, 2)

	f.Read(tag)
	frame.initial = tag[0]
	frame.name = getString(tag)

	f.Read(size)
	frame.size = int(size[0]<<21 | size[1]<<14 | size[2]<<7 | size[3])

	frame_content := make([]byte, frame.size)

	f.Read(padding)
	f.Read(frame_content)
	frame.data = getString(frame_content)

	return frame
}

func isSkippable(name string) bool {
	skip := []string{"TENC", "TLAN", "TXXX", "TPUB", "TSSE", "TLEN", "TCOP", "TPOS", "TPE4", "TMED", "TDRC", "TBPM", "TDEN", "TSOA", "TSOP", "TSOT", "TDTG", "TSRC", "TORY", "Tune", "TunN", "TtP=", "TunP", "To I", "TCMP", "TSIZ", "The ", "TFLT", "Tool", "TPE3", "Team", "TEQ{", "Tour"}

	for i := 0; i < len(skip); i++ {
		if name == skip[i] {
			return true
		}
	}

	return false
}
