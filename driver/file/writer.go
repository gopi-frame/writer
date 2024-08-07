package file

import (
	"os"
	"path/filepath"

	"github.com/gopi-frame/exception"
)

// Writer is a file writer
type Writer struct {
	*Config
	*os.File
}

// NewWriter creates a new file writer
func NewWriter(config *Config) (*Writer, error) {
	w := &Writer{
		Config: config,
	}
	err := w.Open()
	if err != nil {
		return nil, err
	}
	return w, nil
}

// Open opens file writer
func (w *Writer) Open() error {
	if w.Config.File == "" {
		return exception.New("file can't be empty")
	}
	if w.Perm == 0 {
		w.Perm = os.ModePerm
	}
	err := os.MkdirAll(filepath.Dir(w.Config.File), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(w.Config.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, w.Perm)
	if err != nil {
		return err
	}
	w.File = f
	return nil
}
