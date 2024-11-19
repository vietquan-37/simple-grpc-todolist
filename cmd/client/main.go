package main

import (
	"context"
	"log"

	"github.com/vietquan-37/todo-list/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE5NDg1NjAsImlzcyI6MTczMTk0NDk2MCwicm9sZSI6IlVTRVIiLCJzdWIiOjZ9.IEVmpJVmexfJ1DN3IWzomU83FQ6CCVQNCu6txg0qgCs"

	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect to grpc: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoListClient(conn)

	md := metadata.New(map[string]string{
		"authorization": "Bearer " + token,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	updateTask, err := client.UpdateTask(ctx, &pb.UpdateTaskRequest{
		TaskId:      2,
		TaskName:    "Learn Grpc client",
		Description: "I will learn to write grpc client now",
		Status:      pb.Status(99),
	})
	if err != nil {
		log.Fatalf("error while calling UpdateTask: %v", err)
	}

	log.Printf("Task updated successfully: %+v", updateTask)
}
