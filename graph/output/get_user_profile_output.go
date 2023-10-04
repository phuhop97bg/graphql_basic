package output

import "github.com/graphql-go/graphql"

func GetUserProfileOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetUserProfileOutput",
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
				"role": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"created_at": &graphql.Field{
					Type: graphql.String,
				},
				"avatar_url": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
