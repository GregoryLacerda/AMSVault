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

	anime, err := entity.NewAnime(animeViewModel.Name, animeViewModel.Season, animeViewModel.Episode, animeViewModel.Status, animeViewModel.User)
	if err != nil {
		return err
	}

	return c.AnimeService.CreateAnime(anime)
}

func (c *AnimeController) FindByID(id string) (viewmodel.AnimeResponseViewModel, error) {
	anime, err := c.AnimeService.FindByID(id)
	if err != nil {
		return viewmodel.AnimeResponseViewModel{}, err
	}

	return viewmodel.ParseAnimeToResponseViewModel(anime), nil

}

func (c *AnimeController) FindAllByUser(user string) ([]viewmodel.AnimeResponseViewModel, error) {
	animes, err := c.AnimeService.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	var animesResponse []viewmodel.AnimeResponseViewModel
	for _, anime := range animes {
		animesResponse = append(animesResponse, viewmodel.ParseAnimeToResponseViewModel(anime))
	}

	return animesResponse, nil
}
