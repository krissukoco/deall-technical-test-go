package user

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
)

func TestApi(t *testing.T) {
	t.Log("TestApi")
	router := gin.Default()
	now := tests.Now()

	userId := "user1"
	ctl := NewController(NewMockService(NewMockRepository([]*models.User{
		{"user1", "user1@test.com", "password", "User 1", "male", "1995-01-01", "", now, now},
	})))

	secret := "secret"
	header := tests.MockHeaderWithAuth(secret, userId)

	mw := tests.MockAuthMiddlewareWithUserId(userId)
	ctl.RegisterHandlers(router.Group("/users"), mw)

	cases := []*tests.ApiTestCase{
		{"Get Me", "/users/me", "GET", "", header, 200, `*"id":"user1"*`},
	}

	for _, c := range cases {
		tests.API(t, router, c)
	}
}
