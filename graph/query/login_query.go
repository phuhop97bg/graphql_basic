package query

import (
	"backend-food/graph/input"
	"backend-food/graph/output"
	"backend-food/internal/pkg/domain/domain_model/dto"
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/internal/pkg/domain/service"
	"backend-food/pkg/share/middleware"
	"backend-food/pkg/share/utils"
	"errors"
	"time"

	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

func LoginQuery(containerRepo map[string]interface{}) *graphql.Field {
	return &graphql.Field{
		Type:        output.LoginOutput(),
		Description: "User Login",
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: input.UserInput(),
			},
		},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			if p.Args["user"] == nil {
				err = errors.New("user is required")
				return
			}
			req := p.Args["user"].(map[string]interface{})
			loginReq := dto.LoginRequest{}
			utils.ConvertMapToObject(req, &loginReq)
			err = utils.CheckValidate(loginReq)
			if err != nil {
				return
			}
			loginReq.Password = utils.EncryptPassword(loginReq.Password)
			userRepo := containerRepo["user_repository"].(service.UserRepository)

			user, err := userRepo.FirstUser(entity.Users{
				Username: loginReq.Username,
				Password: loginReq.Password,
			})
			if err != nil {
				return
			}

			if user.ID == 0 {
				err = errors.New(utils.LOGIN_FAIL_ERROR)
				return
			}
			timeNow := time.Now()

			timeExpiredAt := timeNow.Add(time.Hour * 48)
			// generate uuid
			uuid := uuid.Must(uuid.NewV4(), nil)
			tokenString, err := middleware.GenerateJWTToken(middleware.JWTParam{
				UUID:       uuid,
				Authorized: true,
				ExpriedAt:  timeExpiredAt,
			})

			if err != nil {
				return
			}

			newUser := entity.Users{
				Token:          &tokenString,
				TokenExpiredAt: &timeExpiredAt,
			}
			_, err = userRepo.UpdateUser(newUser, user)

			loginRes := dto.LoginResponse{
				Token:          tokenString,
				TokenExpiredAt: timeExpiredAt,
				Role:           string(user.Role),
			}
			result = loginRes
			return
		},
	}
}
