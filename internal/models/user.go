package models

import "time"

type User struct {
	Id             string `gorm:"primaryKey"`
	Email          string `gorm:"unique"`
	Password       string `json:"-"`
	Name           string
	Gender         string
	Birthdate      string
	ProfilePicture string
	CreatedAt      int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64 `gorm:"autoUpdateTime:milli"`
}

func (u *User) GetAge() int {
	// Parse birthdate
	t, err := time.Parse("2006-01-02", u.Birthdate)
	if err != nil {
		return 0
	}
	// Calculate age
	return int(time.Since(t).Hours() / 24 / 365)
}
