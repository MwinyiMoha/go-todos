package main

import (
	"context"
	"fmt"
	"go-todos/internal/config"
	"go-todos/internal/core/app"
	"go-todos/internal/core/ports"
	"go-todos/internal/framework/api"
	"go-todos/internal/framework/cli"
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

var storeFactories = map[string]func(cfg *config.Config) (ports.AppRepository, error){
	"inmemory": func(_ *config.Config) (ports.AppRepository, error) {
		return store.New(), nil
	},
	"database": func(cfg *config.Config) (ports.AppRepository, error) {
		return db.NewRepository(cfg)
	},
}

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

	factory, ok := storeFactories[cfg.Store]
	if !ok {
		logger.Fatal("unsupported store type", zap.String("store_type", cfg.Store))
	}

	repo, err := factory(cfg)
	if err != nil {
		logger.Fatal("failed to initialize repository", zap.Error(err))
	}

	svc := app.NewService(repo)

	switch cfg.Interface {
	case "http":
		router := api.NewRouter(svc, logger, cfg.Debug)
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
			Handler: router.Engine,
		}

		ch := make(chan error, 1)
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
		defer stop()

		go func() {
			logger.Info(
				"starting server",
				zap.String("app_name", cfg.AppName),
				zap.String("app_version", cfg.AppVersion),
				zap.Int("port", cfg.ServerPort),
			)

			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("failed to start server", zap.Error(err))
				ch <- err
			}
		}()

		select {
		case <-ctx.Done():
			logger.Info("initiating graceful shutdown")

			if err := srv.Shutdown(context.Background()); err != nil {
				logger.Error("failed graceful shutdown", zap.Error(err))
			}

			if err := repo.Close(); err != nil {
				logger.Error("failed to close repository", zap.Error(err))
			}

			logger.Info("application stopped")
		case err := <-ch:
			logger.Fatal("application stopped with error", zap.Error(err))
		}
	case "cli":
		cmd := cli.NewCMD(cfg, svc)
		if err := cmd.Execute(); err != nil {
			logger.Fatal("execution failed", zap.Error(err))
		}
	default:
		logger.Fatal("unsupported interface type", zap.String("api_type", cfg.Interface))
	}
}
