package models

// Connect user interests, to be used for matching
type Interest struct {
	Id         int64 `gorm:"primaryKey"`
	UserId     string
	InterestId string
}
