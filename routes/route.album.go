package routes

import (
	"go-song-album/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	albumGroup := r.Group("/albums")
	albumGroup.GET("/", controllers.GetAlbums)
	albumGroup.GET("/:id", controllers.GetAlbumByID)
	albumGroup.POST("/", controllers.PostAlbums)
	albumGroup.PUT("/:id", controllers.UpdateAlbum)
	albumGroup.DELETE("/:id", controllers.DeleteAlbum)
}
