package entity

type Song struct {
	ID         int     `gorm:"column:id;primary_key;auto_increment;not null"`
	Title      string  `gorm:"column:title;"`
	Singer     string  `gorm:"column:singer;"`
	ContentURL *string `gorm:"column:content_url"`
	View       int     `gorm:"column:view"`
	UserID     int     `gorm:"column:user_id;"`
	ImageURL   *string `gorm:"column:image_url"`
	Decription string  `gorm:"column:decription"`
	BaseModelWithDeleteAt
}

func (u *Song) TableName() string {
	return "songs"
}
