package routes

import (
	"Book-GOAPI/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GDSC Book API!",
		})
	})
	// Routes
	books := r.Group("/books")
	{
		books.GET("", controller.GetBooks)
		books.GET("/:id", controller.GetBook)
		books.POST("", controller.CreateBook)
		books.PUT("/:id", controller.UpdateBook)
		books.PATCH("/:id", controller.PatchBook)
		books.DELETE("/:id", controller.DeleteBook)
	}

	return r
}
