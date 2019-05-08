package cirrus

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// TODO: make configurable
// 777 – Everyone can read write and execute.
const localPerm = 0777

// Local is a local file storage.
type Local struct {
	dir, url string
}

// NewLocal returns a new local file storage.
func NewLocal(dir, url string) (*Local, error) {
	if err := os.MkdirAll(dir, localPerm); err != nil {
		return nil, fmt.Errorf("failed to create a directory (%q): %v", dir, err)
	}

	s := &Local{
		dir: dir,
		url: strings.TrimSuffix(url, "/"),
	}
	return s, nil
}

// Save saves data from r to file with the given path.
func (s *Local) Save(path string, r io.Reader) error {
	fullpath := filepath.Join(s.dir, path)

	if err := os.MkdirAll(filepath.Dir(fullpath), localPerm); err != nil {
		return fmt.Errorf("failed to create dir for file (%s): %v", fullpath, err)
	}

	w, err := os.Create(fullpath)
	if err != nil {
		return fmt.Errorf("failed to create a file (%s): %v", fullpath, err)
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to copy data to file (%s): %v", fullpath, err)
	}

	w.Close()
	return nil
}

// URL returns an URL of the file with the given path.
func (s *Local) URL(path string) string {
	return s.url + "/" + path
}
