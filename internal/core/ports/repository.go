package ports

import "go-todos/internal/core/domain"

type TodoRepository interface {
	ListTodos() (*[]domain.Todo, error)
	RetrieveTodo(id string) (*domain.Todo, error)
	AddTodo(description string) (*domain.Todo, error)
	EditTodo(id string, description string) (*domain.Todo, error)
	RemoveTodo(id string) error
}
