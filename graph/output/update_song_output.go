package output

import "github.com/graphql-go/graphql"

func UpdateSongOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UpdateSongOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content_url": &graphql.Field{
					Type: graphql.String,
				},
				"image_url": &graphql.Field{
					Type: graphql.String,
				},
				"decription": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
				"singer": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
