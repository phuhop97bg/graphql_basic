package schema

import (
	"backend-food/graph/mutation"

	"github.com/graphql-go/graphql"
)

func GetAnonymousMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"create_user": mutation.CreateUserMutation(containerRepo),
		},
	})
}

func GetClientMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"update_user_profile": mutation.UpdateUserProfileMutation(containerRepo),
			"change_password":     mutation.ChangePasswordMutation(containerRepo),
			"create_song":         mutation.CreateSongMutation(containerRepo),
			"delete_song":         mutation.DeleteSongMutation(containerRepo),
			"update_song":         mutation.UpdateSongMutation(containerRepo),
			"create_comment":      mutation.CreateCommentMutation(containerRepo),
		},
	})
}
func GetAdminMutation(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "mutation",
		Fields: graphql.Fields{},
	})
}
