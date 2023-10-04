package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"errors"

	"github.com/graphql-go/graphql"
)

func GetCommentSongQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetCommetSongOutput(),
		Description: "get comments of song",
		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
			"page": &graphql.ArgumentConfig{
				Type: input.PaginationInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["song"] == nil {
				err = errors.New("song is required")
				return
			}
			if p.Args["page"] == nil {
				err = errors.New("page is required")
				return
			}
			reqPage := p.Args["page"].(map[string]interface{})
			reqSong := p.Args["song"].(map[string]interface{})

			song := entity.Song{
				ID: reqSong["id"].(int),
			}
			cmtRepo := containerRepo["comment_repository"].(service.CommentRepository)
			userRepo := containerRepo["user_repository"].(service.UserRepository)

			comments, err := cmtRepo.FindCommentListWithPagination(entity.Comment{SongID: song.ID}, reqPage["page_num"].(int), reqPage["page_size"].(int))
			if err != nil {
				return
			}
			commentResponseList := make([]dto.CommentResponse, len(comments))
			for i, v := range comments {
				user, err1 := userRepo.FirstUser(entity.Users{ID: v.UserID})
				if err1 != nil {
					err = err1
					return
				}
				commentResponseList[i] = dto.CommentResponse{
					ID:      v.ID,
					Content: v.Content,
					UserComment: dto.UserResponse{
						Username:  user.Username,
						FirstName: user.FirstName,
						LastName:  user.LastName,
						AvatarURL: user.AvatarUrl,
					},
				}
			}
			result = dto.GetCommentSongResponse{
				CommentList: commentResponseList,
			}
			return
		},
	}
}
