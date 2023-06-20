package models

import "time"

type User struct {
	Id             string `gorm:"primaryKey" json:"id"`
	Email          string `gorm:"unique" json:"email"`
	Password       string `json:"-"`
	Name           string `json:"name"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	ProfilePicture string `json:"profile_picture"`
	CreatedAt      int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt      int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
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
