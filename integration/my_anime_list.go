package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type MALIntegration struct {
	cfg *config.Config
}

func NewMALIntegration(cfg *config.Config) *MALIntegration {
	return &MALIntegration{
		cfg: cfg,
	}
}

func (m MALIntegration) GetStoriesByName(name string) ([]entity.Story, error) {

	url := fmt.Sprintf("%s/anime?q=%s&limit=%d", m.cfg.MAL_API_URL, name, 10)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.GetToken()))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	AnimesResponse := AnimesResponse{}
	if err := json.Unmarshal(body, &AnimesResponse); err != nil {
		return nil, err
	}

	stories := []entity.Story{}
	for _, node := range AnimesResponse.Data {
		story := entity.Story{
			Name:        node.Anime.Title,
			MainPicture: entity.MainPicture{Medium: node.Anime.MainPicture.Medium, Large: node.Anime.MainPicture.Large},
		}
		stories = append(stories, story)
	}

	return stories, nil
}

func (m MALIntegration) GetToken() string {

	return m.cfg.MAL_TOKEN
}
