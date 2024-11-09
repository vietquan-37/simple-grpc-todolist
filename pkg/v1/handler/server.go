package handler

import (
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"
)

type Server struct {
	pb.UnimplementedTodoListServer
	Repo interfaces.UserRepo
}

func NewServer(repo interfaces.UserRepo) *Server {
	return &Server{
		Repo: repo,
	}
}
