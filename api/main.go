package main

import (
	"Book-GOAPI/database"
	"Book-GOAPI/model"
	"Book-GOAPI/routes"
	"log"

	"github.com/gin-gonic/gin"
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
	r.Use(CORSMiddleware())

	// Run the server
	r.Run(":8080")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
