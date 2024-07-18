package lumberjack

import (
	"io"

	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/writer"
	"gopkg.in/natefinch/lumberjack.v2"
)

var driverName = "lumberjack"

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
