package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		EncoderConfig:     zapcore.EncoderConfig{
			LevelKey:      "level",
			TimeKey:       "time",
			MessageKey:    "msg",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() interface{} {
	return log
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	_ = log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	_ = log.Sync()
}

func Warn(msg string, tags ...zap.Field) {
	log.Warn(msg, tags...)
	_ = log.Sync()
}
