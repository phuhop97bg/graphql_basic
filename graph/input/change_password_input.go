package input

import "github.com/graphql-go/graphql"

func ChangePasswordInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "ChangePasswordOutput",
		Fields: graphql.InputObjectConfigFieldMap{
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"old_password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}
