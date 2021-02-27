package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GetJWTFromUser(userID string, tokenID string, expiresAt time.Time, key string) (string, time.Time) {
	claims := Claims{
		ID:     tokenID,
		UserID: userID,
	}

	claims.ExpiresAt = expiresAt.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(key))

	return tokenString, expiresAt
}
