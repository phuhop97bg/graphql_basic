package mutation

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func CreateSongMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateSongOutput(),
		Description: "create a song",

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
			createSongReq := dto.CreateSongRequest{}
			utils.ConvertMapToObject(req, &createSongReq)
			err = utils.CheckValidate(createSongReq)
			if err != nil {
				return
			}
			songRepo := containerRepo["song_repository"].(service.SongRepository)
			song := entity.Song{
				Title:      createSongReq.Title,
				UserID:     user.ID,
				Decription: createSongReq.Description,
				Singer:     createSongReq.Singer,
			}
			imgFile, _ := ctx.FormFile("image_file")
			if imgFile != nil {
				ioFile, errFile := imgFile.Open()
				if errFile != nil {
					err = errFile
					return
				}
				url, errUpload := utils.UploadFile(ioFile, imgFile.Filename, []string{"img_music"})
				if errUpload != nil {
					err = errUpload
					return
				}
				song.ImageURL = &url
			}

			contentFile, _ := ctx.FormFile("music_file")
			if contentFile == nil {
				err = errors.New("music File is required")
				return
			}
			ioFile, errFile := contentFile.Open()
			if errFile != nil {
				err = errFile
				return
			}
			url, errUpload := utils.UploadFile(ioFile, utils.RemoveEndFile(contentFile.Filename), []string{"music"})
			if errUpload != nil {
				err = errUpload
				return
			}
			song.ContentURL = &url
			song, err = songRepo.CreateSong(song)
			if err != nil {
				return
			}
			result = dto.CreateSongResponse{
				ID:          song.ID,
				Title:       song.Title,
				ContentURL:  song.ContentURL,
				ImageURL:    song.ImageURL,
				Description: song.Decription,
				Singer:      song.Singer,
				CreatedAt:   utils.FormatTime(song.CreatedAt),
			}
			return
		},
	}
}
