package xLogger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	// Get log level from environment variable
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))

	// Default log level if not provided or invalid
	logLevel := zerolog.InfoLevel

	if logLevelStr != "" {
		switch logLevelStr {
		case "debug":
			logLevel = zerolog.DebugLevel
		case "info":
			logLevel = zerolog.InfoLevel
		case "warn":
			logLevel = zerolog.WarnLevel
		case "error":
			logLevel = zerolog.ErrorLevel
		case "fatal":
			logLevel = zerolog.FatalLevel
		case "panic":
			logLevel = zerolog.PanicLevel
		}
	}

	// UNIX Time is faster and smaller than most timestamps, but logs systems wait RFC3339
	zerolog.TimeFieldFormat = time.RFC3339Nano
	// Log time always in UTC
	zerolog.TimestampFunc = func() time.Time {
		// time always in UTC
		return time.Now().UTC()
	}
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(logLevel)
	Logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
}

// AddRequestIdToLogger Set requestID in meta for every request
func AddRequestIdToLogger(requestID string) {
	Logger = zerolog.New(os.Stdout).With().
		Caller().Timestamp().Str("requestID", requestID).Logger()
}
