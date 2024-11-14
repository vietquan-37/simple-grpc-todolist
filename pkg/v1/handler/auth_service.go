package handler

import (
	"context"
	"errors"
	"time"

	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/val"
	"github.com/vietquan-37/todo-list/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	violation := validateCreateUserRequest(req)
	if violation != nil {
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

func validateCreateUserRequest(req *pb.CreateUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violation = append(violation, ErrorResponse("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violation = append(violation, ErrorResponse("password", err))
	}
	if err := val.ValidatePhoneNumber(req.GetPhoneNumber()); err != nil {
		violation = append(violation, ErrorResponse("phone_number", err))
	}
	if err := val.ValidateFullname(req.GetFullName()); err != nil {
		violation = append(violation, ErrorResponse("full_name", err))
	}

	return violation
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
