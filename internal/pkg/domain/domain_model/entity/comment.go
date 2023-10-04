package entity

// Commet struct
type Comment struct {
	ID      int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Content string `gorm:"column:content;"`
	UserID  int    `gorm:"column:user_id;"`
	SongID  int    `gorm:"column:song_id;"`
	BaseModelWithDeleteAt
}

func (u *Comment) TableName() string {
	return "comment"
}
