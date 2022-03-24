package services

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/core/ports"
	"go-todos/internal/utils/exceptions"
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
		return nil, exceptions.New("MISSING_DESCRIPTION", 400)
	}

	return t.todoRepository.AddTodo(description)
}

func (t *TodoService) UpdateTodo(id string, description string) (*domain.Todo, error) {
	if description == "" {
		return nil, exceptions.New("MISSING_DESCRIPTION", 400)
	}

	return t.todoRepository.EditTodo(id, description)
}

func (t *TodoService) DeleteTodo(id string) error {
	return t.todoRepository.RemoveTodo(id)
}
