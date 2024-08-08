# Overview
[![Go Reference](https://pkg.go.dev/badge/github.com/gopi-frame/writer.svg)](https://pkg.go.dev/github.com/gopi-frame/writer)
[![Go](https://github.com/gopi-frame/writer/actions/workflows/go.yml/badge.svg)](https://github.com/gopi-frame/writer/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gopi-frame/writer/graph/badge.svg?token=9JTZR812XD&flag=writer)](https://codecov.io/gh/gopi-frame/writer?flags[0]=writer)
[![Go Report Card](https://goreportcard.com/badge/github.com/gopi-frame/writer)](https://goreportcard.com/report/github.com/gopi-frame/writer)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Package writer provides a generic writer API.

# Installation
```shell
go get -u -v github.com/gopi-frame/writer
```

# Import
```go
import "github.com/gopi-frame/writer"
```

# Available drivers
- [file](./driver/file/README.md)
- [lumberjack](./driver/lumberjack/README.md)

# How to create a custom driver
Just implement [`writer.Driver`](https://pkg.go.dev/github.com/gopi-frame/contract/writer#Driver) interface
and register it using [`writer.Register`](https://pkg.go.dev/github.com/gopi-frame/writer#Register) function.

Example:

```go
package main

import "github.com/gopi-frame/writer"

func main() {
    writer.Register("custom", new(CustomDriver))
    w, err := writer.Open("custom", map[string]any{
        // options here
    })
}
```