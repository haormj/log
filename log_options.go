package log

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type level string

type encoder string

// Options uber log
type Options struct {
	Encoder        encoder
	Level          level
	MessageKey     string
	LevelKey       string
	TimeKey        string
	NameKey        string
	CallerKey      string
	StacktraceKey  string
	LineEnding     string
	EncodeLevel    zapcore.LevelEncoder
	EncodeTime     zapcore.TimeEncoder
	EncodeDuration zapcore.DurationEncoder
	EncodeCaller   zapcore.CallerEncoder
	EncodeName     zapcore.NameEncoder
	Filename       string
	MaxSize        int
	MaxAge         int
	MaxBackups     int
	LocalTime      bool
	Compress       bool
}

// Option sugar for options
type Option func(*Options)

var (
	DEG level = "debug"
	INF level = "info"
	WRN level = "warn"
	ERR level = "error"
)

var (
	JSON    encoder = "json"
	Console encoder = "console"
)

func newOptions(opts ...Option) (options Options) {
	options = Options{
		Encoder:        Console,
		Level:          INF,
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    LevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		MaxSize:        1,
		MaxAge:         7,
		MaxBackups:     50,
		LocalTime:      true,
		Compress:       false,
	}

	for _, o := range opts {
		o(&options)
	}

	return
}

// Filename log file name
func Filename(name string) Option {
	return func(o *Options) {
		o.Filename = name
	}
}

// Level log level
// default is INF
func Level(l level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

// ParseLevel parse string to level
func ParseLevel(l string) level {
	switch l {
	case "debug":
		return DEG
	case "info":
		return INF
	case "warn":
		return WRN
	case "error":
		return ERR
	}
	return INF
}

// Encoder choose json/console
// default is console
func Encoder(e encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}

// ParseEncoder parse string to encoder
func ParseEncoder(e string) encoder {
	switch e {
	case "json":
		return JSON
	case "console":
		return Console
	}
	return Console
}

// LevelEncoder format level name
func LevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var level string
	switch l {
	case zapcore.DebugLevel:
		level = "DEG"
	case zapcore.InfoLevel:
		level = "INF"
	case zapcore.WarnLevel:
		level = "WRN"
	case zapcore.ErrorLevel:
		level = "ERR"
	}
	enc.AppendString(level)
}

// TimeEncoder format time
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
