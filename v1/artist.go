package v1

// Artist provides the behaviour attached to working with an Artist.
type Artist interface {
	Albums() []Album
	Songs() []Song
}
