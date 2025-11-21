package main

import (
	"go-todos/internal/config"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/mwinyimoha/commons/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger, err := logging.NewLoggerConfig().BuildLogger()
	if err != nil {
		log.Fatal("failed to initialize logger", err)
	}

	defer func() { _ = logger.Sync() }()

	val := validator.New()
	_, err = config.New(val)
	if err != nil {
		logger.Fatal("failed to load app config", zap.Error(err))
	}
}
