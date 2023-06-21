package subscription

import (
	"errors"
	"testing"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestService_Crud(t *testing.T) {
	t.Log("TestService_Crud")
	repo := NewMockRepository(getMockItems())
	service := NewService(repo)
	assert.NotNil(t, service)

	// Get user sub
	sub, err := service.Get("user1")
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, "user1", sub.UserId)
	assert.Equal(t, false, sub.IsActive()) // Expired

	// Get non existent user sub
	sub, err = service.Get("user4")
	assert.NotNil(t, err)
	assert.Nil(t, sub)
	assert.Equal(t, true, errors.Is(err, ErrNoSubscription))

	// Buy sub
	sub, err = service.Buy("user4", "monthly")
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, "user4", sub.UserId)
	assert.Equal(t, true, sub.IsActive())

	// Buy sub with wrong plan
	sub, err = service.Buy("user4", "wrong")
	assert.NotNil(t, err)
	assert.Nil(t, sub)
	assert.Equal(t, true, errors.Is(err, ErrSubscriptionTypeInvalid))

	// Renew sub
	sub, err = service.Renew("user1", "monthly")
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, "user1", sub.UserId)
	assert.Equal(t, true, sub.IsActive())

	// Renew sub with wrong plan
	sub, err = service.Renew("user4", "wrong")
	assert.NotNil(t, err)
	assert.Nil(t, sub)
	assert.Equal(t, true, errors.Is(err, ErrSubscriptionTypeInvalid))
}

func getMockItems() []*models.Subscription {
	return []*models.Subscription{
		{1, "user1", tests.Now(), tests.Now(), tests.Now(), tests.Now()},
		{2, "user2", tests.Now(), tests.Now(), tests.Now(), tests.Now()},
		{3, "user3", tests.Now(), tests.Now(), tests.Now(), tests.Now()},
	}
}
