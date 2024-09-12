package integration

type AnimesResponse struct {
	Data []Node `json:"data"`
}
type Node struct {
	Anime AnimesResponseData `json:"node"`
}
type AnimesResponseData struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	MainPicture MainPicture `json:"main_picture"`
}
type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
