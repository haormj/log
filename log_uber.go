package log

import (
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
	encoder := zapcore.NewJSONEncoder(encoderConfig)

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

// Debug key value
func (u *Uber) Debug(keysAndValues ...interface{}) {
	u.sugar.Debugw("", keysAndValues...)
}

// Debugw with message
func (u *Uber) Debugw(msg string, keysAndValues ...interface{}) {
	u.sugar.Debugw(msg, keysAndValues...)
}

// Info key value
func (u *Uber) Info(keysAndValues ...interface{}) {
	u.sugar.Infow("", keysAndValues...)
}

// Infow with message
func (u *Uber) Infow(msg string, keysAndValues ...interface{}) {
	u.sugar.Infow(msg, keysAndValues...)
}

// Warn key value
func (u *Uber) Warn(keysAndValues ...interface{}) {
	u.sugar.Warnw("", keysAndValues...)
}

// Warnw with message
func (u *Uber) Warnw(msg string, keysAndValues ...interface{}) {
	u.sugar.Warnw(msg, keysAndValues...)
}

// Error key value
func (u *Uber) Error(keysAndValues ...interface{}) {
	u.sugar.Errorw("", keysAndValues...)
}

// Errorw with message
func (u *Uber) Errorw(msg string, keysAndValues ...interface{}) {
	u.sugar.Errorw(msg, keysAndValues...)
}

// With key value
func (u *Uber) With(keysAndValues ...interface{}) Log {
	return &Uber{
		sugar: u.sugar.With(keysAndValues...),
	}
}

// Flush buffered log
func (u *Uber) Flush() error {
	return u.sugar.Sync()
}
