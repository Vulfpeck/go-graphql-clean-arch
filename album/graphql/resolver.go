package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/vulfpeck/go-graphql-test/album"
	"github.com/vulfpeck/go-graphql-test/models"
)

type Resolver interface {
	FetchAlbum(params graphql.ResolveParams) (interface{}, error)
	FetchAlbumById(params graphql.ResolveParams) (interface{}, error)

	SaveAlbum(params graphql.ResolveParams) (interface{}, error)
}

type resolver struct {
	AlbumService album.UseCase
}

func NewResolver(albumService album.UseCase) *resolver {
	return &resolver{AlbumService: albumService}
}

func (r resolver) FetchAlbum(params graphql.ResolveParams) (interface{}, error) {
	return r.AlbumService.GetAll()
}

func (r resolver) FetchAlbumById(params graphql.ResolveParams) (interface{}, error) {
	return r.AlbumService.FindOne(params.Args["id"].(string))
}

func (r resolver) SaveAlbum(params graphql.ResolveParams) (interface{}, error) {
	newAlbum := models.Album{}
	mapstructure.Decode(params.Args["album"], &newAlbum)
	r.AlbumService.Save(newAlbum)
	return newAlbum, nil
}
