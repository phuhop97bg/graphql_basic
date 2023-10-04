package output

import "github.com/graphql-go/graphql"

func CreateCommentOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CreateCommentOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"song_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
