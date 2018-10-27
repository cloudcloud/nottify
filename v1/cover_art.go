package v1

// Artwork allows for interaction with the art making up album covers, artist
// images, etc.
type Artwork interface {
	GetImages() []string
	GetThumbnail() string
}

// BaseArtwork gives structue to hold artwork.
type BaseArtwork struct {
	Images map[string]string `json:"images"`
}

// GetImages will give a list of all images available for the artwork.
func (a *BaseArtwork) GetImages() []string {
	b := []string{}
	for _, x := range a.Images {
		b = append(b, x)
	}

	return b
}

// GetThumbnail returns just the thumbnail image for this artwork.
func (a *BaseArtwork) GetThumbnail() string {
	i, ok := a.Images["thumbnail"]
	if ok {
		return i
	}

	// first cab off the rank
	x := ""
	for _, x = range a.Images {
		break
	}

	return x
}
