package output

import "github.com/graphql-go/graphql"

func LoginOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "LoginOutput",
			Fields: graphql.Fields{
				"token": &graphql.Field{
					Type: graphql.String,
				},
				"token_expired_at": &graphql.Field{
					Type: graphql.DateTime,
				},
				"role": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
