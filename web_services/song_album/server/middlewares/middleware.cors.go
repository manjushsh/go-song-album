package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ApiUrl = "https://go-song-album.onrender.com"

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := map[string]bool{
		ApiUrl: true,
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Println(c.Request.Header.Get("Origin"))
		if allowed, exists := allowedOrigins[origin]; exists && allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
