package request

type StoryRequestViewModel struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Source       string `json:"source"`
	TotalSeason  int64  `json:"total_season,omitempty"`
	TotalEpisode int64  `json:"total_episode,omitempty"`
	TotalVolume  int64  `json:"total_volume,omitempty"`
	TotalChapter int64  `json:"total_chapter,omitempty"`
	Status       string `json:"status"`
}
