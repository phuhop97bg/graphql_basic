package output

import "github.com/graphql-go/graphql"

func GetSongDetailOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetSongDetailOutput",
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
				"view": &graphql.Field{
					Type: graphql.Int,
				},
				"image_url": &graphql.Field{
					Type: graphql.String,
				},
				"decription": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: userInfoOutput(),
				},
				"singer": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
func userInfoOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "userInfoOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"username": &graphql.Field{
					Type: graphql.String,
				},
				"first_name": &graphql.Field{
					Type: graphql.String,
				},
				"last_name": &graphql.Field{
					Type: graphql.String,
				},
				"avatar_url": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
