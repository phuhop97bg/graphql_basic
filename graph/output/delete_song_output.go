package output

import "github.com/graphql-go/graphql"

func DeleteSongOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "DeleteSongOutput",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
