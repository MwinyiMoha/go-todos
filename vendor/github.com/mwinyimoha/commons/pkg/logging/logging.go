package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Config zap.Config
}

func NewLoggerConfig() *LoggerConfig {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	return &LoggerConfig{Config: cfg}
}

func (lc *LoggerConfig) BuildLogger() (*zap.Logger, error) {
	return lc.Config.Build()
}
