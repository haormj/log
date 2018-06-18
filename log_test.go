package log

import (
	"testing"
)

func TestDebugf(t *testing.T) {
	Logger = NewLog(Level(ParseLevel("debug")), Encoder(Json))
	Logger.Debugf("nihao %v", "haoshijie")
	Logger.Flush()
}
