package user

import "github.com/krissukoco/deall-technical-test-go/internal/models"

type Service interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(email, name, password, gender, birthdate string) (*models.User, error)
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

func (s *service) Create(email, name, password, gender, birthdate string) (*models.User, error) {
	panic("implement me")
}
