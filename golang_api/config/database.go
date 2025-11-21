package config

import (
	"log"
	"os"

	"task-manager-fullstack/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	var err error

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("Database url not found or empty")
	}
	log.Println("dbUrl: ", dbUrl)

	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil || DB == nil {
		log.Fatal("failed to connect to database: ", err)
	}

	DB.AutoMigrate(&models.Task{})
}
