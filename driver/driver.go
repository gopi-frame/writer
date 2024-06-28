package driver

import "io"

// Driver driver
type Driver interface {
	Open(options map[string]any) (io.WriteCloser, error)
}
