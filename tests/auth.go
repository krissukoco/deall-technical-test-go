package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MockToken(secret string, userId string, expHours int) string {
	now := Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"aud": "test",
		"iss": "test",
		"iat": now,
		"nbf": now,
		"exp": now + (int64(expHours) * 60 * 60),
	})
	tokenStr, _ := token.SignedString([]byte(secret))
	return tokenStr
}

func MockHeaderWithAuth(secret, userId string) http.Header {
	h := http.Header{}
	h.Set("Authorization", "Bearer "+MockToken(secret, userId, 24))
	return h
}

func MockAuthMiddlewareWithUserId(userId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userId", userId)
		c.Next()
	}
}
