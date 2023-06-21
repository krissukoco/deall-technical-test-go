package user

import (
	"errors"
	"testing"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestRepository_ReadWrite(t *testing.T) {
	db, err := tests.NewTestDb("test")
	if err != nil {
		t.Errorf("Error creating test db: %s", err.Error())
		return
	}
	repo := NewRepository(db)
	if repo == nil {
		t.Errorf("Error creating repository")
		return
	}

	// Get user not found
	_, err = repo.GetById("user1")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrNoUser))

	// Get user by email not found
	_, err = repo.GetByEmail("user1@test.com")
	assert.NotNil(t, err)
	assert.Equal(t, true, errors.Is(err, ErrNoUser))

	// Create user
	user, err := repo.Create(&models.User{
		Email:     "user1@test.com",
		Password:  "password",
		Name:      "User 1",
		Birthdate: "2006-01-02",
		Gender:    "male",
	})
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEqual(t, "", user.Id)
	assert.Equal(t, "user1@test.com", user.Email)
	assert.Equal(t, "User 1", user.Name)
	assert.Equal(t, "2006-01-02", user.Birthdate)
	assert.Equal(t, "male", user.Gender)

	// Get user
	userId := user.Id
	user, err = repo.GetById(userId)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userId, user.Id)

	// Get user by email
	user, err = repo.GetByEmail("user1@test.com")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userId, user.Id)
	assert.Equal(t, "user1@test.com", user.Email)

	// Get by gender
	users, err := repo.FindByGender("male", 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))

	// Get by gender exclude ids
	users, err = repo.FindByGenderAndExcludeIds("male", 10, []string{userId})
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 0, len(users))

	// Get by gender exclude ids but not the user's id
	users, err = repo.FindByGenderAndExcludeIds("male", 10, []string{"user2"})
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 1, len(users))

}
