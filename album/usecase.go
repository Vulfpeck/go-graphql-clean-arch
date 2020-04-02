package album

import "github.com/vulfpeck/go-graphql-test/models"

type UseCase interface {
	GetAll() ([]models.Album, error)
	FindOne(string) (models.Album, error)
	Save(models.Album) (models.Album, error)
}
