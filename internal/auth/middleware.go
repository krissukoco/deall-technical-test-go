package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Malformed auth header",
			})
			return
		}
		if split[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Auth method is unavailable",
			})
			return
		}
		token := split[1]

		claims, err := ParseToken(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid token",
			})
			return
		}
		userId, ok := (*claims)["sub"].(string)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Invalid token",
			})
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
