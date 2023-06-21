package auth

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
	"github.com/krissukoco/deall-technical-test-go/pkg/utils"
	"github.com/krissukoco/deall-technical-test-go/tests"
)

func TestApi(t *testing.T) {
	t.Log("TestApi")
	secret := "secret"
	router := gin.Default()
	now := tests.Now()

	pwd := "password"
	h, err := utils.HashPassword(pwd)
	if err != nil {
		t.Errorf("Error hashing password: %s", err.Error())
		return
	}

	// userId := "user1"
	userRepo := user.NewMockRepository([]*models.User{
		{"user1", "user1@test.com", h, "User 1", "male", "1995-01-01", "", now, now},
	})
	ctl := NewController(NewMockService(user.NewMockService(userRepo), secret))

	// header := tests.MockHeaderWithAuth(secret, userId)

	ctl.RegisterHandlers(router.Group("/auth"))

	cases := []*tests.ApiTestCase{
		{"Login", "/auth/login", "POST", `{"email":"user1@test.com", "password":"password"}`, nil, 200, `*"token"*`},
		{"Login with invalid email", "/auth/login", "POST", `{"email":"user2@test.com", "password":"password"}`, nil, 400, fmt.Sprintf(`*"code":%d*`, api.CodeInvalidCredentials)},
		{"Login with invalid password", "/auth/login", "POST", `{"email":"user1@test.com", "password":"somethingElse"}`, nil, 400, fmt.Sprintf(`*"code":%d*`, api.CodeInvalidCredentials)},
		{
			"Register",
			"/auth/register",
			"POST",
			`{"email":"user2@test.com", "password":"password", "name":"user 2", "gender":"male", "birthdate":"1995-01-01"}`,
			nil,
			200,
			`*"success":true*`,
		},
		{
			"Register with invalid email",
			"/auth/register",
			"POST",
			`{"email":"1234235123", "password":"password", "name":"user 2", "gender":"male", "birthdate":"1995-01-01"}`,
			nil,
			400,
			fmt.Sprintf(`*"message":"%s"*`, "email invalid"),
		},
		{
			"Register with existing email",
			"/auth/register",
			"POST",
			`{"email":"user1@test.com", "password":"password", "name":"user 2", "gender":"male", "birthdate":"1995-01-01"}`,
			nil,
			400,
			fmt.Sprintf(`*"code":%d*`, api.CodeEmailAlreadyExists),
		},
	}

	for _, c := range cases {
		tests.API(t, router, c)
	}
}
