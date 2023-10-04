package dto

// GetUserProfileResponse struct
type GetUserProfileResponse struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Role      string  `json:"role"`
	CreatedAt string  `json:"created_at"`
	AvatarURL *string `json:"avatar_url"`
}
