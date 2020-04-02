package song

import (
	"errors"
	"github.com/vulfpeck/go-graphql-test/models"
)

var songs = []models.Song{
	{
		ID:       "1",
		Album:    "1",
		Title:    "Lmao",
		Duration: "230",
		Type:     "2",
	},
	{
		ID:       "2",
		Album:    "1",
		Title:    "Lmao",
		Duration: "230",
		Type:     "2",
	},
	{
		ID:       "3",
		Album:    "2",
		Title:    "Lmao",
		Duration: "230",
		Type:     "2",
	},
}

type RepositoryInterface interface {
	GetAll() []models.Song
	FindOne(string) (models.Song, error)
	Save(models.Song) models.Song
}

type RepositorySong struct {
}

func (r RepositorySong) GetAll() []models.Song {
	return songs
}

func (r RepositorySong) FindOne(id string) (models.Song, error) {
	for _, song := range songs {
		if song.ID == id {
			return song, nil
		}
	}
	return models.Song{}, errors.New("song not found")
}

func (r RepositorySong) Save(song models.Song) models.Song {
	songs = append(songs, song)
	return song
}
