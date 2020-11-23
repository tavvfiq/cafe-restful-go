package apiutils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJwt generate jwt from secret key and payload
func GenerateJwt(secretKey []byte, payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	return tokenString, err
}
