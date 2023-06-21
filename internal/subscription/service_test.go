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
	repo := newMockRepository(getMockItems())
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

type mockRepository struct {
	items []*models.Subscription
}

func newMockRepository(items []*models.Subscription) Repository {
	return &mockRepository{
		items: items,
	}
}

func (m *mockRepository) Get(userId string) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.UserId == userId {
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) GetById(id int64) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) Renew(userId string, add int64) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.UserId == userId {
			item.EndAt += add
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) Create(userId string, add int64) (*models.Subscription, error) {
	start := tests.Now()
	end := start + add
	item := &models.Subscription{
		UserId:  userId,
		StartAt: start,
		EndAt:   end,
	}
	m.items = append(m.items, item)
	return item, nil
}
