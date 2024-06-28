package file

import (
	"io"

	"github.com/gopi-frame/writer"
)

var driverName = "file"

func init() {
	if driverName != "" {
		writer.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (d Driver) Open(options map[string]any) (io.WriteCloser, error) {
	config, err := UnmarshalOptions(options)
	if err != nil {
		return nil, err
	}
	return NewFileWriter(config)
}
