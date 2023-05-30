package initalizers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable
var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed Connecting to Database")
	}
}
