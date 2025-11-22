package main

import (
	"context"
	"fmt"
	"go-todos/internal/config"
	"go-todos/internal/core/app"
	"go-todos/internal/core/ports"
	"go-todos/internal/framework/api"
	"go-todos/internal/framework/db"
	"go-todos/internal/framework/store"
	"log"
	"net/http"
	"os/signal"
	"syscall"

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

	svc := app.NewService(repo)
	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	ch := make(chan error, 1)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	switch cfg.APIType {
	case "rest":
		router := api.NewRouter(svc, cfg.Debug)
		srv := &http.Server{
			Addr:    addr,
			Handler: router.Engine,
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal("failed to start server", zap.Error(err))
				ch <- err
			}
		}()
	default:
		logger.Fatal("unsupported api type", zap.String("api_type", cfg.APIType))
	}

	select {
	case <-ctx.Done():
		logger.Info("initiating graceful shutdown")
	case err := <-ch:
		logger.Fatal("application stopped with error", zap.Error(err))
	}
}
