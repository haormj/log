package log

import (
	"testing"
)

func TestDebugf(t *testing.T) {
	Logger = NewLog(Level(ParseLevel("debug")), Encoder(JSON))
	Logger.Debugf("nihao %v", "haoshijie")
	Logger.Flush()
}

func TestWith(t *testing.T) {
	l := Logger.Clone()
	l.With("nihao", "world")
	hello := struct {
		Name string
		Age  int
	}{"hao", 12}
	l.With("hello", hello)
	l.Infow("hello")
}
