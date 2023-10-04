package dto

// CreateFoodPostRequest struct
type CreateCommentRequest struct {
	SongID  int    `json:"song_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}
