package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "postgresql://abrahamsamudra:HKxomRx5AcnOr6Y9lHTIzQ@smart-dormitory-7433.8nk.cockroachlabs.cloud:26257/smart-dorm?sslmode=verify-full"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database.")
	}

	fmt.Println("Connection Opened to Database")
}
