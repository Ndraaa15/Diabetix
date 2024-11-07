package config

import (
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap(env *env.Env) *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logger, err := config.Build()
	if err != nil {
		zap.S().Fatalf("Unable to create zap logger: %v", err)
	}

	return logger
}
