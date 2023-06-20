package models

type Match struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	UserId    string `json:"-"`
	MatcheeId string `json:"matchee_id"`
	Liked     bool   `json:"liked"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}
