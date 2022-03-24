package main

import (
	"context"
	"go-todos/internal/core/services"
	"go-todos/internal/framework/api"
	"go-todos/internal/framework/store"
	"go-todos/internal/utils/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf := config.New()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.AppTimeout)*time.Second)
	defer cancel()

	repository := store.NewStore()
	service := services.NewTodoService(repository)
	controller := api.NewHTTPController(service)
	router := api.NewRouter(controller)
	router.AddRoutes()

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
	<-c

	log.Println("Initiating graceful shutdown!!")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
