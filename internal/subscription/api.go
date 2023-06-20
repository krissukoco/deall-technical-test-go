package subscription

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
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
	group.GET("", authMiddleware, ctl.Get)
	group.POST("/buy/:type", authMiddleware, ctl.Buy)
	group.POST("/renew/:type", authMiddleware, ctl.Renew)
}

func (ctl *controller) Get(c *gin.Context) {
	userId := c.GetString("userId")
	subscription, err := ctl.service.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNoSubscription) {
			c.JSON(404, gin.H{
				"message": "You don't have any subscription",
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, subscription)
}

func (ctl *controller) Buy(c *gin.Context) {
	sub, err := ctl.service.Buy(c.GetString("userId"), c.Param("type"))
	if err != nil {
		if errors.Is(err, ErrSubscriptionTypeInvalid) {
			c.JSON(400, api.NewError(api.CodeSubscriptionTypeInvalid, "Invalid subscription type"))
			return
		}
		if errors.Is(err, ErrAlreadySubscribed) {
			c.JSON(400, api.NewError(api.CodeAlreadySubscribed, "You already subscribed to this subscription type"))
			return
		}

		c.JSON(400, api.NewError(api.CodeUnknown, err.Error()))
		return
	}
	c.JSON(200, sub)
}

func (ctl *controller) Renew(c *gin.Context) {
	panic("not implemented")
}
