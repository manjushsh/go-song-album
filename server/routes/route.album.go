package routes

import (
	"go-song-album/controllers"
	"go-song-album/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(r *gin.Engine) {
	albumGroup := r.Group("/albums")
	albumGroup.Use(middlewares.AuthMiddleware())
	albumGroup.GET("/", controllers.GetAlbums)
	albumGroup.GET("/:id", controllers.GetAlbumByID)
	albumGroup.POST("/", controllers.PostAlbums)
	albumGroup.PUT("/:id", controllers.UpdateAlbum)
	albumGroup.DELETE("/:id", controllers.DeleteAlbum)
}
