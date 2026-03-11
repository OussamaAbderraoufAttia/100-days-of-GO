package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// Setup zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(os.Stdout)

	// Log with structured fields
	log.Info().
		Str("component", "auth").
		Int("user_id", 12345).
		Msg("User successfully logged in")

	log.Error().
		Str("error", "database connection timeout").
		Msg("Failed to reach database")
}
