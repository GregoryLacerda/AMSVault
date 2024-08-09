package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type AnimeController struct {
	cfg          *config.Config
	AnimeService *service.AnimeService
}

func newAnimeController(cfg *config.Config, service *service.Service) *AnimeController {
	return &AnimeController{
		cfg:          cfg,
		AnimeService: service.AnimeService,
	}
}

func (c *AnimeController) CreateAnime(animeViewModel *viewmodel.AnimeRequestViewModel) error {

	anime, err := entity.NewAnime(animeViewModel.Name, animeViewModel.Season, animeViewModel.Episode, animeViewModel.Status)
	if err != nil {
		return err
	}

	return c.AnimeService.CreateAnime(anime)
}
