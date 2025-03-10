package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	OneHour = 1
	HalfDay = 12
	OneDay  = 24
	TwoDays = 48
)

type LoginResponse struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Token generation
func GenerateToken(userId, jwtSecret string, expiresHours time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresHours * time.Hour).UTC()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// Token validation
func ValidateToken(tokenString, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token is expired")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
