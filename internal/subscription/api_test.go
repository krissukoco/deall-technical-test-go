package subscription

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
	"github.com/krissukoco/deall-technical-test-go/internal/auth"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
)

func TestApi(t *testing.T) {
	t.Log("TestApi")
	router := gin.Default()
	userId := "user1"
	ctl := NewController(NewService(newMockRepository([]*models.Subscription{
		{1, userId, tests.Now(), tests.Now(), tests.Now(), tests.Now()},
	})))

	secret := "secret"
	header := auth.MockHeaderWithAuth(secret, userId)

	mw := auth.MockAuthMiddlewareWithUserId(userId)
	ctl.RegisterHandlers(router.Group("/subscriptions"), mw)

	cases := []*tests.ApiTestCase{
		{"Get user's subscription", "/subscriptions", "GET", "", header, 200, `*"id":1*`},
		{"Buy subscription monthly", "/subscriptions/buy/monthly", "POST", "", header, 200, `*"id":1*`},
		{"Cannot buy subscription", "/subscriptions/buy/monthly", "POST", "", header, 400, fmt.Sprintf(`*"code":%d`, api.CodeAlreadySubscribed)},
		{"Renew subscription yearly", "/subscriptions/renew/yearly", "POST", "", header, 200, `*"id":1*`},
	}

	for _, c := range cases {
		tests.API(t, router, c)
	}
}
