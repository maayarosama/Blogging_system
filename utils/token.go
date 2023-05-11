package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	// "github.com/maayarosama/Blogging_system/models"
)

// func GenerateToken(user_id string, email string, JWTKey string, timeout int) {
// 	expirationTime := time.Now().Add(time.Duration(timeout) * time.Minute)
// 	claims := &models.Claims{
// 		UserID: user_id,
// 		Email:  email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			// In JWT, the expiry time is expressed as unix milliseconds
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// }

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix() // expiration
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey)) // signing token with secret in conf file

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}
