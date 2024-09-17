package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/integration/models"
)

func (m *MALIntegration) getAnimeList(name string) (*models.AnimesResponse, error) {

	fields := "num_episodes,synopsis,status"
	query := strings.ReplaceAll(name, " ", "%20")

	url := fmt.Sprintf("%s/anime?q=%s&limit=%d&fields=%s", m.cfg.MAL_API_URL, query, 10, fields)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.GetToken()))
	req.Header.Add("Content-Type", "application/json")

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

	AnimesResponse := models.AnimesResponse{}
	if err := json.Unmarshal(body, &AnimesResponse); err != nil {
		return nil, err
	}

	return &AnimesResponse, nil
}
