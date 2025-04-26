package integration

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/integration/models"
)

type MALIntegration struct {
	cfg *config.Config
}

func newMALIntegration(cfg *config.Config) *MALIntegration {
	return &MALIntegration{
		cfg: cfg,
	}
}

func (m *MALIntegration) GetStoriesByName(name string) ([]entity.Story, error) {

	animeResponse, err := m.getAnimeByName(name)
	if err != nil {
		return nil, err
	}

	stories := []entity.Story{}
	for _, node := range animeResponse.Data {
		story := m.mapResponseToStory(node.Anime)
		stories = append(stories, story)
	}

	return stories, nil
}

func (m *MALIntegration) GetStoryByID(id int64) (entity.Story, error) {

	animeResponse, err := m.getAnimeById(id)
	if err != nil {
		return entity.Story{}, err
	}

	return m.mapResponseToStory(animeResponse), nil
}

func (m *MALIntegration) GetToken() string {

	return m.cfg.MAL_TOKEN
}

func (m *MALIntegration) mapResponseToStory(anime models.AnimesResponseData) entity.Story {
	return entity.Story{
		MALID: anime.ID,
		Name:  anime.Title,
		MainPicture: entity.MainPicture{
			Medium: anime.MainPicture.Medium,
			Large:  anime.MainPicture.Large,
		},
		Description:  anime.Synopsis,
		TotalEpisode: anime.NumEpisodes,
		Status:       anime.Status,
		Source:       "anime",
	}
}
