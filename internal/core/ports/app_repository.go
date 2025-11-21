package ports

import "go-todos/internal/core/domain"

type AppRepository interface {
	ListTodos() ([]*domain.Todo, error)
	RetrieveTodo(id string) (*domain.Todo, error)
	SaveTodo(todo *domain.Todo) error
	EditTodo(todo *domain.Todo) error
	RemoveTodo(id string) error
}
