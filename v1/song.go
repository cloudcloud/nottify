package v1

// Song is a behaviour to work with individual Songs.
type Song interface {
}

// BaseSong is a structure to contain data for a Song.
type BaseSong struct {
	FilePath string `json:"file_path"`
}

// FromFile will take a specific filename and turn it into a Song.
func FromFile(f string) Song {
	//

	return nil
}
