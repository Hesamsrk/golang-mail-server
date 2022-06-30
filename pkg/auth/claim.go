package auth

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

