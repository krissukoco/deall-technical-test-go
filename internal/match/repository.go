package match

import (
	"errors"
	"time"

	"github.com/krissukoco/deall-technical-test-go/internal/models"
	"gorm.io/gorm"
)

var (
	ErrMatchNotFound = errors.New("match not found")
)

type Repository interface {
	Get(id int64) (*models.Match, error)
	GetLast24Hours(userId string, limit int) ([]*models.Match, error)
	Create(userId, matcheeId string) (*models.Match, error)
	Like(id int64) (*models.Match, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&models.Match{})

	return &repository{
		db: db,
	}
}

func (r *repository) Get(id int64) (*models.Match, error) {
	var m models.Match
	err := r.db.Where("id = ?", id).Limit(1).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMatchNotFound
		}
		return nil, err
	}
	return &m, nil
}

func (r *repository) GetLast24Hours(userId string, limit int) ([]*models.Match, error) {
	matches := make([]*models.Match, 0)
	offset24Hours := time.Now().Add(-24 * time.Hour).UnixMilli()
	err := r.db.Where("user_id = ? AND created_at > ?", userId, offset24Hours).
		Limit(limit).
		Find(&matches).
		Error
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (r *repository) Create(userId, matcheeId string) (*models.Match, error) {
	match := &models.Match{
		UserId:    userId,
		MatcheeId: matcheeId,
		Liked:     false,
	}

	err := r.db.Create(match).Error
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (r *repository) Like(id int64) (*models.Match, error) {
	var match models.Match
	err := r.db.Where("id = ?", id).Limit(1).First(&match).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMatchNotFound
		}
		return nil, err
	}

	match.Liked = true
	err = r.db.Save(&match).Error
	if err != nil {
		return nil, err
	}

	return &match, nil
}
