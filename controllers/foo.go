package controllers

import (
	"net/http"

	"github.com/ttnsp/go-boilerplate/models"
	"github.com/ttnsp/go-boilerplate/repositories"

	"github.com/gin-gonic/gin"
)

func FindFoo(c *gin.Context) {
	var foo models.Foo
	if err := repositories.DB.Model(&models.Foo{}).Where("id = ?", c.Param("id")).First(&foo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foo})
}

func FindFoos(c *gin.Context) {
	var foo []models.Foo
	repositories.DB.Model(&models.Foo{}).Find(&foo)

	c.JSON(http.StatusOK, gin.H{"data": foo})
}

func CreateFoo(c *gin.Context) {
	// Validate input
	var input models.Foo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create foo
	repositories.DB.Create(&input)

	c.JSON(http.StatusCreated, input)
}

func UpdateFoo(c *gin.Context) {
	// Validate input
	var input models.Foo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var foo models.Foo
	if err := repositories.DB.Model(&models.Foo{}).Where("id = ?", c.Param("id")).First(&foo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := repositories.DB.Model(&foo).Updates(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": foo})
}
