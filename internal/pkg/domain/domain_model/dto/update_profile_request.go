package dto

// UpdateUserProfileResponse struct
type UpdateUserProfileRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
