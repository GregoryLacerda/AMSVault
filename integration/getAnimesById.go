package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/integration/models"
)

func (m *MALIntegration) getAnimeById(id int64) (retVal models.AnimesResponseData, err error) {

	fields := "id,num_episodes,synopsis,status"

	url := fmt.Sprintf("%s/anime/%d?fields=%s", m.cfg.MAL_API_URL, id, fields)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return retVal, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.GetToken()))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return retVal, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return retVal, err
	}

	bodyStr := string(body)

	if res.StatusCode != http.StatusOK {
		return retVal, fmt.Errorf("MALApi error => %s", bodyStr)
	}

	if err := json.Unmarshal(body, &retVal); err != nil {
		return retVal, err
	}

	return retVal, nil
}
