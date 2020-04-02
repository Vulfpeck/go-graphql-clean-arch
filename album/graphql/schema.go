package graphql

import (
	"github.com/graphql-go/graphql"
)

var AlbumGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"title": &graphql.Field{Type: graphql.String},
			"year":  &graphql.Field{Type: graphql.String},
			"genre": &graphql.Field{Type: graphql.String},
		},
	},
)

var AlbumInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "AlbumInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"title": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"year": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"genre": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
		Description: "Used when creating new objects",
	},
)
