package handler

import (
	"context"

	"github.com/bufbuild/connect-go"

	"go-grpc-api/internal/application/service"
	"go-grpc-api/internal/infrastructure/repository"
	userpb "go-grpc-api/tools/grpc/go_grpc_api/v1"
)

var getUserUsecase = service.GetUserUsecase{
	UserRepo: &repository.UserRepositoryImpl{},
}

type UserServiceServer struct{}

func (s *UserServiceServer) GetUser(ctx context.Context, req *connect.Request[userpb.GetUserRequest]) (*connect.Response[userpb.GetUserResponse], error) {
	user, err := getUserUsecase.Execute(string(rune(req.Msg.GetId())))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&userpb.GetUserResponse{
		User: &userpb.User{
			Id:       req.Msg.GetId(),
			Username: user.Name,
		},
	}), nil
}
