package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	model "github.com/liju-github/FoodBuddyMicroserviceUser/models"
)

// CustomClaims extends jwt.StandardClaims
type CustomClaims struct {
	UserID     string `json:"userid"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Reputation int32  `json:"reputation"`
	jwt.RegisteredClaims
}

// JWT related constants and variables
var (
	TokenExpiry  = 24 * time.Hour
	JWTSecretKey string
)

// SetJWTSecretKey sets the JWT secret key
func SetJWTSecretKey(secret string) {
	JWTSecretKey = secret
}

// ValidateToken verifies the JWT token
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, model.ErrInvalidToken
	}

	return claims, nil
}
