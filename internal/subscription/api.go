package subscription

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
	"github.com/krissukoco/deall-technical-test-go/internal/models"
)

var (
	_ = models.Subscription{}
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
	group.GET("/packages", ctl.GetPackages)

	// Private routes
	group.Use(authMiddleware)
	group.GET("", ctl.Get)
	group.POST("/buy/:type", ctl.Buy)
	group.POST("/renew/:type", ctl.Renew)
}

// Get godoc
// @Summary Get user's subscription
// @Schemes
// @Security AccessToken
// @Description Get user's current subscription. Returns 404 and code if user doesn't have any
// @Tags Subscription
// @Produce json
// @Success 200 {object} models.Subscription
// @Router /subscriptions [get]
func (ctl *controller) Get(c *gin.Context) {
	userId := c.GetString("userId")
	subscription, err := ctl.service.Get(userId)
	if err != nil {
		if errors.Is(err, ErrNoSubscription) {
			c.JSON(404, api.NewError(api.CodeNoSubscription, "You don't have any subscription"))
			return
		}
		c.JSON(500, api.Internal())
		return
	}

	c.JSON(200, subscription)
}

// GetPackages godoc
// @Summary Get all subscription packages
// @Schemes
// @Tags Subscription
// @Produce json
// @Success 200 {array} subscriptionPackage
// @Router /subscriptions/packages [get]
func (ctl *controller) GetPackages(c *gin.Context) {
	packages := ctl.service.Packages()
	c.JSON(200, packages)
}

// Buy godoc
// @Summary Buy a subscription by type
// @Schemes
// @Security AccessToken
// @Tags Subscription
// @Param type path string true "Subscription type"
// @Produce json
// @Success 200 {array} models.Subscription
// @Router /subscriptions/buy/{type} [post]
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

// Renew godoc
// @Summary Renew user's subscription
// @Schemes
// @Security AccessToken
// @Tags Subscription
// @Param type path string true "Subscription type"
// @Produce json
// @Success 200 {object} models.Subscription
// @Router /subscriptions/renew/{type} [post]
func (ctl *controller) Renew(c *gin.Context) {
	sub, err := ctl.service.Renew(c.GetString("userId"), c.Param("type"))
	if err != nil {
		if errors.Is(err, ErrSubscriptionTypeInvalid) {
			c.JSON(400, api.NewError(api.CodeSubscriptionTypeInvalid, "Invalid subscription type"))
			return
		}
		c.JSON(400, api.NewError(api.CodeUnknown, err.Error()))
		return
	}
	c.JSON(200, sub)
}
