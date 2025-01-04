package controllers

import (
	"context"
	"go-song-album/models"
	"go-song-album/services"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Generate random UUID
func GenerateUUID() string {
	return strconv.Itoa(rand.Int())
}

func getMongoService(c *gin.Context) *services.MongoService {
	mongoInstance, err := services.NewMongoService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create mongo service"})
		return nil
	}
	return mongoInstance
}

// GetAlbums handles GET requests to retrieve the list of albums.
func GetAlbums(c *gin.Context) {
	mongoInstance := getMongoService(c)
	if mongoInstance == nil {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"isdeleted": bson.M{"$ne": true}}
	cursor, err := mongoInstance.FindAll("albums", filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve albums"})
		return
	}
	defer cursor.Close(context.Background())

	var albums []models.Album
	if err = cursor.All(context.Background(), &albums); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to decode albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	newAlbum.ID = GenerateUUID()
	newAlbum.IsDeleted = false

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	mongoInstance := getMongoService(c)
	if mongoInstance == nil {
		return
	}
	defer mongoInstance.Disconnect()

	_, err := mongoInstance.Insert("albums", newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to add album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id parameter sent by the client.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	mongoInstance := getMongoService(c)
	if mongoInstance == nil {
		return
	}
	defer mongoInstance.Disconnect()

	var album models.Album
	err := mongoInstance.FindOne("albums", bson.M{"id": id}).Decode(&album)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// UpdateAlbum - updates an existing album with the provided ID.
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	mongoInstance := getMongoService(c)
	if mongoInstance == nil {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedAlbum}

	result, err := mongoInstance.Update("albums", filter, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to update album"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// DeleteAlbum - marks an album as deleted with the provided ID.
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	mongoInstance := getMongoService(c)
	if mongoInstance == nil {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"isdeleted": true}}

	result, err := mongoInstance.Update("albums", filter, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete album"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album marked as deleted"})
}
