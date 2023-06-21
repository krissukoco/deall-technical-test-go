package match

import (
	"testing"

	"github.com/krissukoco/deall-technical-test-go/tests"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	db, err := tests.NewTestDb("test")
	if err != nil {
		t.Errorf("Error creating test db: %s", err.Error())
		return
	}
	repo := NewRepository(db)
	assert.NotNil(t, repo)

	userId := "user1"

	// Get last 24 hours matches
	matches, err := repo.GetLast24Hours(userId, 10)
	assert.Nil(t, err)
	assert.NotNil(t, matches)

	// Create match
	matcheeId := "matchee1"
	match, err := repo.Create(userId, matcheeId)
	assert.Nil(t, err)
	assert.NotNil(t, match)
	assert.Equal(t, userId, match.UserId)
	assert.Equal(t, matcheeId, match.MatcheeId)

	// Get match
	match, err = repo.Get(match.Id)
	assert.Nil(t, err)
	assert.NotNil(t, match)
	assert.Equal(t, userId, match.UserId)
	assert.Equal(t, matcheeId, match.MatcheeId)

	// Get last 24 hours matches
	matches, err = repo.GetLast24Hours(userId, 10)
	assert.Nil(t, err)
	assert.NotNil(t, matches)
	assert.Equal(t, 1, len(matches))

	// Like match
	match, err = repo.Like(match.Id)
	assert.Nil(t, err)
	assert.NotNil(t, match)
	assert.Equal(t, userId, match.UserId)
	assert.Equal(t, matcheeId, match.MatcheeId)
	assert.Equal(t, true, match.Liked)

}
