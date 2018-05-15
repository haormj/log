package log

import (
	"context"
)

// Log interface
type Log interface {
	Debug(keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	With(keysAndValues ...interface{}) Log
	Flush() error
}

type logKey struct{}

var (
	// DefaultLog default use uber log
	Logger = NewLog()
)

// NewContext put Log to context
func NewContext(ctx context.Context, l Log) context.Context {
	return context.WithValue(ctx, logKey{}, l)
}

// FromContext get Log from context
func FromContext(ctx context.Context) (Log, bool) {
	log, ok := ctx.Value(logKey{}).(Log)
	return log, ok
}
