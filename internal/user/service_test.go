package user

import (
	"errors"
	"testing"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestService(t *testing.T) {
	pwd := "password"
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), 3)
	if err != nil {
		t.Errorf("Error hashing password: %s", err.Error())
		return
	}
	h := string(hashed)
	now := tests.Now()

	repo := NewMockRepository([]*models.User{
		{"user1", "user1@test.com", h, "User 1", "male", "1995-12-12", "", now, now},
		{"user2", "user2@test.com", h, "User 2", "female", "1993-12-12", "", now, now},
		{"user3", "user3@test.com", h, "User 3", "male", "1992-12-12", "", now, now},
	})

	s := NewService(repo)

	// Get user not found
	_, err = s.GetById("user4")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrNoUser))

	// Get user by email not found
	_, err = s.GetByEmail("user4@test.com")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrNoUser))

	// Create user
	user, err := s.Create("user4@test.com", "User 4", h, "male", "1995-12-12")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEqual(t, "", user.Id)

	// Get user
	userId := user.Id
	user, err = s.GetById(userId)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userId, user.Id)

	// Create user with no email
	_, err = s.Create("", "User 5", h, "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrEmailInvalid))

	// Create user with invalid email
	_, err = s.Create("test", "User 5", h, "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrEmailInvalid))

	// Create user with invalid email 2
	_, err = s.Create("test@test", "User 5", h, "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrEmailInvalid))

	// Create user with existing email
	_, err = s.Create("user1@test.com", "User 1", h, "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrEmailAlreadyExists))

	// Create user with no name
	_, err = s.Create("user5@test.com", "", h, "male", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrNameMinLen))

	// Create user with invalid gender
	_, err = s.Create("user5@test.com", "User 5", h, "wrong", "1995-12-12")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrGenderInvalid))

	// Create user with random birthdate
	_, err = s.Create("user5@test.com", "User 5", h, "male", "1234567890")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrBirthdateInvalid))

	// Create user with invalid birthdate
	_, err = s.Create("user5@test.com", "User 5", h, "male", "1995-30-02")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrBirthdateInvalid))

}
