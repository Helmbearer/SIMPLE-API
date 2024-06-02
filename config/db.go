package config

import (
	"SIMPLE-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})

	DB = database
}
