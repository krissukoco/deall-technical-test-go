package user

import (
	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&models.User{})

	return &repository{
		db: db,
	}
}

func (r *repository) GetById(id string) (*models.User, error) {
	var u models.User
	err := r.db.Where("id = ?", id).Limit(1).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.Where("email = ?", email).Limit(1).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
