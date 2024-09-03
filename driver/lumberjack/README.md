# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/writer/driver/lumberjack.svg)](https://pkg.go.dev/github.com/gopi-frame/writer/driver/lumberjack)
[![Test driver lumberjack](https://github.com/gopi-frame/writer/actions/workflows/driver.lumberjack.yml/badge.svg)](https://github.com/gopi-frame/writer/actions/workflows/driver.lumberjack.yml)
[![codecov](https://codecov.io/gh/gopi-frame/writer/graph/badge.svg?token=9JTZR812XD&flag=lumberjack)](https://codecov.io/gh/gopi-frame/writer?flags[0]=lumberjack)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/writer/driver/lumberjack)](https://goreportcard.com/report/github.com/gopi-frame/writer/driver/lumberjack)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package `lumberjack` is a rotating file writer driver for [writer](https://pkg.go.dev/github.com/gopi-frame/writer)
package.

# Installation
```shell
go get -u -v github.com/gopi-frame/writer/driver/lumberjack
```

# Import
```go
import _ "github.com/gopi-frame/writer/driver/lumberjack"
```

# Usage

Use `lumberjack` as driver name and valid options as below.

```go
package main

import (
	"github.com/gopi-frame/writer"
    _ "github.com/gopi-frame/writer/driver/lumberjack"
)


func main() {
    w, err := writer.Open("lumberjack", map[string]any{
        "filename": "data.txt",
        "maxSize": 100,
        "maxAge": 7,
        "maxBackups": 10,
        "compress": true,
    })
    if err != nil {
        panic(err)
    }
    if _, err := w.Write([]byte("Hello World")); err != nil {
        panic(err)
    }
}
```

# Options

This package uses [mapstructure](https://pkg.go.dev/github.com/go-viper/mapstructure/v2) to decode the options.

The following options are available:

## Filename
Type: `string`

Filename is the file to write logs to.
For more information about the `Filename`,
see [lumberjack.Logger.Filename](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.Filename).

## MaxSize
Type: `int`

MaxSize is the maximum size in megabytes of the log file before it gets rotated.
For more information about the `MaxSize`,
see [lumberjack.Logger.MaxSize](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.MaxSize).

## MaxAge
Type: `int`

MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename.
For more information about the `MaxAge`,
see [lumberjack.Logger.MaxAge](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.MaxAge).

## MaxBackups
Type: `int`

MaxBackups is the maximum number of old log files to retain.
For more information about the `MaxBackups`,
see [lumberjack.Logger.MaxBackups](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.MaxBackups).

## LocalTime
Type: `bool`

LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time.
For more information about the `LocalTime`,
see [lumberjack.Logger.LocalTime](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.LocalTime).

## Compress
Type: `bool`

Compress determines if the rotated log files should be compressed using gzip.
For more information about the `Compress`,
see [lumberjack.Logger.Compress](https://pkg.go.dev/gopkg.in/natefinch/lumberjack.v2#Logger.Compress).
