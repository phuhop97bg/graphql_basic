package dto

// GetSongListResponse struct
type GetSongListResponse struct {
	SongList []SongResponse `json:"songs"`
}
type SongResponse struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	ContentURL *string `json:"content_url"`
	View       int     `json:"view"`
	ImageURL   *string `json:"image_url"`
	Decription string  `json:"decription"`
	CreatedAt  string  `json:"created_at"`
	Singer     string  `json:"singer"`
}
