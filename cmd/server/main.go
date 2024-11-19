package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vietquan-37/todo-list/internal/db"
	"github.com/vietquan-37/todo-list/middleware"
	"github.com/vietquan-37/todo-list/util"

	"github.com/vietquan-37/todo-list/internal/model"

	"github.com/vietquan-37/todo-list/pb" // Import your protobuf package
	"github.com/vietquan-37/todo-list/pkg/v1/handler"
	"github.com/vietquan-37/todo-list/pkg/v1/repository"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	config, err := util.LoadConfig("../../")

	if err != nil {
		log.Fatalf("cannot load from configuration: %v", err)
	}
	db := db.DbConn(config.DBSource)

	migrations(db)

	lis, err := net.Listen("tcp", config.GRPCAddress)
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER: %v", err)
	}
	jwtMaker, err := util.NewService(config.SignatureSercret)
	userRepo := initUserRepo(db)
	taskRepo := initTaskRepo(db)
	server := handler.NewServer(userRepo, taskRepo, *jwtMaker)

	if err != nil {
		log.Fatalf("Failed to create JwtMaker: %v", err)
	}

	authInterceptor, err := middleware.NewAuthInterceptor(jwtMaker)
	if err != nil {
		log.Fatalf("Failed to create AuthInterceptor: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.UnaryAuthMiddleware))
	pb.RegisterTodoListServer(grpcServer, server)
	// reflection.Register(grpcServer)
	fmt.Println("Server is listening on port 5051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func initUserRepo(db *gorm.DB) interfaces.UserRepo {
	return repository.NewUserRepo(db)
}
func initTaskRepo(db *gorm.DB) interfaces.TaskRepo {
	return repository.NewTaskRepo(db)
}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Task{}, &model.Session{}, &model.VerifyEmail{})
	if err != nil {
		fmt.Println("Migration error:", err)
	} else {
		fmt.Println("Database migration completed.")
	}
}
