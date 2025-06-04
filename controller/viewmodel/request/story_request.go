package request

type StoryRequestViewModel struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	MALID        int64       `json:"mal_id"`
	Description  string      `json:"description"`
	Source       string      `json:"source"`
	TotalSeason  int64       `json:"total_season,omitempty"`
	TotalEpisode int64       `json:"total_episode,omitempty"`
	TotalVolume  int64       `json:"total_volume,omitempty"`
	TotalChapter int64       `json:"total_chapter,omitempty"`
	Status       string      `json:"status"`
	MainPicture  MainPicture `json:"main_picture"`
}

type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}
