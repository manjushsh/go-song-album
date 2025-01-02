package controllers

import (
	"go-song-album/models"
	"go-song-album/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.Auth{}

func Login(c *gin.Context) {
	var loginRequest models.Auth
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if loginRequest.Username != "admin" || loginRequest.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT
	token, err := services.GenerateJWT(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Implement register API
func Register(c *gin.Context) {
	var registerRequest models.Auth
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Implement the registration logic here. store the user in the auth array
	users = append(users, registerRequest)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
