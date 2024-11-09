package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vietquan-37/todo-list/internal/db"

	"github.com/vietquan-37/todo-list/internal/model"

	"github.com/vietquan-37/todo-list/pb" // Import your protobuf package
	"github.com/vietquan-37/todo-list/pkg/v1/handler"
	"github.com/vietquan-37/todo-list/pkg/v1/repository"
	"github.com/vietquan-37/todo-list/pkg/v1/repository/interfaces"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {

	db := db.DbConn()

	migrations(db)

	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER: %v", err)
	}

	userUseCase := initUserServer(db)
	server := handler.NewServer(userUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterTodoListServer(grpcServer, server)
	reflection.Register(grpcServer)
	fmt.Println("Server is listening on port 5051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func initUserServer(db *gorm.DB) interfaces.UserRepo {
	return repository.NewUserRepo(db)
}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Task{}, &model.Session{}, &model.VerifyEmail{})
	if err != nil {
		fmt.Println("Migration error:", err)
	} else {
		fmt.Println("Database migration completed.")
	}
}
