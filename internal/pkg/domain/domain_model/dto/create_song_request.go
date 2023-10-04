package dto

// CreateSongRequest struct
type CreateSongRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"decription"`
	Singer      string `json:"singer"`
}
