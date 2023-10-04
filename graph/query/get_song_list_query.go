package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/utils"
	"errors"

	"github.com/graphql-go/graphql"
)

func GetSongListQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.GetSongListOutput(),
		Description: "get song list",
		Args: graphql.FieldConfigArgument{
			"song": &graphql.ArgumentConfig{
				Type: input.SongInput(),
			},
			"page": &graphql.ArgumentConfig{
				Type: input.PaginationInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["page"] == nil {
				err = errors.New("page is required")
				return
			}
			song := entity.Song{}
			if p.Args["song"] != nil {
				reqSong := p.Args["song"].(map[string]interface{})
				if reqSong["title"] != nil {
					song.Title = reqSong["title"].(string)
				}
				if reqSong["singer"] != nil {
					song.Singer = reqSong["singer"].(string)
				}
				if reqSong["user_id"] != nil {
					song.UserID = reqSong["user_id"].(int)
				}
			}
			songRepo := containerRepo["song_repository"].(service.SongRepository)
			reqPage := p.Args["page"].(map[string]interface{})
			songList, err := songRepo.FindSongListWithPagination(song, reqPage["page_num"].(int), reqPage["page_size"].(int))
			if err != nil {
				return
			}
			songListRes := make([]dto.SongResponse, len(songList))
			for i, v := range songList {
				songListRes[i] = dto.SongResponse{
					ID:         v.ID,
					Title:      v.Title,
					ContentURL: v.ContentURL,
					View:       v.View,
					ImageURL:   v.ImageURL,
					Decription: v.Decription,
					CreatedAt:  utils.FormatTime(v.CreatedAt),
					Singer:     v.Singer,
				}
			}
			result = dto.GetSongListResponse{
				SongList: songListRes,
			}
			return
		},
	}
}
