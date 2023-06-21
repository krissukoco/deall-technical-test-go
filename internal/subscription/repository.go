package subscription

import (
	"errors"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"github.com/krissukoco/deall-technical-test-go/tests"
	"gorm.io/gorm"
)

type Repository interface {
	Get(userId string) (*models.Subscription, error)
	GetById(id int64) (*models.Subscription, error)
	Renew(userId string, add int64) (*models.Subscription, error)
	Create(userId string, add int64) (*models.Subscription, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&models.Subscription{})

	return &repository{
		db: db,
	}
}

func (r *repository) Get(userId string) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.Where("user_id = ?", userId).Limit(1).First(&sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoSubscription
		}
		return nil, err
	}
	return &sub, nil
}

func (r *repository) GetById(id int64) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.Where("id = ?", id).Limit(1).First(&sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoSubscription
		}
		return nil, err
	}
	return &sub, nil
}

func (r *repository) Create(userId string, add int64) (*models.Subscription, error) {
	start := time.Now().UnixMilli()
	sub := models.Subscription{
		UserId:  userId,
		StartAt: start,
		EndAt:   start,
	}
	sub.Renew(add)
	err := r.db.Create(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *repository) Renew(userId string, add int64) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.Where("user_id = ?", userId).Limit(1).First(&sub).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sub, err := r.Create(userId, add)
			if err != nil {
				return nil, err
			}
			sub.Renew(add)
			return sub, nil
		}
		return nil, err
	}
	sub.Renew(add)

	err = r.db.Save(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}

type mockRepository struct {
	items []*models.Subscription
}

func NewMockRepository(items []*models.Subscription) Repository {
	return &mockRepository{
		items: items,
	}
}

func (m *mockRepository) Get(userId string) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.UserId == userId {
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) GetById(id int64) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) Renew(userId string, add int64) (*models.Subscription, error) {
	for _, item := range m.items {
		if item.UserId == userId {
			item.EndAt += add
			return item, nil
		}
	}
	return nil, ErrNoSubscription
}

func (m *mockRepository) Create(userId string, add int64) (*models.Subscription, error) {
	start := tests.Now()
	end := start + add
	item := &models.Subscription{
		UserId:  userId,
		StartAt: start,
		EndAt:   end,
	}
	m.items = append(m.items, item)
	return item, nil
}
