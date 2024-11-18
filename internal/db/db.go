package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn(DbSource string) *gorm.DB {

	db, err := gorm.Open(
		postgres.Open(DbSource), &gorm.Config{TranslateError: true},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
