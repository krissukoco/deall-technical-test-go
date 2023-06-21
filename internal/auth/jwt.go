package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userId string, expHours int, secret string) (string, error) {
	now := time.Now()
	nowTs := now.Unix()

	// Create JWT token with standard claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"aud": "deall-technical-test-go",
		"iss": "deall-technical-test-go",
		"iat": nowTs,
		"nbf": nowTs,
		"exp": now.Add(time.Hour * time.Duration(expHours)).Unix(),
		"jti": "jwt_" + uuid.New().String(),
	})

	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string, secret string) (*jwt.MapClaims, error) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return &claims, nil
}

func MockToken(secret, userId string) string {
	token, _ := GenerateToken(userId, 24, secret)
	return token
}
