# Log [![Build Status](https://travis-ci.org/haormj/log.svg?branch=master)](https://travis-ci.org/haormj/log) [![GoDoc](https://godoc.org/github.com/haormj/log?status.svg)](https://godoc.org/github.com/haormj/log) [![Go Report Card](https://goreportcard.com/badge/github.com/haormj/log)](https://goreportcard.com/report/github.com/haormj/log)

A simple go log.


## Install

```shell
go get github.com/haormj/log
```

## Usage

```go
package main

import (
	"context"

	"github.com/haormj/log"
)

func main() {
	l := log.With("username", "haormj")
	l.Info("age", 11)

	ctx := log.NewContext(context.Background(), l)
	hello(ctx)
}

func hello(ctx context.Context) {
	l, _ := log.FromContext(ctx)
	l.Info("hello", "world")
}
```

If you use `With`, and pass `Log` through `Context`, pay attention to the life cycle of the `Context` to prevent memory increase
