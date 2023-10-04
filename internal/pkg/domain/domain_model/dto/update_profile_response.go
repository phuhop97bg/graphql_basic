package dto

// UpdateUserProfileResponse struct
type UpdateUserProfileResponse struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarURL *string `json:"avatar_url"`
}
