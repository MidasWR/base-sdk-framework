package midas

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
)

type LoggerConfig struct {
	LogLevel string
	Out      io.Writer
}

func InitLogger(cfg LoggerConfig, service string) zerolog.Logger {
	cw := zerolog.ConsoleWriter{Out: cfg.Out}
	cw.FormatFieldName = func(i interface{}) string {
		if s, ok := i.(string); ok && s == "service" {
			return "\x1b[35m " + s + ":\x1b[0m"
		}
		return fmt.Sprintf("%s:", i)
	}
	if cfg.LogLevel == "DEV" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else if cfg.LogLevel == "PROD" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	logger := zerolog.New(cw).With().
		Timestamp().
		Caller().
		Str("service", service).
		Logger()

	return logger
}
