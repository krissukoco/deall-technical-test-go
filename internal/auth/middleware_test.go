package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krissukoco/deall-technical-test-go/tests"
)

func TestMiddleware_UserId(t *testing.T) {
	secretKey := "secret"
	userId := "some-user-id"

	mw := Middleware(secretKey)

	now := time.Now()
	nowTs := now.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"iat": nowTs,
		"exp": now.Add(time.Hour * 24).Unix(),
		"nbf": nowTs,
	})
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		t.Error(err)
		return
	}

	c := tests.NewGinContext()
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

	mw(c)
	if c.IsAborted() {
		t.Errorf("Context is aborted with status %d", c.Writer.Status())
		return
	}

	userIdFromContext, ok := c.Get("userId")
	if !ok {
		t.Error("userId not found in context")
		return
	}

	if userIdFromContext != userId {
		t.Error("userId is not equal to expected")
		return
	}

	t.Logf("Middleware_UserId: userIdFromContext: %s", userIdFromContext)
}
