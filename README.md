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
	l := log.Logger.Clone()

	l.With("main", "I'm main")
	ctx := log.NewContext(context.Background(), l)
	hello(ctx)
	world(ctx)
	l.Info("main", "end")
}

func hello(ctx context.Context) {
	l, _ := log.FromContext(ctx)
	l.With("hello", "1")
	l.Infow("this is hello function")
}

func world(ctx context.Context) {
	l, _ := log.FromContext(ctx)
	l.With("world", 2)
	l.Infof("this is %s function", "world")
}
```

If pass `Log` through `Context`, pay attention to the life cycle of the `Context` to prevent memory increase
