package dto

import "time"

// LoginResponse struct
type LoginResponse struct {
	Token          string    `json:"token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
	Role           string    `json:"role"`
}
