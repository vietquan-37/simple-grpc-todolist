package handler

import (
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"
	"github.com/vietquan-37/todo-list/util"
)

type Server struct {
	pb.UnimplementedTodoListServer
	Repo     interfaces.UserRepo
	TaskRepo interfaces.TaskRepo
	Token    util.JwtMaker
}

func NewServer(repo interfaces.UserRepo, taskRepo interfaces.TaskRepo, token util.JwtMaker) *Server {
	return &Server{
		Repo:     repo,
		TaskRepo: taskRepo,
		Token:    token,
	}
}
