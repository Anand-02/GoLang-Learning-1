package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=your_password dbname=blog port=5432 sslmode=disable timezone=Asia/Kolkata"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to Database")
	}

	database.AutoMigrate(&BlogItem{})

	DB = database
}
