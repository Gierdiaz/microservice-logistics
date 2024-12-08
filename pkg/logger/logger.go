package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func InitLogger() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Sync() {
	log.Sync()
}
