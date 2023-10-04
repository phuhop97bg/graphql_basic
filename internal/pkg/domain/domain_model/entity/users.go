package entity

import (
	"time"
)

type userRole string

const (
	AdminRole  userRole = "admin"
	ClientRole userRole = "client"
)

type activeStatusUser string

const (
	NO_ACTIVE_STATUS = "no_active"
	ACTIVE_STATUS    = "active"
	BANNED_STATUS    = "banned"
)

// Users struct
type Users struct {
	ID             int              `gorm:"column:id;primary_key;auto_increment;not null"`
	FirstName      string           `gorm:"column:first_name;"`
	LastName       string           `gorm:"column:last_name"`
	Username       string           `gorm:"column:username;not null"`
	Email          string           `gorm:"column:email;not null"`
	Password       string           `gorm:"column:password;not null"`
	Role           userRole         `gorm:"column:role"`
	ActiveStatus   activeStatusUser `gorm:"column:active_status"`
	AvatarUrl      *string          `gorm:"column:avatar_url"`
	Token          *string          `gorm:"column:token"`
	TokenExpiredAt *time.Time       `gorm:"column:token_expired_at"`
	BaseModel
}

func (u *Users) TableName() string {
	return "users"
}
