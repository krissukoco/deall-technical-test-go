package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/config"
	"github.com/krissukoco/deall-technical-test-go/internal/auth"
	"github.com/krissukoco/deall-technical-test-go/internal/database"
	"github.com/krissukoco/deall-technical-test-go/internal/match"
	"github.com/krissukoco/deall-technical-test-go/internal/subscription"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
)

const (
	DefaultPort = 8080
)

func getPort(cfgPort int) int {
	if cfgPort == 0 {
		return DefaultPort
	}
	return cfgPort
}

func main() {
	dbConfig := config.Config.Database

	db, err := database.NewPostgresGorm(
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName,
		dbConfig.Timezone,
		dbConfig.Port,
		dbConfig.EnableSsl,
	)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/api/v1")

	// Middlewares
	authMiddleware := auth.Middleware(config.Config.JwtSecret)

	// Services
	userService := user.NewService(user.NewRepository(db))
	authService := auth.NewService(config.Config.JwtSecret, userService)
	subsService := subscription.NewService(subscription.NewRepository(db))

	{
		authCtl := auth.NewController(authService)
		authCtl.RegisterHandlers(v1.Group("/auth"))
	}
	{
		userCtl := user.NewController(userService)
		userCtl.RegisterHandlers(v1.Group("/users"), authMiddleware)
	}
	{
		subsCtl := subscription.NewController(subsService)
		subsCtl.RegisterHandlers(v1.Group("/subscriptions"), authMiddleware)
	}
	{
		matchCtl := match.NewController(match.NewService(match.NewRepository(db), subsService, userService))
		matchCtl.RegisterHandlers(v1.Group("/matches"), authMiddleware)
	}

	port := getPort(config.Config.Port)

	router.Run(fmt.Sprintf(":%d", port))
}
