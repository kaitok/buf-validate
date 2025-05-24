package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-grpc-api/internal/interface/grpc/handler"
	grpc "go-grpc-api/tools/grpc/go_grpc_api/v1/go_grpc_apiv1connect"
)

func main() {

	// Run grpc server
	grpcPort := ":8081"
	go func() {
		log.Printf("gRPC server (connect-go) listening on %s", grpcPort)
		mux := http.NewServeMux()
		mux.Handle(grpc.NewTaskServiceHandler(&handler.TaskServiceServer{}))
		mux.Handle(grpc.NewUserServiceHandler(&handler.UserServiceServer{}))
		if err := http.ListenAndServe(grpcPort, mux); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}
