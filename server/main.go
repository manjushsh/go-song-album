package main

import (
	"go-song-album/middlewares"
	"go-song-album/routes"
	"go-song-album/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	services.InitEnv()
	router := gin.Default()

	// Static route to serve HTML, CSS, and JS from client/public
	router.Use(middlewares.CORSMiddleware())
	router.Static("/home", "../client/public")

	// Health API
	router.GET("/", func(c *gin.Context) {
		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
		c.JSON(200, gin.H{
			"message": "Welcome to the Go-Song-Album API!",
			"homeUrl": scheme + ":" + c.Request.Host + "/home",
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
