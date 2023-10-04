package dto

type ChangePasswordRequest struct {
	Password    string `json:"password" validate:"required,password"`
	OldPassword string `json:"old_password" validate:"required,password"`
}
