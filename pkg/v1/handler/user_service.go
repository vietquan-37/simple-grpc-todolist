package handler

import (
	"context"

	"github.com/vietquan-37/todo-list/pb"
	"github.com/vietquan-37/todo-list/pkg/v1/val"
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
	model := convertUser(req)

	check, err := server.Repo.GetUserByEmail(req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error when retrieving user: %s", err)
	}
	if check != nil {
		return nil, status.Errorf(codes.AlreadyExists, "email %s already register before", check.Email)
	}
	user, err := server.Repo.CreateUser(model)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while creating user: %s", err)
	}

	return convertUserResponse(*user), nil
}
func (server *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userId := req.GetId()
	err := server.Repo.DeleteUser(int(userId))
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "error while deleting user: %s", err)
	}
	return &pb.DeleteUserResponse{
		Message: "Delete user successfully",
	}, nil

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
func (server *Server) GetAllUser(ctx context.Context, req *pb.Pagination) (*pb.UserListResponse, error) {
	userList, err := server.Repo.GetAllUser(req.GetFullName(), req.GetPageNumber(), req.GetPageSize())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while retrieving users: %s", err)
	}
	return convertUserListResponse(userList), nil
}
