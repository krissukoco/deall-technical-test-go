package auth

import (
	"testing"
)

const (
	JwtSecret = "secret"
)

func TestGenerateAndParseToken(t *testing.T) {
	userId := "123"
	expHours := 24

	token, err := GenerateToken(userId, expHours, JwtSecret)
	if err != nil {
		t.Error(err)
	}

	claims, err := ParseToken(token, JwtSecret)
	if err != nil {
		t.Error(err)
		return
	}

	if (*claims)["sub"] != userId {
		t.Error("userId is not equal")
	}
}
