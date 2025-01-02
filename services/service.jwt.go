package services

import (
	"go-song-album/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Define the JWT secret key
var jwtSecret = []byte("123456-secret-key")

// GenerateJWT generates a new JWT token
func GenerateJWT(username string) (string, error) {
	claims := models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT token
func ValidateJWT(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims
	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
