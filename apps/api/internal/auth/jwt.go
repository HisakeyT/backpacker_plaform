package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
	}
}

func (j *JWTManager) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id":     userID,
		"exprires_at": time.Now().Add(time.Hour * 24).Unix(),
		"issued_at":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}
