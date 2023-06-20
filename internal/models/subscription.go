package models

import (
	"time"
)

type Subscription struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    string
	StartAt   int64
	EndAt     int64
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func (s *Subscription) Renew(add int64) {
	now := time.Now().UnixMilli()
	if s.EndAt < now {
		s.StartAt = now
		s.EndAt = now
	}
	s.EndAt += add
}

func (s *Subscription) IsActive() bool {
	return s.EndAt > s.StartAt
}
