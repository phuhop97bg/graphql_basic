package dto

//GetCommentSongResponse struct
type GetCommentSongResponse struct {
	CommentList []CommentResponse `json:"comments"`
}
type CommentResponse struct {
	ID          int          `json:"id"`
	Content     string       `json:"content"`
	UserComment UserResponse `json:"user"`
}
