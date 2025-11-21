package app

import (
	"go-todos/internal/core/domain"

	"github.com/google/uuid"
	"github.com/mwinyimoha/commons/pkg/errors"
)

func (svc *Service) GetTodos() ([]*domain.Todo, error) {
	return svc.repository.ListTodos()
}

func (svc *Service) GetTodo(id string) (*domain.Todo, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewErrorf(errors.BadRequest, "invalid id")
	}

	return svc.repository.RetrieveTodo(id)
}

func (svc *Service) CreateTodo(description string) (*domain.Todo, error) {
	if description == "" {
		return nil, errors.NewErrorf(errors.BadRequest, "description is required")
	}

	todo := domain.NewTodo(description)
	if err := svc.repository.SaveTodo(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (svc *Service) UpdateTodo(id string, description string) (*domain.Todo, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.NewErrorf(errors.BadRequest, "invalid id")
	}

	if description == "" {
		return nil, errors.NewErrorf(errors.BadRequest, "description is required")
	}

	todo, err := svc.repository.RetrieveTodo(id)
	if err != nil {
		return nil, err
	}

	todo.Description = description
	if err := svc.repository.EditTodo(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (svc *Service) DeleteTodo(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.NewErrorf(errors.BadRequest, "invalid id")
	}

	_, err = svc.repository.RetrieveTodo(id)
	if err != nil {
		return err
	}

	return svc.repository.RemoveTodo(id)
}
