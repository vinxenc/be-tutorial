package log

import (
	"os"

	"github.com/rs/zerolog"
)

type ConsoleTransport struct {
	logger zerolog.Logger
}

func NewConsoleTransport() Transport {
	return &ConsoleTransport{
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}

func (ct *ConsoleTransport) log(json *JsonMessage) error {
	logger := ct.logger

	switch json.Level {
	case "INFO":
		logger.Info().Msg(json.Message)
	case "ERROR":
		logger.Error().Msg(json.Message)
	case "WARN":
		logger.Warn().Msg(json.Message)
	case "DEBUG":
		logger.Debug().Msg(json.Message)
	case "FATAL":
		logger.Fatal().Msg(json.Message)
	}

	return nil
}
