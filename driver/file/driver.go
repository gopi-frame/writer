package file

import (
	"io"

	"github.com/gopi-frame/writer"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/writer/file.driverName=custom`
var driverName = "file"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		writer.Register(driverName, new(Driver))
	}
}

// Driver is a file writer driver
type Driver struct{}

// Open opens file writer
func (Driver) Open(options map[string]any) (io.WriteCloser, error) {
	config, err := UnmarshalOptions(options)
	if err != nil {
		return nil, err
	}
	return NewWriter(config)
}

// Open is a convenience function that calls [Driver.Open].
func Open(options map[string]any) (io.WriteCloser, error) {
	return new(Driver).Open(options)
}
