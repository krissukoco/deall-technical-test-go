package subscription

import (
	"testing"

	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	// Get database
	db, err := tests.NewTestDb()
	if err != nil {
		t.Errorf("Error creating test db: %s", err.Error())
		return
	}
	repo := NewRepository(db)
	assert.NotNil(t, repo)

	userId := "user1"
	addMilliseconds := int64(100)

	// Create subscription
	sub, err := repo.Create(userId, addMilliseconds)
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, userId, sub.UserId)
	assert.Equal(t, true, sub.IsActive())

	// Wait to make sure the subscription is expired
	tests.Sleep(1)
	assert.Equal(t, false, sub.IsActive())

	// Get subscription
	sub, err = repo.Get(userId)
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, userId, sub.UserId)
	assert.Equal(t, false, sub.IsActive())

	// Get by id
	subId := sub.Id
	sub, err = repo.GetById(subId)
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, userId, sub.UserId)
	assert.Equal(t, sub.Id, subId)

	// Renew
	renewAdd := int64(100000)
	sub, err = repo.Renew(userId, renewAdd)
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	assert.Equal(t, subId, sub.Id)
	assert.Equal(t, true, sub.IsActive())

}
