package match

import (
	"fmt"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
	"github.com/krissukoco/deall-technical-test-go/internal/auth"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/internal/subscription"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
	"github.com/krissukoco/deall-technical-test-go/tests"
)

func TestApi(t *testing.T) {
	t.Log("TestApi")
	router := gin.Default()

	repo := NewMockRepository([]*models.Match{})

	now := tests.Now()
	oneMonth := time.Now().Add(30 * 24 * time.Hour).UnixMilli()

	subService := subscription.NewService(subscription.NewMockRepository([]*models.Subscription{
		{1, "user1", now, oneMonth, now, now},
	}))

	users := GetMockUsers(100)
	userService := user.NewService(user.NewMockRepository(users))

	userId := "user1"
	secret := "secret"
	mw := auth.MockAuthMiddlewareWithUserId(userId)
	s := NewService(repo, subService, userService)

	ctl := NewController(s)
	ctl.RegisterHandlers(router.Group("/matches"), mw)

	header := auth.MockHeaderWithAuth(secret, userId)

	cases := []*tests.ApiTestCase{
		{"Get new match", "/matches/new", "GET", "", header, 200, `*"id":1*`},
		{"Like match", "/matches/like/1", "POST", "", header, 200, `*"liked":true*`},
		{"Like match again", "/matches/like/1", "POST", "", header, 400, fmt.Sprintf(`*"code":%d*`, api.CodeMatchAlreadyLiked)},
		{"Like non existing match", "/matches/like/2", "POST", "", header, 404, fmt.Sprintf(`*"code":%d*`, api.CodeMatchNotFound)},
	}

	for _, c := range cases {
		tests.API(t, router, c)
	}
}
