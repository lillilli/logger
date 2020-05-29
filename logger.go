package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger - wrapper for zerolog.Logger type
type Logger = zerolog.Logger

var writer io.Writer = zerolog.ConsoleWriter{Out: os.Stderr}

// Params - represents logger configuration
// Fields:
// HumanFriendly - sets up human friendly log formatting
// MinLevel - sets up min logged level, logger support levels: trace, debug, info, warn, error, fatal, panic
type Params struct {
	HumanFriendly bool   `default:"false" json:"humanFriendly"`
	MinLevel      string `default:"debug" json:"minLevel"`
}

// Init - initialize logger via specified configuration
// params:
// logger support levels: trace, debug, info, warn, error, fatal, panic
func Init(params Params) {
	level, err := zerolog.ParseLevel(params.MinLevel)
	if err != nil {
		log.Warn().Str("unit", "logger").
			Str("level", params.MinLevel).
			Msg("Unexpected initial logger min level, level will be debug")

		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	if !params.HumanFriendly {
		writer = os.Stderr
	}
}

// New - returns new logger instance with unit name
func New(unit string) Logger {
	return log.Output(writer).With().Str("unit", unit).Timestamp().Logger()
}

// NewLogger - returns new logger instance
func NewLogger() Logger {
	return log.Output(writer).With().Timestamp().Logger()
}
