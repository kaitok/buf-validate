package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-grpc-api/internal/interface/grpc/handler"
	grpc "go-grpc-api/tools/grpc/go_grpc_api/v1/go_grpc_apiv1connect"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Run grpc server
	grpcPort := ":8081"
	go func() {
		log.Printf("gRPC server (connect-go) listening on %s", grpcPort)
		pathTask, handlerTask := grpc.NewTaskServiceHandler(&handler.TaskServiceServer{})
		r.Handle(pathTask, handlerTask)
		pathUser, handlerUser := grpc.NewUserServiceHandler(&handler.UserServiceServer{})
		r.Handle(pathUser, handlerUser)
		if err := http.ListenAndServe(grpcPort, r); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}
