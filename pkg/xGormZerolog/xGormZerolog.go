package xgormzerolog

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	gormzerolog "github.com/vitaliy-art/gorm-zerolog"
	gormlog "gorm.io/gorm/logger"
)

var (
	dbLogLevelMap = map[string]gormlog.LogLevel{
		"silent": 1,
		"error":  2,
		"warn":   3,
		"info":   4,
	}
)

func NewDbLogger(logger *zerolog.Logger) *gormzerolog.GormLogger {
	dbLogger := gormzerolog.NewGormLogger()
	dbLogger.WithInfo(
		func() gormzerolog.Event {
			return &gormzerolog.GormLoggerEvent{Event: logger.Info()}
		},
	)
	dbLogger.WithWarn(
		func() gormzerolog.Event {
			return &gormzerolog.GormLoggerEvent{Event: logger.Warn()}
		},
	)
	dbLogger.WithError(
		func() gormzerolog.Event {
			return &gormzerolog.GormLoggerEvent{Event: logger.Error()}
		},
	)

	// Get log level from environment variable
	dbLogLevelStr := strings.ToLower(os.Getenv("DB_LOG_LEVEL"))

	// set default log level for driver
	var dbLogLevel gormlog.LogLevel = dbLogLevelMap["warn"]
	if dbLogLevelStr != "" {
		dbLogLevel = dbLogLevelMap[dbLogLevelStr]
	}
	dbLogger.LogMode(dbLogLevel)
	return dbLogger
}
