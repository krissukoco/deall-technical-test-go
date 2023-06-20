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

	// Swagger docs
	docs "github.com/krissukoco/deall-technical-test-go/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title           Dating App REST API
// @version         1.0
// @description     REST API for Dating App - Deall Technical Test - Written in Go

// @contact.name   Kris Sukoco
// @contact.url    https://github.com/krissukoco
// @contact.email  kristianto.sukoco@gmail.com

// @BasePath  /api/v1

// @securityDefinitions.apiKey AccessToken
// @in header
// @name Authorization

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

	// Swagger Documentation
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middlewares
	authMiddleware := auth.Middleware(config.Config.JwtSecret)

	// Services
	userService := user.NewService(user.NewRepository(db))
	authService := auth.NewService(config.Config.JwtSecret, userService)
	subsService := subscription.NewService(subscription.NewRepository(db))

	{
		authGroup := v1.Group("/auth")
		authCtl := auth.NewController(authService)
		authCtl.RegisterHandlers(authGroup)
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
