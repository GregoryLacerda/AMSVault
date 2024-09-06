package service

import (
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type AnimeService struct {
	data *data.Data
}

func newAnimeService(data *data.Data) *AnimeService {
	return &AnimeService{
		data: data,
	}
}

func (s *AnimeService) CreateAnime(anime *entity.Anime) error {
	if err := anime.Validate(); err != nil {
		return err
	}

	err := s.data.Mongo.Insert("anime", anime)
	if err != nil {
		return err
	}

	return nil
}

func (s *AnimeService) FindAllByUser(user string) ([]entity.Anime, error) {
	animes, err := s.data.Mongo.FindAllByField("anime", "user_hash", user)
	if err != nil {
		return nil, err
	}

	return animes, nil
}
