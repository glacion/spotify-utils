package config

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ConfigureLogger sets up zerolog
func ConfigureLogger(verbose bool) *zerolog.Logger {
	logger := log.Output(os.Stdout)

	if verbose {
		logger = logger.Level(zerolog.DebugLevel)
	} else {
		logger = logger.Level(zerolog.InfoLevel)
	}

	// Use pretty output in a TTY
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) ||
		isatty.IsTerminal(os.Stdout.Fd()) {

		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
		logger.Debug().Msg("enabled pretty logging")
	}

	log.Logger = logger
	return &logger
}
