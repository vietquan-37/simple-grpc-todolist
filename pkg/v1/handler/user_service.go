package handler

import (
	"context"

	"github.com/vietquan-37/todo-list/pb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	model := convertUser(req)
	user, err := server.UserUseCase.CreateUser(model)
	if err != nil {
		return nil, err
	}
	rsp := &pb.CreateUserResponse{
		User: convertUserResponse(*user),
	}
	return rsp, nil
}
