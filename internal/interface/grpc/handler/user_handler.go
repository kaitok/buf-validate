package handler

import (
	"context"

	"github.com/bufbuild/connect-go"

	"buf-validate-example/internal/application/service"
	"buf-validate-example/internal/infrastructure/repository"
	userpb "buf-validate-example/tools/grpc/buf_validate_example/v1"
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
