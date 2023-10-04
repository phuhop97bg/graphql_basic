package schema

import (
	"backend-food/internal/pkg/repository"
	"backend-food/pkg/infrastucture/db"
)

func GetContainerRepo(data db.Database) map[string]interface{} {

	return map[string]interface{}{
		"user_repository":    repository.NewUserRepository(data),
		"song_repository":    repository.NewSongRepository(data),
		"comment_repository": repository.NewCommentRepository(data),
		//"food_post_repository": repository.NewFoodPostRepository(data),
	}
}
