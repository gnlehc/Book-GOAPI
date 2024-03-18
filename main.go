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

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if there is nothing in the environment variable
	}
	r.Run(":" + port)
}
