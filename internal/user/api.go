package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
)

var (
	_ = models.User{}
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{
		service: service,
	}
}

func (ctl *controller) RegisterHandlers(group *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	group.GET("/me", authMiddleware, ctl.GetMe)
	// group.GET("/:id", ctl.GetById)
}

// GetMe godoc
// @Summary Get user account
// @Schemes
// @Security AccessToken
// @Tags User
// @Produce json
// @Success 200 {object} models.User
// @Router /users/me [get]
func (ctl *controller) GetMe(c *gin.Context) {
	userId := c.GetString("userId")
	user, err := ctl.service.GetById(userId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

// func (ctl *controller) GetById(c *gin.Context) {
// 	id := c.Param("id")
// 	user, err := ctl.service.GetById(id)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(200, user)
// }
