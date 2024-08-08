package writer

import (
	"bytes"
	"io"
)

func init() {
	Register("test", new(mockDriver))
}

type mockDriver struct{}

func (mockDriver) Open(_ map[string]any) (io.WriteCloser, error) {
	return &mockWriter{buffer: bytes.NewBuffer(nil)}, nil
}
