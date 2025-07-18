package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	DATABASE_URL := LoadEnv("DATABASE_URL")

	var err error

	DB, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database")
}
