package user

import (
	"errors"
	"log"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

var (
	ErrNoUser = errors.New("no user")
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoUser
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.Where("email = ?", email).Limit(1).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoUser
		}
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

type mockRepository struct {
	items []*models.User
}

func NewMockRepository(items []*models.User) Repository {
	return &mockRepository{items}
}

func (m *mockRepository) GetById(id string) (*models.User, error) {
	for _, item := range m.items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, ErrNoUser
}

func (r *mockRepository) GetByEmail(email string) (*models.User, error) {
	for _, item := range r.items {
		if item.Email == email {
			return item, nil
		}
	}
	return nil, ErrNoUser
}

func (r *mockRepository) FindByGender(gender string, limit int) ([]*models.User, error) {
	items := make([]*models.User, 0)
	for _, item := range r.items {
		if item.Gender == gender {
			items = append(items, item)
		}
	}
	return items, nil
}

func (r *mockRepository) FindByGenderAndExcludeIds(gender string, limit int, excludeIds []string) ([]*models.User, error) {
	items, _ := r.FindByGender(gender, limit)
	for _, excludeId := range excludeIds {
		for i, item := range items {
			if item.Id == excludeId {
				items = append(items[:i], items[i+1:]...)
			}
		}
	}
	return items, nil
}

func (m *mockRepository) Create(u *models.User) (*models.User, error) {
	u.Id = newUserId()
	u.CreatedAt = time.Now().UnixMilli()
	u.UpdatedAt = time.Now().UnixMilli()
	m.items = append(m.items, u)
	return u, nil
}
