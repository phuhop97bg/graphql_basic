package output

import "github.com/graphql-go/graphql"

func GetSongListOutput() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "getSongListOutput",
			Fields: graphql.Fields{
				"songs": &graphql.Field{
					Type: &graphql.List{
						OfType: graphql.NewObject(graphql.ObjectConfig{
							Name: "SongOutput",
							Fields: graphql.Fields{
								"id": &graphql.Field{
									Type: graphql.Int,
								},
								"title": &graphql.Field{
									Type: graphql.String,
								},
								"content_url": &graphql.Field{
									Type: graphql.String,
								},
								"view": &graphql.Field{
									Type: graphql.Int,
								},
								"image_url": &graphql.Field{
									Type: graphql.String,
								},
								"decription": &graphql.Field{
									Type: graphql.String,
								},
								"created_at": &graphql.Field{
									Type: graphql.String,
								},
								"singer": &graphql.Field{
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
