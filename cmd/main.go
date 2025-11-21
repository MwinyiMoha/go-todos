package main

import (
	"go-todos/internal/config"
	"go-todos/internal/core/app"
	"go-todos/internal/core/ports"
	"go-todos/internal/framework/db"
	"go-todos/internal/framework/store"
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
	cfg, err := config.New(val)
	if err != nil {
		logger.Fatal("failed to load app config", zap.Error(err))
	}

	var repo ports.AppRepository
	if cfg.StoreType == "inmemory" {
		repo = store.New()
	} else {
		repo, err = db.NewRepository(cfg)
		if err != nil {
			logger.Fatal("failed to initialize db repository", zap.Error(err))
		}
	}

	_ = app.NewService(repo)
}
