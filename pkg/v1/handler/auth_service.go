package handler

import (
	"context"
	"errors"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	validator, err := protovalidate.New()
	if err != nil {
		panic(err)
	}
	if err := validator.Validate(req); err != nil {
		violation := ErrorResponses(err)
		return nil, invalidArgumentError(violation)
	}
	hashPasword, err := util.HashedPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while hashing password: %s", err)
	}
	req.Password = hashPasword
	model := convertUser(req)

	user, err := server.Repo.CreateUser(model)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, status.Errorf(codes.AlreadyExists, "email %s already register before", req.Email)
		}
		return nil, status.Errorf(codes.Internal, "error while creating user: %s", err)
	}

	return convertUserResponse(*user), nil
}

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.Repo.GetUserByEmail(req.GetEmail())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, status.Errorf(codes.NotFound, "user with email %s not found", req.GetEmail())

		}
		return nil, status.Errorf(codes.Internal, "error while retrieving user: %v", err)
	}
	if err = util.CheckPassword(req.GetPassword(), user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	accessToken, err := server.Token.GenerateJWT(user, time.Hour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating token: %v", err)
	}
	refreshToken, err := server.Token.GenerateJWT(user, time.Hour*5)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating token: %v", err)
	}
	rsp := &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return rsp, nil
}
