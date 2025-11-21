package ports

import "go-todos/internal/core/domain"

type AppService interface {
	GetTodos() ([]*domain.Todo, error)
	GetTodo(id string) (*domain.Todo, error)
	CreateTodo(description string) (*domain.Todo, error)
	UpdateTodo(id string, description string) (*domain.Todo, error)
	DeleteTodo(id string) error
}
