package viewmodel

type AnimeRequestViewModel struct {
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
}

type AnimeResponseViewModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
}
