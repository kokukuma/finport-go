package log

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a new zap logger with the given log level.
func New(level string) (*zap.Logger, error) {
	l, err := logLevel(level)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level")
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(l)
	config.DisableStacktrace = true
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	return config.Build()
}

// NewDiscard creates logger which output to ioutil.Discard.
// This can be used for testing.
func NewDiscard() *zap.Logger {
	return zap.NewNop()
}

func logLevel(level string) (zapcore.Level, error) {
	level = strings.ToUpper(level)
	var l zapcore.Level
	switch level {
	case "DEBUG":
		l = zapcore.DebugLevel
	case "INFO":
		l = zapcore.InfoLevel
	case "ERROR":
		l = zapcore.ErrorLevel
	default:
		return l, errors.Errorf("invalid loglevel: %s", level)
	}
	return l, nil
}
