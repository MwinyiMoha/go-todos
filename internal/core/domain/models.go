package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string    `json:"id" bson:"id"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updatedAt"`
}

func NewTodo(description string) *Todo {
	return &Todo{
		ID:          uuid.New().String(),
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
