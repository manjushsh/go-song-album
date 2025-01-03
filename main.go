package main

import (
	"go-song-album/routes"
	"go-song-album/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	services.InitEnv()
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
		port = "3060"
	}
	router.Run(":" + port)
}
