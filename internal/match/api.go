package match

import (
	"errors"
	"strconv"

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
	group.GET("/new", authMiddleware, ctl.newMatch)
	group.POST("/like/:id", authMiddleware, ctl.likeMatch)
	// group.GET("/:id", ctl.GetById)
}

func (ctl *controller) newMatch(c *gin.Context) {
	data, err := ctl.service.GenerateMatch(c.GetString("userId"))
	if err != nil {
		if errors.Is(err, ErrMaxMatchPerDay) {
			c.JSON(402, gin.H{
				"message": "You have reached your max match per day",
			})
			return
		}
		if errors.Is(err, ErrNoMatchAvailable) {
			c.JSON(404, api.NewError(api.CodeNoMatchAvailable, "No match available"))
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, data)
}

func (ctl *controller) likeMatch(c *gin.Context) {
	userId := c.GetString("userId")
	matchIdParam := c.Param("id")
	matchId, err := strconv.Atoi(matchIdParam)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Match not found",
		})
		return
	}
	m, err := ctl.service.Like(userId, int64(matchId))
	if err != nil {
		if errors.Is(err, ErrMatchNotFound) {
			c.JSON(404, api.NewError(api.CodeMatchNotFound, "Match not found"))
			return
		}
		if errors.Is(err, ErrMatchAlreadyLiked) {
			c.JSON(400, api.NewError(api.CodeMatchAlreadyLiked, "You already liked this match"))
			return
		}
		c.JSON(500, api.Internal())
		return
	}
	c.JSON(200, m)
}
