package db

import (
	"go-todos/internal/core/domain"
	"go-todos/internal/utils/exceptions"
	"go-todos/internal/utils/factories"

	"github.com/mwinyimoha/commons/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) ListTodos() ([]*domain.Todo, error) {
	ctx, cancel := r.getContext()
	defer cancel()

	todos := make([]*domain.Todo, 0)

	findOptions := options.Find().SetSort(bson.M{"createdAt": -1})
	cursor, err := r.todos.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	if err := cursor.All(ctx, &todos); err != nil {
		return nil, errors.NewErrorf(errors.Internal, "could not fetch records")
	}

	return todos, nil
}

func (r *Repository) RetrieveTodo(id string) (*domain.Todo, error) {
	filter := map[string]any{"id": id}
	return r.getTodo(filter)
}

func (r *Repository) SaveTodo(todo *domain.Todo) error {
	ctx, cancel := r.getContext()
	defer cancel()

	_, err := r.todos.InsertOne(ctx, todo)
	if err != nil {
		return errors.NewErrorf(errors.Internal, "could not save record")
	}

	return nil
}

func (r *Repository) EditTodo(todo *domain.Todo) error {
	ctx, cancel := factories.NewContext()
	defer cancel()

	_, err := r.todos.ReplaceOne(ctx, bson.M{"id": todo.ID}, todo)
	if err != nil {
		return errors.NewErrorf(errors.Internal, "could not update record")
	}

	return nil
}

func (r *Repository) RemoveTodo(id string) error {
	ctx, cancel := r.getContext()
	defer cancel()

	_, err := r.todos.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return errors.NewErrorf(errors.Internal, "could not delete record")
	}

	return nil
}

func (r *Repository) getTodo(filter map[string]any) (*domain.Todo, error) {
	ctx, cancel := r.getContext()
	defer cancel()

	var todo domain.Todo

	result := r.todos.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewErrorf(errors.NotFound, "record not found")
		}

		return nil, errors.NewErrorf(errors.Internal, "could not fetch record")
	}

	if err := result.Decode(&todo); err != nil {
		return nil, errors.NewErrorf(errors.Internal, "could not serialize record")
	}

	return &todo, nil
}
