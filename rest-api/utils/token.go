package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate token using jwt package
const SecretKey = "supersecretkeywillbesavehere"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})
	return token.SignedString([]byte(SecretKey))

}
