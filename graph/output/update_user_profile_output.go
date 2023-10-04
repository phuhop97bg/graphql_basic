package output

import "github.com/graphql-go/graphql"

func UpdateUserProfileOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "UpdateUserProfileOutput",
			Fields: graphql.Fields{
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
