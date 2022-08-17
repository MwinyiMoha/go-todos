package db

import (
	"fmt"
	"go-todos/internal/core/domain"
	"go-todos/internal/utils/config"
	"go-todos/internal/utils/exceptions"
	"go-todos/internal/utils/factories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	// Create Mongo database URL
	URL := fmt.Sprintf("mongodb://%v:%v@%v:%v/?authSource=%v", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.AuthDB)

	// Pull a context and a cancel function
	ctx, cancel := factories.NewContext()
	defer cancel()

	client, err := connectDatabase(ctx, URL)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Client:     client,
		Collection: client.Database(cfg.DBName).Collection(cfg.CollectionName),
	}, nil
}

func (r *Repository) ListTodos() (*[]domain.Todo, error) {
	/*
		Returns all todos in the database
	*/
	ctx, cancel := factories.NewContext()
	defer cancel()

	todos := make([]domain.Todo, 0)

	findOptions := options.Find().SetSort(bson.M{"createdAt": -1})
	cursor, err := r.Collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	if err := cursor.All(ctx, &todos); err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	return &todos, nil
}

func (r *Repository) RetrieveTodo(id string) (*domain.Todo, error) {
	// Returns a todo with a given [id]
	return r.GetTodoOr404(id)
}

func (r *Repository) AddTodo(description string) (*domain.Todo, error) {
	/*
		Adds a new todo in the collection
		Returns the added todo
	*/
	ctx, cancel := factories.NewContext()
	defer cancel()

	todo := domain.NewTodo(description)
	_, err := r.Collection.InsertOne(ctx, *todo)
	if err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	return todo, nil
}

func (r *Repository) EditTodo(id string, description string) (*domain.Todo, error) {
	/*
		Updates a todo with a given [id], with new description in the argument passed to this method
	*/
	todo, err := r.GetTodoOr404(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := factories.NewContext()
	defer cancel()

	todo.Description = description
	_, err = r.Collection.ReplaceOne(ctx, bson.M{"id": id}, todo)
	if err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	return todo, nil
}

func (r *Repository) RemoveTodo(id string) error {
	/*
		Removes a todo with the given [id] from the collection
	*/
	_, err := r.GetTodoOr404(id)
	if err != nil {
		return err
	}

	ctx, cancel := factories.NewContext()
	defer cancel()

	_, err = r.Collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return exceptions.New(err.Error(), 500)
	}

	return nil
}

func (r *Repository) GetTodoOr404(id string) (*domain.Todo, error) {
	/*
		Returns a todo with the given [id] if found, else returns an error
	*/
	ctx, cancel := factories.NewContext()
	defer cancel()

	var todo domain.Todo

	result := r.Collection.FindOne(ctx, bson.M{"id": id})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exceptions.New(err.Error(), 404)
		}

		return nil, exceptions.New(err.Error(), 500)
	}

	if err := result.Decode(&todo); err != nil {
		return nil, exceptions.New(err.Error(), 500)
	}

	return &todo, nil
}
