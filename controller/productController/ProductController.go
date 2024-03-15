package productcontroller

import (
	"encoding/json"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func Show(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data is not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := models.DB.Create(&product)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create product", "error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func Update(c *gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	check := models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0

	if check {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot update the product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update product"})

}

func Delete(c *gin.Context) {

	var product models.Product

	// ambil id dari JSON
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// konversi id dari JSON ke int
	id, err := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete product!"})

}
