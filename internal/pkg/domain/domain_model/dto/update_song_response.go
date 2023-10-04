package dto

// UpdateSongResponse struct
type UpdateSongResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ContentURL  *string `json:"content_url"`
	ImageURL    *string `json:"image_url"`
	Description string  `json:"decription"`
	UpdatedAt   string  `json:"updated_at"`
	Singer      string  `json:"singer"`
}
