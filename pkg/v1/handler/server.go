package handler

import (
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/usecase/interfaces"
)

type Server struct {
	pb.UnimplementedTodoListServer
	interfaces.UserUseCase
}

func NewServer(userUseCase interfaces.UserUseCase) *Server {
	return &Server{
		UserUseCase: userUseCase,
	}
}
