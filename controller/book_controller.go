package controller

import (
	"Book-GOAPI/database"
	"Book-GOAPI/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all books
func GetBooks(c *gin.Context) {
	var books []model.Book
	database.GlobalDB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Get single book
func GetBook(c *gin.Context) {
	var book model.Book
	if err := database.GlobalDB.Where("book_id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// Create new book
func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.StoreBookRecord()
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(500, gin.H{
	// 		"StatusCode": "500",
	// 		"Message":    "Create book error, Book Already Exists",
	// 	})
	// 	c.Abort()
	// 	return
	// }
	c.JSON(http.StatusCreated, book)
}

// Update book
func UpdateBook(c *gin.Context) {
	var book model.Book
	if err := database.GlobalDB.Where("book_id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.GlobalDB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Patch book
func PatchBook(c *gin.Context) {
	var book model.Book
	if err := database.GlobalDB.Where("book_id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var updatedBook model.Book
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Only update fields that are provided in the request payload
	database.GlobalDB.Model(&book).Updates(updatedBook)

	c.JSON(http.StatusOK, book)
}

// Delete book
func DeleteBook(c *gin.Context) {
	var book model.Book
	if err := database.GlobalDB.Where("book_id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	database.GlobalDB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
