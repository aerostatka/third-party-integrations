package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}

type ZapLogger struct {
	logger *zap.Logger
}

func CreateNewZapLogger() ZapLogger {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	log, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	return ZapLogger{
		logger: log,
	}
}

func (log ZapLogger) Debug(message string) {
	log.logger.Debug(message)
}

func (log ZapLogger) Info(message string) {
	log.logger.Info(message)
}

func (log ZapLogger) Warn(message string) {
	log.logger.Warn(message)
}

func (log ZapLogger) Error(message string) {
	log.logger.Error(message)
}

func (log ZapLogger) Fatal(message string) {
	log.logger.Fatal(message)
}
