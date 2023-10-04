package output

import "github.com/graphql-go/graphql"

func ChangePasswordOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "ChangePasswordOutput",
			Fields: graphql.Fields{
				"token": &graphql.Field{
					Type: graphql.String,
				},
				"token_expired_at": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		},
	)
}
