package file

import (
	"io"

	"github.com/gopi-frame/writer"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/writer/file.driverName=custom`
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
