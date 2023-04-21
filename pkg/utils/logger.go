package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sort"
	"sync/atomic"
)

// sharedLogger holds the global LoggerInterface
var sharedLogger atomic.Value

var zapLoggerConfig *zap.Config

type LogLevel string
type LogFields map[string]interface{}

type LoggerInterface interface {
	WithField(key string, value interface{}) LoggerInterface
	WithFields(fields LogFields) LoggerInterface
	WithError(err error) LoggerInterface
	Debug(message string, args ...zapcore.Field)
	Info(message string, args ...zapcore.Field)
	Warn(message string, args ...zapcore.Field)
	Error(message string, args ...zapcore.Field)
	Fatal(message string, args ...zapcore.Field)
	Panic(message string, args ...zapcore.Field)
}

type ZapLogger struct {
	*zap.Logger
}

func newZapLogger() LoggerInterface {
	logger, err := zapLoggerConfig.Build()
	if err != nil {
		// if we can't build the logger, it will terminate the program
		os.Exit(1)
	}

	return &ZapLogger{logger}
}

func init() {
	zapLoggerConfig = setupLoggerConfig()
}

func setupLoggerConfig() *zap.Config {
	return &zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// Logger returns the global log entry.
func Logger() LoggerInterface {
	if sharedLogger.Load() == nil {
		sharedLogger.Store(newZapLogger())
	}
	return sharedLogger.Load().(LoggerInterface)
}

// WithField returns the logger at the supplied field.
func (zl *ZapLogger) WithField(key string, value interface{}) LoggerInterface {
	newLogger := zl.Logger.WithOptions(zap.Fields(zap.Any(key, value)))
	return &ZapLogger{newLogger}
}

// WithFields returns the logger at the supplied fields.
func (zl *ZapLogger) WithFields(fields LogFields) LoggerInterface {
	// sort the keys of the fields map
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	// the purpose of sorting the keys is to ensure that the order of the fields
	sort.Strings(keys)

	// iterate over the sorted keys and append the corresponding zap.Field
	zapFields := make([]zap.Field, 0, len(fields))
	for _, k := range keys {
		v := fields[k]
		zapFields = append(zapFields, zap.Any(k, v))
	}

	newLogger := zl.Logger.WithOptions(zap.Fields(zapFields...))
	return &ZapLogger{newLogger}
}

// WithError returns the logger at the supplied error.
func (zl *ZapLogger) WithError(err error) LoggerInterface {
	newLogger := zl.Logger.WithOptions(zap.Fields(zap.Error(err)))
	return &ZapLogger{newLogger}
}
