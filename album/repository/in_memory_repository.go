package repository

import (
	"errors"
	"github.com/vulfpeck/go-graphql-test/album"
	"github.com/vulfpeck/go-graphql-test/models"
)

type inMemoryAlbumRepo struct {
	albums []models.Album
}

func NewInMemoryAlbumRepo() album.Repository {
	return &inMemoryAlbumRepo{
		albums: []models.Album{},
	}
}

func (i inMemoryAlbumRepo) GetAll() ([]models.Album, error) {
	return i.albums, nil
}

func (i inMemoryAlbumRepo) FindOne(id string) (models.Album, error) {
	for _, album := range i.albums {
		if id == album.ID {
			return album, nil
		}
	}
	return models.Album{}, errors.New("not found")
}

func (i inMemoryAlbumRepo) Save(newAlbum models.Album) (models.Album, error) {
	i.albums = append(i.albums, newAlbum)
	return newAlbum, nil
}
