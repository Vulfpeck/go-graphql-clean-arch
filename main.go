package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	_albumDelivery "github.com/vulfpeck/go-graphql-test/album/graphql"
	_albumRepo "github.com/vulfpeck/go-graphql-test/album/repository"
	_albumUsecase "github.com/vulfpeck/go-graphql-test/album/usecase"
	rootSchema "github.com/vulfpeck/go-graphql-test/root_schema"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	albumRepo := _albumRepo.NewInMemoryAlbumRepo()
	albumUsecase := _albumUsecase.NewAlbumUsecase(albumRepo)
	albumResolver := _albumDelivery.NewResolver(albumUsecase)
	schema := rootSchema.NewRootSchema(albumResolver)
	gqlSchema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    schema.RootQuery(),
			Mutation: schema.RootMutation(),
		},
	)
	if err != nil {
		panic("Could not parse schema")
	}
	h := handler.New(&handler.Config{
		Schema:     &gqlSchema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
		FormatErrorFn: func(err error) gqlerrors.FormattedError {
			fmt.Println(err)
			return gqlerrors.FormatError(err)
		},
	})

	http.Handle("/graphql", h)
	fmt.Print("serving at :1234")
	http.ListenAndServe(":1234", nil)
}
