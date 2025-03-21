package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("text.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	//create db

	db.Create(&Product{
		Code:  "D42",
		Price: 2001,
	})

	// Read
	var product Product
	db.First(&product, 1)
	// db.First(*&product, "Code = ?", "D42")
}
