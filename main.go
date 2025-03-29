package main

import (
	"shopping_cart/database"
	"shopping_cart/routes"
)

func init() {
	database.ConnectDB()
}

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")

	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
}