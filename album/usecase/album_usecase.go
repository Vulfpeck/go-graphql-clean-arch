package usecase

import (
	"github.com/vulfpeck/go-graphql-test/album"
	"github.com/vulfpeck/go-graphql-test/models"
)

type albumUsecase struct {
	albumRepository album.Repository
}

func (a albumUsecase) GetAll() ([]models.Album, error) {
	return a.albumRepository.GetAll()
}

func (a albumUsecase) FindOne(id string) (models.Album, error) {
	return a.albumRepository.FindOne(id)
}

func (a albumUsecase) Save(newAlbum models.Album) (models.Album, error) {
	return a.albumRepository.Save(newAlbum)
}

func NewAlbumUsecase(albumRepository album.Repository) *albumUsecase {
	return &albumUsecase{albumRepository: albumRepository}
}
