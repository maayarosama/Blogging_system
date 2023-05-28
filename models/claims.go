package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user_id int, email string, jwtKey string, timeout int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(timeout) * time.Minute)
	claims := &Claims{
		UserID: user_id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}
	return tokenString, nil
}

func ValidateToken(token, secret string, timeout int) (Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if !tkn.Valid {
		return Claims{}, fmt.Errorf("token '%s' is invalid", token)
	}

	if time.Until(claims.ExpiresAt.Time) > time.Duration(timeout)*time.Minute {
		return Claims{}, fmt.Errorf("token '%s' is expired", token)
	}

	return *claims, nil
}
