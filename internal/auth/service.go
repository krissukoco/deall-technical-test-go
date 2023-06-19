package auth

import (
	"errors"

	"github.com/krissukoco/deall-technical-test-go/internal/user"
)

const (
	DefaultJwtExpirationHours = 24 * 7
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type Service interface {
	Login(email, password string) (string, error)
	Register(email, password, name, gender, birthdate string) error
}

type service struct {
	jwtSecret       string
	userService     user.Service
	expirationHours int
}

func NewService(jwtSecret string, userService user.Service, jwtExpirationHours ...int) Service {
	exp := DefaultJwtExpirationHours
	if len(jwtExpirationHours) > 0 {
		exp = jwtExpirationHours[0]
	}
	return &service{
		jwtSecret:       jwtSecret,
		userService:     userService,
		expirationHours: exp,
	}
}

func (s *service) Login(email, password string) (string, error) {
	user, err := s.userService.GetByEmail(email)
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (s *service) Register(email, password, name, gender, birthdate string) error {
	_, err := s.userService.GetByEmail(email)
	if err == nil {
		return ErrEmailAlreadyExists
	}
	_, err = s.userService.Create(email, password, name, gender, birthdate)
	return err
}
