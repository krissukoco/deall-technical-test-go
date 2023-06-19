package models

type User struct {
	Id             string `gorm:"primaryKey"`
	Email          string `gorm:"unique"`
	Name           string
	Gender         string
	Birthdate      string
	ProfilePicture string
	Premium        bool
	CreatedAt      int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt      int64 `gorm:"autoUpdateTime:milli"`
}
