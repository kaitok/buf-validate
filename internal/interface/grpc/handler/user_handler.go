package handler

import (
	"context"

	userpb "buf-validate/tools/grpc/v1/user"

	"buf-validate/internal/application/service"
	"buf-validate/internal/infrastructure/repository"
)

var getUserUsecase = service.GetUserUsecase{
	UserRepo: &repository.UserRepositoryImpl{},
}

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := getUserUsecase.Execute(string(rune(req.GetId())))
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:       req.GetId(),
			Username: user.Name,
		},
	}, nil
}
