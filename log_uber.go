package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Uber log
type Uber struct {
	sugar *zap.SugaredLogger
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
	case Json:
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
	}
}

func (u *Uber) logv(args []interface{}) []interface{} {
	kvs := make([]interface{}, 0)
	for _, v := range args {
		kvs = append(kvs, fmt.Sprintf("%+v", v))
	}
	return kvs
}

// Debug key value
func (u *Uber) Debug(keysAndValues ...interface{}) {
	u.sugar.Debugw("", keysAndValues...)
}

// Debugv use %+v
func (u *Uber) Debugv(msg string, keysAndValues ...interface{}) {
	u.sugar.Debugw(msg, u.logv(keysAndValues)...)
}

// Debugw with message
func (u *Uber) Debugw(msg string, keysAndValues ...interface{}) {
	u.sugar.Debugw(msg, keysAndValues...)
}

// Debugf formats according to a format specifier
func (u *Uber) Debugf(format string, a ...interface{}) {
	u.sugar.Debugf(format, a...)
}

// Info key value
func (u *Uber) Info(keysAndValues ...interface{}) {
	u.sugar.Infow("", keysAndValues...)
}

// Infov use %+v
func (u *Uber) Infov(msg string, keysAndValues ...interface{}) {
	u.sugar.Infow(msg, u.logv(keysAndValues)...)
}

// Infow with message
func (u *Uber) Infow(msg string, keysAndValues ...interface{}) {
	u.sugar.Infow(msg, keysAndValues...)
}

// Infof formats according to a format specifier
func (u *Uber) Infof(format string, a ...interface{}) {
	u.sugar.Infof(format, a...)
}

// Warn key value
func (u *Uber) Warn(keysAndValues ...interface{}) {
	u.sugar.Warnw("", keysAndValues...)
}

// Warnv use %+v
func (u *Uber) Warnv(msg string, keysAndValues ...interface{}) {
	u.sugar.Warnw(msg, u.logv(keysAndValues)...)
}

// Warnw with message
func (u *Uber) Warnw(msg string, keysAndValues ...interface{}) {
	u.sugar.Warnw(msg, keysAndValues...)
}

// Warnf formats according to a format specifier
func (u *Uber) Warnf(format string, a ...interface{}) {
	u.sugar.Warnf(format, a...)
}

// Error key value
func (u *Uber) Error(keysAndValues ...interface{}) {
	u.sugar.Errorw("", keysAndValues...)
}

// Errorv use %+v
func (u *Uber) Errorv(msg string, keysAndValues ...interface{}) {
	u.sugar.Errorw(msg, u.logv(keysAndValues)...)
}

// Errorw with message
func (u *Uber) Errorw(msg string, keysAndValues ...interface{}) {
	u.sugar.Errorw(msg, keysAndValues...)
}

// Errorf formats according to a format specifier
func (u *Uber) Errorf(format string, a ...interface{}) {
	u.sugar.Errorf(format, a...)
}

// With key value
func (u *Uber) With(keysAndValues ...interface{}) Log {
	return &Uber{
		sugar: u.sugar.With(keysAndValues...),
	}
}

// Withv use %+v
func (u *Uber) Withv(keysAndValues ...interface{}) Log {
	return &Uber{
		sugar: u.sugar.With(u.logv(keysAndValues)...),
	}
}

// Flush buffered log
func (u *Uber) Flush() error {
	return u.sugar.Sync()
}
