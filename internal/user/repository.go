package user

import (
	"log"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	FindByGender(gender string, limit int) ([]*models.User, error)
	FindByGenderAndExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error)
	Create(user *models.User) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

func newUserId() string {
	return ulid.Make().String()
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

func (r *repository) FindByGender(gender string, limit int) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := r.db.Where("gender = ?", gender).
		Limit(limit).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindByGenderAndExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error) {
	users := make([]*models.User, 0)
	err := r.db.Where("gender = ?", gender).
		Not(excludeIds).
		Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	log.Println()
	return users, nil
}

func (r *repository) Create(user *models.User) (*models.User, error) {
	user.Id = newUserId()
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
