package dto

// UpdateSongRequest struct
type UpdateSongRequest struct {
	SongID      int    `json:"id" validate:"required"`
	Title       string `json:"title"`
	Description string `json:"decription"`
	Singer      string `json:"singer"`
}
