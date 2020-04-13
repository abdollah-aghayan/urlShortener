package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerInterface interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
}

var logger *zap.Logger

func init() {
	// init Zap logger

	config := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ = config.Build()

}

func Info(msg string, args ...zap.Field) {
	logger.Info(msg, args...)
	logger.Sync()
}

func Debug(msg string, args ...zap.Field) {
	logger.Debug(msg, args...)
	logger.Sync()
}

func Error(msg string, err error, args ...zap.Field) {
	args = append(args, zap.NamedError("error", err))
	logger.Error(msg, args...)
	logger.Sync()
}
