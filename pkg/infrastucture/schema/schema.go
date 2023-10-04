package schema

import (
	"backend-food/pkg/infrastucture/db"

	"github.com/graphql-go/graphql"
)

func NewAnonymousSchema(database db.Database) *graphql.Schema {
	repoContainer := GetContainerRepo(database)
	myschema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    GetAnonymousQuery(repoContainer),
			Mutation: GetAnonymousMutation(repoContainer),
		},
	)
	return &myschema
}

func NewClientSchema(database db.Database) *graphql.Schema {
	repoContainer := GetContainerRepo(database)
	myschema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    GetClientQuery(repoContainer),
			Mutation: GetClientMutation(repoContainer),
		},
	)
	return &myschema
}
func NewAdminSchema(database db.Database) *graphql.Schema {
	repoContainer := GetContainerRepo(database)
	myschema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    GetAdminQuery(repoContainer),
			Mutation: GetAdminMutation(repoContainer),
		},
	)
	return &myschema
}
