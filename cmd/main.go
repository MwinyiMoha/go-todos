package main

import (
	"go-todos/internal/core/services"
	"go-todos/internal/framework/api"
	"go-todos/internal/framework/store"
	"log"
)

func main() {
	repository := store.NewStore()
	service := services.NewTodoService(repository)
	controller := api.NewHTTPController(service)
	router := api.NewRouter(controller)
	router.AddRoutes()

	log.Fatal(router.Serve())
}
