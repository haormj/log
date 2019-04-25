package log

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Uber log
type Uber struct {
	sugar *zap.SugaredLogger
	kvs   []interface{}
	rw    sync.RWMutex
}

// NewLog uber log
func NewLog(opts ...Option) Log {

	options := newOptions(opts...)

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     options.MessageKey,
		LevelKey:       options.LevelKey,
		TimeKey:        options.TimeKey,
		NameKey:        options.NameKey,
		CallerKey:      options.CallerKey,
		StacktraceKey:  options.StacktraceKey,
		LineEnding:     options.LineEnding,
		EncodeLevel:    options.EncodeLevel,
		EncodeTime:     options.EncodeTime,
		EncodeDuration: options.EncodeDuration,
		EncodeCaller:   options.EncodeCaller,
		EncodeName:     options.EncodeName,
	}

	var encoder zapcore.Encoder
	switch options.Encoder {
	case Console:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case JSON:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var writeSyncer zapcore.WriteSyncer
	if len(options.Filename) == 0 {
		writeSyncer = zapcore.NewMultiWriteSyncer(os.Stdout)
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjack.Logger{
			Filename:   options.Filename,
			MaxSize:    options.MaxSize,
			MaxAge:     options.MaxAge,
			MaxBackups: options.MaxBackups,
			LocalTime:  options.LocalTime,
			Compress:   options.Compress,
		}))
	}
	var l zapcore.Level
	l.Set(string(options.Level))
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(writeSyncer)), l)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &Uber{
		sugar: logger.Sugar(),
		kvs:   make([]interface{}, 0),
		rw:    sync.RWMutex{},
	}
}

func (u *Uber) get() []interface{} {
	u.rw.RLock()
	kvs := make([]interface{}, len(u.kvs))
	copy(kvs, u.kvs)
	u.rw.RUnlock()
	return kvs
}

func (u *Uber) format(format string, a []interface{}) string {
	// Format with Sprint, Sprintf, or neither.
	msg := format
	if msg == "" && len(a) > 0 {
		msg = fmt.Sprint(a...)
	} else if msg != "" && len(a) > 0 {
		msg = fmt.Sprintf(format, a...)
	}
	return msg
}

// Debug key value
func (u *Uber) Debug(keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Debugw("", kvs...)
}

// Debugw with message
func (u *Uber) Debugw(msg string, keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Debugw(msg, kvs...)
}

// Debugf formats according to a format specifier
func (u *Uber) Debugf(format string, a ...interface{}) {
	kvs := u.get()
	msg := u.format(format, a)

	u.sugar.Debugw(msg, kvs...)
}

// Info key value
func (u *Uber) Info(keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Infow("", kvs...)
}

// Infow with message
func (u *Uber) Infow(msg string, keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Infow(msg, kvs...)
}

// Infof formats according to a format specifier
func (u *Uber) Infof(format string, a ...interface{}) {
	kvs := u.get()
	msg := u.format(format, a)

	u.sugar.Infow(msg, kvs...)
}

// Warn key value
func (u *Uber) Warn(keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Warnw("", kvs...)
}

// Warnw with message
func (u *Uber) Warnw(msg string, keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Warnw(msg, kvs...)
}

// Warnf formats according to a format specifier
func (u *Uber) Warnf(format string, a ...interface{}) {
	kvs := u.get()
	msg := u.format(format, a)

	u.sugar.Warnw(msg, kvs...)
}

// Error key value
func (u *Uber) Error(keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Errorw("", kvs...)
}

// Errorw with message
func (u *Uber) Errorw(msg string, keysAndValues ...interface{}) {
	kvs := u.get()
	kvs = append(kvs, keysAndValues...)

	u.sugar.Errorw(msg, kvs...)
}

// Errorf formats according to a format specifier
func (u *Uber) Errorf(format string, a ...interface{}) {
	kvs := u.get()
	msg := u.format(format, a)

	u.sugar.Errorw(msg, kvs...)
}

// With key value
func (u *Uber) With(keysAndValues ...interface{}) {
	u.rw.Lock()
	u.kvs = append(u.kvs, keysAndValues...)
	u.rw.Unlock()
}

func (u *Uber) Clone() Log {
	kvs := u.get()

	return &Uber{
		sugar: u.sugar.With(),
		kvs:   kvs,
		rw:    sync.RWMutex{},
	}
}

// Flush buffered log
func (u *Uber) Flush() error {
	return u.sugar.Sync()
}
