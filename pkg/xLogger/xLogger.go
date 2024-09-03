package xLogger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const logDir = "logs"

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

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = time.RFC3339Nano
	// Log time always in UTC
	zerolog.TimestampFunc = func() time.Time {
		// time always in UTC
		return time.Now().UTC()
	}
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(logLevel)
	isDev := os.Getenv("IS_DEV")
	if isDev == "true" {
		checkLogFolder(logDir)
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
		fileName := fmt.Sprintf("./%s/%v.log", logDir, time.Now().Unix())
		runLogFile, _ := os.OpenFile(
			fileName,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		multi := zerolog.MultiLevelWriter(consoleWriter, runLogFile)
		Logger = zerolog.New(multi).With().Caller().Timestamp().Logger()
	} else {
		Logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
	}
}

// AddRequestIdToLogger Set requestID in meta for every request
func AddRequestIdToLogger(requestID string) {
	Logger = zerolog.New(os.Stdout).With().
		Caller().Timestamp().Str("requestID", requestID).Logger()
}

func GetLogger(componentName string) zerolog.Logger {
	return Logger.With().Str("component", componentName).Logger()
}

// checkLogFolder creating logs dir if not exist
func checkLogFolder(folder string) {
	// Check if the directory exists
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created:", folder)
	} else {
		fmt.Println("Directory already exists:", folder)
	}
}
