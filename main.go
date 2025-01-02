package main

import (
	"go-song-album/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Health API
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go-Song-Album API!",
		})
	})

	routes.RegisterAuthRoutes(router)
	routes.RegisterAlbumRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
