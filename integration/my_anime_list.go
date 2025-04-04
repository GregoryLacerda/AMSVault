package integration

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type MALIntegration struct {
	cfg *config.Config
}

func newMALIntegration(cfg *config.Config) *MALIntegration {
	return &MALIntegration{
		cfg: cfg,
	}
}

func (m MALIntegration) GetStoriesByName(name string) ([]entity.Story, error) {

	animeResponse, err := m.getAnimeList(name)
	if err != nil {
		return nil, err
	}

	stories := []entity.Story{}
	for _, node := range animeResponse.Data {
		story := entity.Story{
			Name:         node.Anime.Title,
			MainPicture:  entity.MainPicture{Medium: node.Anime.MainPicture.Medium, Large: node.Anime.MainPicture.Large},
			Description:  node.Anime.Synopsis,
			TotalEpisode: int64(node.Anime.NumEpisodes),
			Status:       node.Anime.Status,
		}
		stories = append(stories, story)
	}

	return stories, nil
}

func (m MALIntegration) GetToken() string {

	return m.cfg.MAL_TOKEN
}
