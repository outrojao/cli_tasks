package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken() (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	sub := os.Getenv("JWT_HEADER_SUB")
	iss := os.Getenv("JWT_HEADER_ISS")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iss": iss,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}
