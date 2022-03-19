package domain

import (
	"time"

	"github.com/google/uuid"
)

type TodoPayload struct {
	Description string `json:"description"`
}

type Todo struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(desc string) *Todo {
	return &Todo{
		ID:          uuid.New().String(),
		Description: desc,
		CreatedAt:   time.Now(),
	}
}
