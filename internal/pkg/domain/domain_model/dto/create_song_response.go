package dto

// CreateSongResponse struct
type CreateSongResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ContentURL  *string `json:"content_url"`
	ImageURL    *string `json:"image_url"`
	Description string  `json:"decription"`
	CreatedAt   string  `json:"created_at"`
	Singer      string  `json:"singer"`
}
