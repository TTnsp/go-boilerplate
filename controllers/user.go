package controllers

import (
	"net/http"

	"github.com/ttnsp/go-boilerplate/auth"
	"github.com/ttnsp/go-boilerplate/models"
	"github.com/ttnsp/go-boilerplate/repositories"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	// Validate input
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users
	if err := repositories.DB.Model(&models.Users{}).Where("name = ?", input.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !auth.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response LoginResponse = LoginResponse{Token: token}
	c.JSON(http.StatusOK, gin.H{"data": response})
}
