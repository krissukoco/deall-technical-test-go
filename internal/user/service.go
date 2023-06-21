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
	FindByGenderExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error)
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

func ValidateCreate(email, name, hashedPassword, gender, birthdate string) error {
	// Validate email
	if !utils.ValidEmail(email) {
		return ErrEmailInvalid
	}

	if len(name) < 3 {
		return ErrNameMinLen
	}

	// Validate birthdate. Should be in format YYYY-MM-DD
	t, err := time.Parse("2006-01-02", birthdate)
	if err != nil {
		return ErrBirthdateInvalid
	}
	// Check if birthdate is in the past
	if t.After(time.Now()) {
		return ErrBirthdateInvalid
	}
	// Check minimum age
	if time.Since(t).Hours()/24/365 < MinimumAge {
		return ErrMinimumAge
	}
	// Gender can only be 'male' or 'female'
	if gender != "male" && gender != "female" {
		return ErrGenderInvalid
	}
	return nil
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

func (s *service) FindByGenderExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error) {
	return s.repo.FindByGenderAndExcludeIds(gender, limit, excludeIds)
}

func (s *service) Create(email, name, hashedPassword, gender, birthdate string) (*models.User, error) {
	// Check if email already exists
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	err = ValidateCreate(email, name, hashedPassword, gender, birthdate)
	if err != nil {
		return nil, err
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

type mockService struct {
	repo Repository
}

func NewMockService(repo Repository) Service {
	return &mockService{
		repo: repo,
	}
}

func (m *mockService) GetById(id string) (*models.User, error) {
	return m.repo.GetById(id)
}

func (m *mockService) GetByEmail(email string) (*models.User, error) {
	return m.repo.GetByEmail(email)
}

func (m *mockService) Create(email, name, hashedPassword, gender, birthdate string) (*models.User, error) {
	_, err := m.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	err = ValidateCreate(email, name, hashedPassword, gender, birthdate)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Email:     email,
		Name:      name,
		Password:  hashedPassword,
		Gender:    gender,
		Birthdate: birthdate,
	}
	return m.repo.Create(user)
}

func (m *mockService) FindByGender(gender string, limit int) ([]*models.User, error) {
	return m.repo.FindByGender(gender, limit)
}

func (m *mockService) FindByGenderExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error) {
	return m.repo.FindByGenderAndExcludeIds(gender, limit, excludeIds)
}
