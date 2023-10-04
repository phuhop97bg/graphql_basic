package mutation

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func DeleteSongMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.DeleteSongOutput(),
		Description: "delete a song",

		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["song"] == nil {
				err = errors.New("song is required")
				return
			}
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["song"].(map[string]interface{})
			songRepo := containerRepo["song_repository"].(service.SongRepository)
			err = songRepo.DeleteSong(entity.Song{
				ID:     req["id"].(int),
				UserID: user.ID,
			})
			if err != nil {
				return
			}
			result = map[string]interface{}{
				"id": req["id"].(int),
			}
			return
		},
	}
}
