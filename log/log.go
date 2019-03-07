package log

import (
	"errors"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Config struct {
	LogLevel  string
	LogFormat string
}

func New(config Config) (log.Logger, error) {
	var logger log.Logger
	if config.LogFormat == "json" {
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	} else if config.LogFormat == "logfmt" {
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	} else {
		return nil, errors.New("Unrecognized log format: " + config.LogFormat)
	}
	switch config.LogLevel {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	default:
		return nil, errors.New("Unrecognized log level: " + config.LogLevel)
	}

	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	return logger, nil
}
