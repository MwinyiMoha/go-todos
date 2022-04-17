package main

import (
	"go-todos/internal/core/services"
	"go-todos/internal/framework/api"
	"go-todos/internal/framework/db"
	"go-todos/internal/utils/config"
	"go-todos/internal/utils/factories"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// repository, err := store.NewStore()

	repository, err := db.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	service := services.NewTodoService(repository)
	controller := api.NewHTTPController(service)
	router := api.NewRouter(controller)
	router.AddRoutes()

	conf := config.New()
	server := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: router.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("ERROR -- %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	received := <-c

	ctx, cancel := factories.NewContext()
	defer cancel()

	func() {
		log.Printf("Handling (%v): Initiating graceful shutdown!!", received)

		log.Printf("Closing database connection!!")
		if err := repository.Client.Disconnect(ctx); err != nil {
			log.Fatalf("Error closing database connection-- %v", err)
		}

		log.Printf("Shutting down server!!")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown error -- %v", err)
		}
	}()

}
