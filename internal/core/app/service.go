package app

import "go-todos/internal/core/ports"

type Service struct {
	repository ports.AppRepository
}

func NewService(r ports.AppRepository) *Service {
	return &Service{
		repository: r,
	}
}
