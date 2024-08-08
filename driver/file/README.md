# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/writer/driver/file.svg)](https://pkg.go.dev/github.com/gopi-frame/writer/driver/file)
[![Go](https://github.com/gopi-frame/writer/actions/workflows/driver.file.yml/badge.svg)](https://github.com/gopi-frame/writer/actions/workflows/driver.file.yml)
[![codecov](https://codecov.io/gh/gopi-frame/writer/graph/badge.svg?token=9JTZR812XD&flag=file)](https://codecov.io/gh/gopi-frame/writer?flags[0]=file)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/writer/driver/file)](https://goreportcard.com/report/github.com/gopi-frame/writer/driver/file)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package file is a file backed writer driver for [writer](https://pkg.go.dev/github.com/gopi-frame/writer) package.

# Installation
```shell
go get -u -v github.com/gopi-frame/writer/driver/file
```

# Import
```go
import _ "github.com/gopi-frame/writer/driver/file"
```

# Usage

Use `file` as driver name and valid options as below.

```go
package main

import (
    "github.com/gopi-frame/writer"
    _ "github.com/gopi-frame/writer/driver/file"
)

func main() {
    w, err := writer.Open("file", map[string]any{
        "path": "data.txt",
        "perm": 0777,
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

## Path
Type: `string`

The path to the file, if the file does not exist, it will be created.

## Perm
Type: `os.FileMode` or any types can be converted to `os.FileMode`, such as `uint32`

The permission of the file if it is created. Default is `os.ModePerm`
