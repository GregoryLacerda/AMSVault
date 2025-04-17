package models

type AnimesResponse struct {
	Data []Node `json:"data"`
}
type Node struct {
	Anime AnimesResponseData `json:"node"`
}
type AnimesResponseData struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	MainPicture MainPicture `json:"main_picture"`
	NumEpisodes int64       `json:"num_episodes"`
	Synopsis    string      `json:"synopsis"`
	Status      string      `json:"status"`
}
type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
