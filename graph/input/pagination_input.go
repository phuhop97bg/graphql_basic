package input

import "github.com/graphql-go/graphql"

func PaginationInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "PaginationInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"page_num": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"page_size": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	})
}
