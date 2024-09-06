package viewmodel

import "github.com.br/GregoryLacerda/AMSVault/entity"

type AnimeRequestViewModel struct {
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
	User    string `json:"user"`
}

type AnimeResponseViewModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
}

func ParseAnimeToResponseViewModel(anime entity.Anime) AnimeResponseViewModel {
	return AnimeResponseViewModel{
		ID:      anime.ID,
		Name:    anime.Name,
		Season:  anime.Season,
		Episode: anime.Episode,
		Status:  anime.Status,
	}
}
