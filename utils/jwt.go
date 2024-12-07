package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

func GenerateJWT(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}
