package controllers

import (
	"go-song-album/models"
	"go-song-album/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context) {
	var loginRequest models.Auth
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !services.IsValidUsername(loginRequest.Username) || !services.IsValidPassword(loginRequest.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	var foundUser models.Auth
	err := mongoInstance.FindOne("users", bson.M{"username": loginRequest.Username, "isdeleted": false}).Decode(&foundUser)
	if err != nil {
		services.HandleError(c, err, http.StatusInternalServerError, "Invalid credentials - No User Found")
		return
	}

	if !services.CheckPasswordHash(loginRequest.Password, foundUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var registerRequest models.Auth
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !services.IsValidUsername(registerRequest.Username) || !services.IsValidPassword(registerRequest.Password) {
		services.HandleError(c, nil, http.StatusBadRequest, "Invalid credentials")
		return
	}

	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	var existingUser models.Auth
	err := mongoInstance.FindOne("users", bson.M{"username": registerRequest.Username}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	}

	hashedPassword, err := services.HashPassword(registerRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	registerRequest.Password = hashedPassword
	registerRequest.Project = "go-song-album"
	registerRequest.IsDeleted = false

	_, err = mongoInstance.Insert("users", registerRequest)
	if err != nil {
		services.HandleError(c, err, http.StatusInternalServerError, "Failed to register user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func GetUsers(c *gin.Context) {
	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	var users []models.Auth
	cursor, err := mongoInstance.FindAll("users", bson.M{})
	if err != nil {
		services.HandleError(c, err, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	if err = cursor.All(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	username := c.Param("username")

	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	var user models.Auth
	err := mongoInstance.FindOne("users", bson.M{"username": username}).Decode(&user)
	if err != nil {
		services.HandleError(c, err, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var updatedUser models.Auth
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !services.IsValidUsername(updatedUser.Username) || !services.IsValidPassword(updatedUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	result, err := mongoInstance.Update("users", bson.M{"username": username}, bson.M{"$set": updatedUser})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")

	if !services.IsValidUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	mongoInstance, ok := services.ConnectToMongo(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	result, err := mongoInstance.Delete("users", bson.M{"username": username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
