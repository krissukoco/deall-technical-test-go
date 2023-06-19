package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/auth"
	"github.com/krissukoco/deall-technical-test-go/internal/database"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
)

const (
	DefaultPort = 8080
)

func getPort() int {
	portStr, exists := os.LookupEnv("PORT")
	if !exists {
		return DefaultPort
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return DefaultPort
	}
	return port
}

type Server struct {
}

func main() {
	db, err := database.NewDefaultPostgresGorm()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/api/v1")

	jwtsecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		panic("JWT_SECRET env var is not set")
	}

	// Services
	userService := user.NewService(user.NewRepository(db))
	authService := auth.NewService(jwtsecret, userService)
	authMiddleware := auth.Middleware(jwtsecret)

	{
		authCtl := auth.NewController(authService)
		authCtl.RegisterHandlers(v1.Group("/auth"))
	}
	{
		userCtl := user.NewController(userService)
		userCtl.RegisterHandlers(v1.Group("/users"), authMiddleware)
	}

	port := getPort()

	router.Run(fmt.Sprintf(":%d", port))
}
