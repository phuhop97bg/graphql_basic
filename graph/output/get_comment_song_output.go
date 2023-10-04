package output

import "github.com/graphql-go/graphql"

func GetCommetSongOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "GetSongOutput",
			Fields: graphql.Fields{
				"comments": &graphql.Field{
					Type: &graphql.List{
						OfType: graphql.NewObject(graphql.ObjectConfig{
							Name: "commentsDetailSongOutput",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"user": &graphql.Field{
									Type: userInfoOutput(),
								},
								"content": &graphql.Field{
									Type: graphql.String,
								},
							},
						}),
					},
				},
			},
		},
	)
}
