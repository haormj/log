package log

import (
	"context"
)

// Log interface
type Log interface {
	Debug(keysAndValues ...interface{})
	Debugv(msg string, keysAndValues ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Debugf(format string, a ...interface{})

	Info(keysAndValues ...interface{})
	Infov(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Infof(format string, a ...interface{})

	Warn(keysAndValues ...interface{})
	Warnv(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Warnf(format string, a ...interface{})

	Error(keysAndValues ...interface{})
	Errorv(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Errorf(format string, a ...interface{})

	With(keysAndValues ...interface{}) Log
	Withv(keysAndValues ...interface{}) Log
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
