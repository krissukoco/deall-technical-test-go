package match

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/internal/subscription"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func GetMockUsers(count int) []*models.User {
	users := make([]*models.User, 0)
	now := tests.Now()
	for i := 0; i < count; i++ {
		u := &models.User{
			Id:             fmt.Sprintf("user%d", i),
			Email:          fmt.Sprintf("user%d@test.com", i),
			Password:       "password",
			Name:           fmt.Sprintf("User %d", i),
			Gender:         []string{"male", "female"}[rand.Intn(2)],
			Birthdate:      "1995-12-12",
			ProfilePicture: "",
			CreatedAt:      now,
			UpdatedAt:      now,
		}
		users = append(users, u)
	}
	return users
}

func getUserGender(userId string, users []*models.User) string {
	for _, u := range users {
		if u.Id == userId {
			return u.Gender
		}
	}
	return ""
}

func TestService_CRUD(t *testing.T) {
	repo := NewMockRepository([]*models.Match{})

	now := tests.Now()
	oneMonth := time.Now().Add(30 * 24 * time.Hour).UnixMilli()

	subService := subscription.NewService(subscription.NewMockRepository([]*models.Subscription{
		{1, "user1", now, oneMonth, now, now},
	}))

	users := GetMockUsers(100)
	userService := user.NewService(user.NewMockRepository(users))

	s := NewService(repo, subService, userService)

	// Utils -> get user gender
	matcheeGender := "male"
	gender := getUserGender("user1", users)
	if gender == "male" {
		matcheeGender = "female"
	}
	// Get count of matchees
	count := 0
	for _, u := range users {
		if u.Gender == matcheeGender {
			count++
		}
	}

	// Get a match
	match, err := s.GenerateMatch("user1")
	assert.Nil(t, err)
	assert.NotNil(t, match)

	// Like match
	m, err := s.Like("user1", match.Id)
	assert.Nil(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, "user1", m.UserId)
	assert.Equal(t, match.UserId, m.MatcheeId)
	assert.Equal(t, true, m.Liked)

	// Like match again
	m, err = s.Like("user1", 1)
	assert.NotNil(t, err)
	assert.Nil(t, m)
	assert.Equal(t, true, errors.Is(err, ErrMatchAlreadyLiked))

	// Like non existent match
	m, err = s.Like("user1", 2)
	assert.NotNil(t, err)
	assert.Nil(t, m)
	assert.Equal(t, true, errors.Is(err, ErrMatchNotFound))

	// Get matches if user not subscribed
	for i := 0; i < 10; i++ {
		match, err = s.GenerateMatch("user2")
		assert.Nil(t, err)
		assert.NotNil(t, match)
	}
	match, err = s.GenerateMatch("user2")
	assert.NotNil(t, err)
	assert.Nil(t, match)
	assert.Equal(t, true, errors.Is(err, ErrMaxMatchPerDay))

	// Get matches if user subscribed
	for i := 0; i < count-1; i++ {
		match, err = s.GenerateMatch("user1")
		if ok := assert.Nil(t, err); !ok {
			break
		}
		if ok := assert.NotNil(t, match); !ok {
			break
		}
	}

	// No more matches
	match, err = s.GenerateMatch("user1")
	assert.NotNil(t, err)
	assert.Nil(t, match)
	assert.Equal(t, true, errors.Is(err, ErrNoMatchAvailable))

}
