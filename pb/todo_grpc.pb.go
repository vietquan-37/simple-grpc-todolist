// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: todo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TodoList_CreateUser_FullMethodName     = "/pb.TodoList/CreateUser"
	TodoList_DeleteUser_FullMethodName     = "/pb.TodoList/DeleteUser"
	TodoList_GetAllUser_FullMethodName     = "/pb.TodoList/GetAllUser"
	TodoList_UpdateUser_FullMethodName     = "/pb.TodoList/UpdateUser"
	TodoList_Login_FullMethodName          = "/pb.TodoList/Login"
	TodoList_CreateTask_FullMethodName     = "/pb.TodoList/CreateTask"
	TodoList_GetAllUserTask_FullMethodName = "/pb.TodoList/GetAllUserTask"
	TodoList_DeleteTask_FullMethodName     = "/pb.TodoList/DeleteTask"
	TodoList_UpdateTask_FullMethodName     = "/pb.TodoList/UpdateTask"
	TodoList_Logout_FullMethodName         = "/pb.TodoList/Logout"
	TodoList_RefreshToken_FullMethodName   = "/pb.TodoList/RefreshToken"
)

// TodoListClient is the client API for TodoList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoListClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetAllUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	GetAllUserTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	RefreshToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
}

type todoListClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoListClient(cc grpc.ClientConnInterface) TodoListClient {
	return &todoListClient{cc}
}

func (c *todoListClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, TodoList_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TodoList_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) GetAllUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, TodoList_GetAllUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, TodoList_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, TodoList_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, TodoList_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) GetAllUserTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, TodoList_GetAllUserTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TodoList_DeleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, TodoList_UpdateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) Logout(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TodoList_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoListClient) RefreshToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, TodoList_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoListServer is the server API for TodoList service.
// All implementations should embed UnimplementedTodoListServer
// for forward compatibility.
type TodoListServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*CommonResponse, error)
	GetAllUser(context.Context, *GetUserRequest) (*UserListResponse, error)
	UpdateUser(context.Context, *UserUpdateRequest) (*UserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error)
	GetAllUserTask(context.Context, *GetTaskRequest) (*TaskListResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*CommonResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error)
	Logout(context.Context, *TokenRequest) (*CommonResponse, error)
	RefreshToken(context.Context, *TokenRequest) (*RefreshTokenResponse, error)
}

// UnimplementedTodoListServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTodoListServer struct{}

func (UnimplementedTodoListServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTodoListServer) DeleteUser(context.Context, *DeleteUserRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedTodoListServer) GetAllUser(context.Context, *GetUserRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}
func (UnimplementedTodoListServer) UpdateUser(context.Context, *UserUpdateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedTodoListServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedTodoListServer) CreateTask(context.Context, *CreateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTodoListServer) GetAllUserTask(context.Context, *GetTaskRequest) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserTask not implemented")
}
func (UnimplementedTodoListServer) DeleteTask(context.Context, *DeleteTaskRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTodoListServer) UpdateTask(context.Context, *UpdateTaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTodoListServer) Logout(context.Context, *TokenRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedTodoListServer) RefreshToken(context.Context, *TokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedTodoListServer) testEmbeddedByValue() {}

// UnsafeTodoListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoListServer will
// result in compilation errors.
type UnsafeTodoListServer interface {
	mustEmbedUnimplementedTodoListServer()
}

func RegisterTodoListServer(s grpc.ServiceRegistrar, srv TodoListServer) {
	// If the following call pancis, it indicates UnimplementedTodoListServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TodoList_ServiceDesc, srv)
}

func _TodoList_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_GetAllUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).GetAllUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).UpdateUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_GetAllUserTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).GetAllUserTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_GetAllUserTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).GetAllUserTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).Logout(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoList_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoListServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoList_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoListServer).RefreshToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoList_ServiceDesc is the grpc.ServiceDesc for TodoList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoList",
	HandlerType: (*TodoListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _TodoList_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _TodoList_DeleteUser_Handler,
		},
		{
			MethodName: "GetAllUser",
			Handler:    _TodoList_GetAllUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _TodoList_UpdateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _TodoList_Login_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _TodoList_CreateTask_Handler,
		},
		{
			MethodName: "GetAllUserTask",
			Handler:    _TodoList_GetAllUserTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TodoList_DeleteTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TodoList_UpdateTask_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _TodoList_Logout_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _TodoList_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
