package schema

import (
	"backend-food/graph/query"

	"github.com/graphql-go/graphql"
)

func GetAnonymousQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"login": query.LoginQuery(containerRepo),
		},
	})
}

func GetClientQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"get_user_profile": query.GetProfileUserQuery(containerRepo),
			"get_song_list":    query.GetSongListQuery(containerRepo),
			"get_song_detail":  query.GetSongDetailQuery(containerRepo),
			"get_comment_song": query.GetCommentSongQuery(containerRepo),
		},
	})
}
func GetAdminQuery(containerRepo map[string]interface{}) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "query",
		Fields: graphql.Fields{},
	})
}
