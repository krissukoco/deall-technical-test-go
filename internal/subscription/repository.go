package subscription

import (
	"errors"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	Get(userId string) (*models.Subscription, error)
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
			return r.Create(userId, add)
		}
		return nil, err
	}

	err = r.db.Save(&sub).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}
