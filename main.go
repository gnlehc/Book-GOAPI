package main

import (
	"Book-GOAPI/database"
	"Book-GOAPI/model"
	"Book-GOAPI/routes"
	"log"
	"os"
)

func main() {
	err := database.DatabaseConnection()
	if err != nil {
		panic("Failed to connect to the database")
	}
	db := database.GlobalDB

	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
