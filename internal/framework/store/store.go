package store

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/utils/exceptions"
)

type Store struct {
	todos []domain.Todo
}

func NewStore() (*Store, error) {
	return &Store{
		todos: []domain.Todo{},
	}, nil
}

func (s *Store) FindTodoById(id string) (*domain.Todo, int, error) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			return &s.todos[i], i, nil
		}
	}

	return nil, -1, exceptions.New("TODO_NOT_FOUND", 404)
}

func (s *Store) ListTodos() (*[]domain.Todo, error) {
	return &s.todos, nil
}

func (s *Store) RetrieveTodo(id string) (*domain.Todo, error) {
	todo, _, err := s.FindTodoById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Store) AddTodo(description string) (*domain.Todo, error) {
	todo := domain.NewTodo(description)
	s.todos = append(s.todos, *todo)

	return todo, nil
}

func (s *Store) EditTodo(id string, description string) (*domain.Todo, error) {
	todo, ix, err := s.FindTodoById(id)
	if err != nil {
		return nil, err
	}

	todo.Description = description
	s.todos[ix] = *todo
	return todo, nil

}

func (s *Store) RemoveTodo(id string) error {
	_, ix, err := s.FindTodoById(id)
	if err != nil {
		return err
	}

	s.todos = append(s.todos[:ix], s.todos[ix+1:]...)
	return nil
}
