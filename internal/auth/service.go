package auth

import (
	"errors"
	"log"

	"github.com/krissukoco/deall-technical-test-go/internal/user"
	"golang.org/x/crypto/bcrypt"
)

const (
	DefaultJwtExpirationHours = 24 * 7
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrCredentialsInvalid = errors.New("credentials invalid")
	ErrPasswordMinLen     = errors.New("password must be at least 8 characters long")
)

type Service interface {
	Login(email, password string) (string, error)
	Register(email, password, name, gender, birthdate string) error
}

type service struct {
	jwtSecret       string
	userService     user.Service
	expirationHours int
	saltCost        int
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
		saltCost:        7,
	}
}

func (s *service) comparePassword(password string, hashedPassword string) error {
	log.Println("password:", password, "hashed: ", hashedPassword)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *service) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.saltCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *service) Login(email, password string) (string, error) {
	user, err := s.userService.GetByEmail(email)
	if err != nil {
		return "", ErrCredentialsInvalid
	}
	// Check password
	if err := s.comparePassword(password, user.Password); err != nil {
		log.Println("compare password error:", err)
		return "", ErrCredentialsInvalid
	}
	// Generate JWT
	token, err := GenerateToken(user.Id, s.expirationHours, s.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Register(email, password, name, gender, birthdate string) error {
	// Hash password
	if len(password) < 8 {
		return ErrPasswordMinLen
	}
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		return err
	}

	_, err = s.userService.Create(email, name, hashedPassword, gender, birthdate)
	return err
}
