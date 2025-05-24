package main

import (
	"buf-validate-example/internal/interface/grpc/handler"
	taskpb "buf-validate-example/tools/grpc/v1/task"
	userpb "buf-validate-example/tools/grpc/v1/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Run grpc server
	grpcPort := ":8081"
	go func() {
		log.Printf("gRPC server (connect-go) listening on %s", grpcPort)
		mux := http.NewServeMux()
		mux.Handle(taskpb.NewTaskServiceHandler(&handler.TaskServiceServer{}))
		mux.Handle(userpb.NewUserServiceHandler(&handler.UserServiceServer{}))
		if err := http.ListenAndServe(grpcPort, mux); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}
