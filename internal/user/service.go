package user

import (
	"errors"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/pkg/utils"
)

const (
	MinimumAge = 12
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrEmailInvalid       = errors.New("email invalid")
	ErrBirthdateInvalid   = errors.New("birthdate invalid")
	ErrGenderInvalid      = errors.New("gender can only be male or female")
	ErrMinimumAge         = errors.New("minimum age is 12 years old")
	ErrNameMinLen         = errors.New("name must be at least 3 characters long")
)

type Service interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	FindByGender(gender string, limit int) ([]*models.User, error)
	Create(email, name, hashedPassword, gender, birthdate string) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetById(id string) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *service) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *service) FindByGender(gender string, limit int) ([]*models.User, error) {
	return s.repo.FindByGender(gender, limit)
}

func (s *service) Create(email, name, hashedPassword, gender, birthdate string) (*models.User, error) {
	// Check if email already exists
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	// Validate email
	if !utils.ValidEmail(email) {
		return nil, ErrEmailInvalid
	}

	if len(name) < 3 {
		return nil, ErrNameMinLen
	}

	// Validate birthdate. Should be in format YYYY-MM-DD
	t, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return nil, ErrBirthdateInvalid
	}
	// Check if birthdate is in the past
	if t.After(time.Now()) {
		return nil, ErrBirthdateInvalid
	}
	// Check minimum age
	if time.Since(t).Hours()/24/365 < MinimumAge {
		return nil, ErrMinimumAge
	}
	// Gender can only be 'male' or 'female'
	if gender != "male" && gender != "female" {
		return nil, ErrGenderInvalid
	}

	user := &models.User{
		Email:     email,
		Name:      name,
		Gender:    gender,
		Birthdate: birthdate,
		Password:  hashedPassword,
		// Premium:   false,
	}
	return s.repo.Create(user)
}
