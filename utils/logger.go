package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func prettyLogger() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{Out: os.Stderr,
		TimeFormat: zerolog.TimeFormatUnix,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
	}
}

func CreateLogger() {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(prettyLogger()).With().Caller().Logger()
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = log.Output(prettyLogger()).With().Caller().Logger()
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Logger = log.Output(prettyLogger()).With().Caller().Logger()
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		log.Logger = log.Output(prettyLogger()).With().Caller().Logger()
	case "none":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		log.Logger = log.With().Caller().Logger()
	}
}
