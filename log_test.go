package log

import (
	"testing"
)

func TestDebugf(t *testing.T) {
	Logger = NewLog(Level(ParseLevel("debug")), Encoder(Json))
	Logger.Debugf("nihao %v", "haoshijie")
	Logger.Flush()
}

func TestInfov(t *testing.T) {
	Logger = NewLog(Encoder(Json))
	hello := struct {
		Name string
		Age  int
	}{"hao", 12}
	Logger.Infov("hello world", "hello", hello)
}
