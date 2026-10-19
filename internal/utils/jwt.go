
package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID string) (string, error) {
	secretKey := os.Getenv("ACCESS_TOKEN_SECRET")
	if secretKey == "" {
		return "", jwt.ErrInvalidKey
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(userID string) (string, error) {
	secretKey := os.Getenv("REFRESH_TOKEN_SECRET")
	if secretKey == "" {
		return "", jwt.ErrInvalidKey
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
