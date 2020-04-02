package root_schema

import (
	"github.com/graphql-go/graphql"
	albumSchema "github.com/vulfpeck/go-graphql-test/album/graphql"
)

type Schema struct {
	albumResolver albumSchema.Resolver
}

func NewRootSchema(albumResolver albumSchema.Resolver) *Schema {
	return &Schema{albumResolver: albumResolver}
}

func (s *Schema) RootQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"Albums": &graphql.Field{
					Type: graphql.NewList(albumSchema.AlbumGraphQL),
					//Resolve: s.albumResolver.FetchAlbum,
				},
				"Album": &graphql.Field{
					Type: albumSchema.AlbumGraphQL,
					//Resolve: s.albumResolver.FetchAlbumById,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
					},
				},
			},
		},
	)
}

func (s *Schema) RootMutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"CreateAlbum": &graphql.Field{
				Type: albumSchema.AlbumGraphQL,
				Args: graphql.FieldConfigArgument{
					"album": &graphql.ArgumentConfig{Type: albumSchema.AlbumInput},
				},
				Resolve: s.albumResolver.SaveAlbum,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}
