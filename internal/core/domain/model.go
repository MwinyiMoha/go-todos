package domain

import (
	"time"

	"github.com/google/uuid"
)

type TodoPayload struct {
	Description string `json:"description"`
}

type Todo struct {
	ID          string    `json:"id" bson:"id"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"createdAt"`
}

func NewTodo(desc string) *Todo {
	return &Todo{
		ID:          uuid.New().String(),
		Description: desc,
		CreatedAt:   time.Now(),
	}
}
