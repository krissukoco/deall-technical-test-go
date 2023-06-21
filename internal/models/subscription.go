package models

import (
	"time"
)

type Subscription struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	UserId    string `json:"-"`
	StartAt   int64  `json:"start_at"`
	EndAt     int64  `json:"end_at"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (s *Subscription) Renew(add int64) {
	now := time.Now().UnixMilli()
	if s.StartAt == 0 {
		s.StartAt = now
	}
	if s.EndAt < now {
		s.EndAt = now
	}
	s.EndAt += add
}

func (s *Subscription) IsActive() bool {
	if s == nil {
		return false
	}
	return s.EndAt > time.Now().UnixMilli()
}
