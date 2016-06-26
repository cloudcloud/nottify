package config

import "io"

// File is the structure for all content within the configuration file itself
type File struct {
	Path []struct {
		Title string
		Dir   string
	}
	Database struct {
		Dsn string
	}
	TmpDir string
}

// Load will take content from a reader
func Load(r io.Reader) *File {
	return nil
}
