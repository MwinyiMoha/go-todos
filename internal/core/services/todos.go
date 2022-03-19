package services

import (
	"errors"
	"go-todos/internal/core/domain"
	"go-todos/internal/core/ports"
)

type TodoService struct {
	todoRepository ports.TodoRepository
}

func NewTodoService(repository ports.TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: repository,
	}
}

func (t *TodoService) GetTodos() (*[]domain.Todo, error) {
	return t.todoRepository.ListTodos()
}

func (t *TodoService) GetTodo(id string) (*domain.Todo, error) {
	return t.todoRepository.RetrieveTodo(id)
}

func (t *TodoService) CreateTodo(description string) (*domain.Todo, error) {
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	return t.todoRepository.AddTodo(description)
}

func (t *TodoService) UpdateTodo(id string, description string) (*domain.Todo, error) {
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}

	return t.todoRepository.EditTodo(id, description)
}

func (t *TodoService) DeleteTodo(id string) error {
	return t.todoRepository.RemoveTodo(id)
}
