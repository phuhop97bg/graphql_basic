package input

import "github.com/graphql-go/graphql"

func CommentInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "Comment Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"content": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"song_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
