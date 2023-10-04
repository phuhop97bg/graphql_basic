package dto

// CreateCommentResponse struct

type CreateCommentResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	SongID  int    `json:"song_id"`
}
