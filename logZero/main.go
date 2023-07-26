package main

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	// If you set zerolog.TimeFieldFormat to an empty string,
	// logs will write with UNIX time
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Msg("basic: Hello world")
	log.Debug().Msg("basic: Debug Info")
	log.Log().
		Str("foo", "bar").
		Msg("someMess")
	log.Info().
		Str("foo", "bar").
		Msg("someMess")
	logger.Info().Msg("Hello world")
	logger.Debug().Msg("Debug Info")
	logger.Warn().Msg("This is a warning")
	logger.Log().
		Str("foo", "bar").
		Msg("someMess")

	err := errors.New("seems we have an error here")
	logger.Error().Err(err).Msg("unable to open file:")
	log.Error().Err(err).Msgf("unable to open file: ")

	logger.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	// service := "myservice"
	// logger.Fatal().
	// 	Err(err).
	// 	Str("service", service).
	// 	Msgf("Cannot start %s", service)

}
