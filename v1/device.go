package v1

// Device is a behaviour associate with usage with individual devices.
type Device interface {
	NowPlaying() NowPlaying
	PlayCount() PlayCount
}

// BaseDevice provides a data structure for basic devices to use.
type BaseDevice struct {
	Name string `json:"name"`
}
