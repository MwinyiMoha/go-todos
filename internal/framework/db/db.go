package db

import (
	"context"
	"go-todos/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "todos"
)

type Repository struct {
	client  *mongo.Client
	todos   *mongo.Collection
	timeout time.Duration // in seconds
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	timeout := time.Duration(cfg.AppTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &Repository{
		client:  client,
		todos:   client.Database(cfg.DatabaseName).Collection(collectionName),
		timeout: timeout,
	}, nil
}

func (r *Repository) Close() error {
	ctx, cancel := r.getContext()
	defer cancel()

	return r.client.Disconnect(ctx)
}

func (r *Repository) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.timeout)
}
