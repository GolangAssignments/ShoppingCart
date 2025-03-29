package database

import (
	"fmt"
	"log"
	"shopping_cart/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	var err error 
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect Sqlite Database:", err)
	}
	err = DB.AutoMigrate(&models.Item{}, &models.Cart{}, &models.CartItem{})
	if err != nil {
		log.Fatal("Failed to migrate the table:", err)
	}

	fmt.Println("Sqlite database connected successfully!")
}