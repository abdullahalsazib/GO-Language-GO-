package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"unique;not null"`
	Age   int
}

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=root dbname=otherdb port=5432 sslmode=disable  TimeZone=Asia/Dhaka"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL!")
}

func main() {
	fmt.Println("postgress setup")
	ConnectDb()
	if DB == nil {
		fmt.Println("❌ Database connection failed!")
		return
	}

	// Auto Migrate Tables
	// config.DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(User{})
	fmt.Println("✅ Database migrated successfully!")
}
