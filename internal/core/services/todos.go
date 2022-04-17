package services

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/core/ports"
	"go-todos/internal/utils/exceptions"
)

type TodoService struct {
	TodoRepository ports.TodoRepository
}

func NewTodoService(repository ports.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepository: repository,
	}
}

func (t *TodoService) GetTodos() (*[]domain.Todo, error) {
	return t.TodoRepository.ListTodos()
}

func (t *TodoService) GetTodo(id string) (*domain.Todo, error) {
	return t.TodoRepository.RetrieveTodo(id)
}

func (t *TodoService) CreateTodo(description string) (*domain.Todo, error) {
	if description == "" {
		return nil, exceptions.New("description is required", 400)
	}

	return t.TodoRepository.AddTodo(description)
}

func (t *TodoService) UpdateTodo(id string, description string) (*domain.Todo, error) {
	if description == "" {
		return nil, exceptions.New("description is required", 400)
	}

	return t.TodoRepository.EditTodo(id, description)
}

func (t *TodoService) DeleteTodo(id string) error {
	return t.TodoRepository.RemoveTodo(id)
}
