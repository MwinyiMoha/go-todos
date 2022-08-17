package main

import (
	"go-todos/internal/core/services"
	"go-todos/internal/framework/db"
	"go-todos/internal/framework/rpcserver"
	pb "go-todos/internal/framework/rpcserver/proto"
	"go-todos/internal/utils/config"
	"go-todos/internal/utils/factories"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf := config.New()

	repository, err := db.NewRepository(conf)
	if err != nil {
		log.Fatalf("DB_CONN_ERROR -- %v", err)
	}

	service := services.NewTodoService(repository)
	srv := rpcserver.NewRPCServer(service)

	lis, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		log.Fatalf("NET_ERROR -- %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoRPCServiceServer(s, srv)

	go func() {
		log.Printf("Starting RPC server on port: %v\n", conf.Port)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("SERVE_ERRROR -- %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	received := <-c

	func() {
		log.Printf("Handling (%v): Initiating graceful server shutdown!!", received)

		done := make(chan int)
		go func(c chan int) {
			s.GracefulStop()
			c <- 1
		}(done)

		countdown := time.NewTimer(time.Duration(conf.AppTimeout))
		select {
		case <-done:
			break
		case <-countdown.C:
			s.Stop()
		}
	}()

	func() {
		log.Println("Closing client connections!!")

		ctx, cancel := factories.NewContext()
		defer cancel()

		if err := repository.Client.Disconnect(ctx); err != nil {
			log.Fatalf("DB_DISCONNECT_ERROR -- %v", err)
		}
	}()
}
