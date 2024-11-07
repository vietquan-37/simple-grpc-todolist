package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {

	db, err := gorm.Open(
		postgres.Open("postgresql://postgres:12345@localhost:5431/todo_list?sslmode=disable"), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
