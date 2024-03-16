package main

import (
	"Book-GOAPI/database"
	"Book-GOAPI/routes"
)

func main() {
	err := database.DatabaseConnection()
	if err != nil {
		panic("Failed to connect to the database")
	}

	r := routes.SetupRouter()

	// Run the server
	r.Run(":8080")
}
