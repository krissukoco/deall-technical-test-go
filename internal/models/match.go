package models

type Match struct {
	Id        int64 `gorm:"primaryKey"`
	UserId    string
	MatcheeId string
	Liked     bool
	CreatedAt string `gorm:"autoCreateTime:milli"`
}
