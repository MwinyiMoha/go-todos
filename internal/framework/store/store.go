package store

import (
	"go-todos/internal/core/domain"

	"github.com/mwinyimoha/commons/pkg/errors"
)

type Store struct {
	todos []*domain.Todo
}

func New() *Store {
	return &Store{
		todos: []*domain.Todo{},
	}
}

func (s *Store) Close() error {
	return nil
}

func (s *Store) findTodoById(id string) (*domain.Todo, int, error) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			return s.todos[i], i, nil
		}
	}

	return nil, -1, errors.NewErrorf(errors.NotFound, "record not found")
}

func (s *Store) ListTodos() ([]*domain.Todo, error) {
	return s.todos, nil
}

func (s *Store) RetrieveTodo(id string) (*domain.Todo, error) {
	todo, _, err := s.findTodoById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Store) SaveTodo(todo *domain.Todo) error {
	s.todos = append(s.todos, todo)
	return nil
}

func (s *Store) EditTodo(todo *domain.Todo) error {
	_, ix, err := s.findTodoById(todo.ID)
	if err != nil {
		return err
	}

	s.todos[ix] = todo
	return nil

}

func (s *Store) RemoveTodo(id string) error {
	_, ix, err := s.findTodoById(id)
	if err != nil {
		return err
	}

	s.todos = append(s.todos[:ix], s.todos[ix+1:]...)
	return nil
}
