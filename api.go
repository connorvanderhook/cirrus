package cirrus

import "io"

// FileStore provides an interface for file management
// across the different file store api providers
type FileStore interface {
	// Save saves data from the io.Reader to a file named by path
	Save(path string, r io.Reader) error
	// Remove deletes the file found from the path
	Remove(path string) error
	// URL returns the URL location of a file from the path
	URL(path string) string
}
