// go get github.com/golang-jwt/jwt/v5
package models

import "github.com/golang-jwt/jwt/v5"

// Claims structure
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
