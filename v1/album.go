package v1

// Album provides the behaviour attached to working with an Album.
type Album interface {
	GetArtists() []Artist
	GetReleaseYear() int
	GetSongs() []Song
}

// BaseAlbum provides a structure to enable Album efforts.
type BaseAlbum struct {
	Artists     []Artist `json:"albums"`
	ReleaseYear int      `json:"release_year"`
	Songs       []Song   `json:"songs"`
}

// GetArtists will give the Artist list associated with this album.
func (a *BaseAlbum) GetArtists() []Artist {
	return a.Artists
}

// GetReleaseYear provides the year the Album was released.
func (a *BaseAlbum) GetReleaseYear() int {
	return a.ReleaseYear
}

// GetSongs will provide the list of Songs appearing on this Album.
func (a *BaseAlbum) GetSongs() []Song {
	return a.Songs
}
