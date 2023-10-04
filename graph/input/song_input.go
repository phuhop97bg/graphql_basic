package input

import "github.com/graphql-go/graphql"

func SongInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"singer": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"decription": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"user_id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
