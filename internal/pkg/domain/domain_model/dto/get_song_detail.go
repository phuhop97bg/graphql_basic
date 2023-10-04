package dto

// GetSongDetailResponse struct
type GetSongDetailResponse struct {
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	ContentURL *string      `json:"content_url"`
	View       int          `json:"view"`
	ImageURL   *string      `json:"image_url"`
	Decription string       `json:"decription"`
	CreatedAt  string       `json:"created_at"`
	UserSong   UserResponse `json:"user"`
	Singer     string       `json:"singer"`
}
type UserResponse struct {
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarURL *string `json:"avatar_url"`
}
