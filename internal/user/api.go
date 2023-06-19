package user

import "github.com/gin-gonic/gin"

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{
		service: service,
	}
}

func (ctl *controller) RegisterHandlers(group *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	group.GET("/:id", ctl.GetById)
}

func (ctl *controller) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := ctl.service.GetById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}
