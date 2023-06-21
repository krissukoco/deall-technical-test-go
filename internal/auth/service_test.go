package auth

import (
	"errors"
	"testing"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
	"github.com/krissukoco/deall-technical-test-go/pkg/utils"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestService_CRUD(t *testing.T) {
	secret := "secret"
	pwd := "password"
	h, err := utils.HashPassword(pwd)
	if err != nil {
		t.Errorf("Error hashing password: %s", err.Error())
		return
	}
	now := tests.Now()

	userRepo := user.NewMockRepository([]*models.User{
		{"1234567", "user1@test.com", h, "User 1", "male", "1995-12-12", "", now, now},
	})
	userService := user.NewMockService(userRepo)
	s := NewMockService(userService, secret)

	// Login user not found
	_, err = s.Login("nothing@nothing.com", "1234")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrCredentialsInvalid))

	// Login user with wrong password
	_, err = s.Login("user1@test.com", "1234")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrCredentialsInvalid))

	// Login user
	token, err := s.Login("user1@test.com", pwd)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
	// t.Logf("Token: %s", token)

	// Verify token
	claims, err := ParseToken(token, secret)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	sub, ok := (*claims)["sub"].(string)
	assert.Equal(t, true, ok)
	assert.Equal(t, "1234567", sub)

	// Register user
	err = s.Register("user2@test.com", pwd, "User 2", "male", "1995-12-12")
	assert.Nil(t, err)
	// Login user
	token, err = s.Login("user2@test.com", pwd)
	assert.Nil(t, err)
	assert.NotEqual(t, "", token)
	// t.Logf("Token: %s", token)

	// Register with same email
	err = s.Register("user2@test.com", pwd, "User 2", "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, user.ErrEmailAlreadyExists))

}
