package apicontrollers

import (
	"net/http"
	"sisbus/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexStudi(c *gin.Context) {
	var studi []models.Studi

	models.DB.Find(&studi)
	c.JSON(http.StatusOK, gin.H{"studi": studi})
}

func ShowStudi(c *gin.Context) {
	var studi models.Studi
	id := c.Param("id")

	if err := models.DB.First(&studi, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak Ditemukan"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"studi": studi})

}

func CreateStudi(c *gin.Context) {
			//sesuaikan dengan field pada db dan foldeer models

	var request struct {
		S01 string `form:"S01" binding:"required"`
		S02 string `form:"S02" binding:"required"`
		S03 string `form:"S03" binding:"required"`
		S04 string `form:"S04" binding:"required"`
	}

	// Bind request body to the struct
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	// Ensure userID is of type uint
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is not of type uint"})
		return
	}

	// Create a new bekerja record
	newStudi := models.Studi{ 
		//sesuaikan dengan field pada db dan foldeer models
		S01:    request.S01,
		S02:    request.S02,
		S03:    request.S03,
		S04:    request.S04,
		IdUser: userIDUint, // Assign the user ID
	}

	// Create new entry in the database
	if err := models.DB.Create(&newStudi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entry"})
		return
	}

	// Return a JSON response with the created entry data
	c.JSON(http.StatusOK, gin.H{"studi": newStudi})
}

func UpdateStudi(c *gin.Context) {
	var studi models.Studi
	id := c.Param("id")

	if err := c.ShouldBindJSON(&studi); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&studi).Where("id = ?", id).Updates(&studi).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak Dapat MengUpdate Data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil diUpdate"})

}

func DeleteStudi(c *gin.Context) {
	Id := c.Param("id")
	//var notes models.Note

	// Parse the ID into an integer
	id, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Delete the user by ID
	if err := models.DB.Delete(&models.Studi{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
