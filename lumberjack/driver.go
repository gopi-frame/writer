package lumberjack

import (
	"io"

	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/writer"
	"gopkg.in/natefinch/lumberjack.v2"
)

// This variable can be replaced through `go build -ldflags=-X github.com/gopi-frame/writer/lumberjack.driverName=custom`
var driverName = "lumberjack"

//goland:noinspection GoBoolExpressions
func init() {
	if driverName != "" {
		writer.Register(driverName, new(Driver))
	}
}

type Driver struct{}

func (d Driver) Open(options map[string]any) (io.WriteCloser, error) {
	logger := new(lumberjack.Logger)
	err := mapstructure.WeakDecode(options, logger)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
