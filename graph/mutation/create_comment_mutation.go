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

func CreateCommentMutation(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.CreateCommentOutput(),
		Description: "create a comment",

		Args: graphql.FieldConfigArgument{
			"comment": &graphql.ArgumentConfig{
				Type: input.CommentInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["comment"] == nil {
				err = errors.New("comment is required")
				return
			}
			ctx := p.Context.(*gin.Context)
			user := middleware.GetUserFromContext(ctx)
			req := p.Args["comment"].(map[string]interface{})
			createCommentReq := dto.CreateCommentRequest{}
			utils.ConvertMapToObject(req, &createCommentReq)
			err = utils.CheckValidate(createCommentReq)
			if err != nil {
				return
			}
			cmtRepo := containerRepo["comment_repository"].(service.CommentRepository)
			cmt, err := cmtRepo.CreateComment(entity.Comment{
				Content: createCommentReq.Content,
				SongID:  createCommentReq.SongID,
				UserID:  user.ID,
			})
			if err != nil {
				return
			}
			result = dto.CreateCommentResponse{
				ID:      cmt.ID,
				Content: cmt.Content,
				SongID:  cmt.SongID,
			}
			return
		},
	}
}
