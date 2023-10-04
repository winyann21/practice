package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	dsn := os.Getenv("DB")
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully connected to the database")
    }
}