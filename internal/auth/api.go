package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/krissukoco/deall-technical-test-go/internal/api"
	"github.com/krissukoco/deall-technical-test-go/internal/user"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{
		service: service,
	}
}

func (ctl *controller) RegisterHandlers(group *gin.RouterGroup) error {
	group.POST("/login", ctl.login)
	group.POST("/register", ctl.register)

	return nil
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

func (ctl *controller) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, api.InvalidJson())
		return
	}
	token, err := ctl.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(400, api.InvalidCredentials())
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func (ctl *controller) register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := ctl.service.Register(req.Email, req.Password, req.Name, req.Gender, req.Birthdate)
	if err != nil {
		if errors.Is(err, user.ErrEmailAlreadyExists) {
			c.JSON(400, api.NewError(api.CodeEmailAlreadyExists, err.Error()))
			return
		}
		if errors.Is(err, ErrPasswordMinLen) {
			c.JSON(400, api.NewError(api.CodePasswordInvalid, err.Error()))
			return
		}
		if errors.Is(err, user.ErrNameMinLen) || errors.Is(err, user.ErrGenderInvalid) || errors.Is(err, user.ErrBirthdateInvalid) {
			c.JSON(400, api.NewError(api.CodeUserData, err.Error()))
			return
		}
		c.JSON(400, api.Unknown(err.Error()))
	}
	c.JSON(200, api.Success())
}
