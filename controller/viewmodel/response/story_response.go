package response

import "github.com.br/GregoryLacerda/AMSVault/entity"

type StoryResponseViewModel struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Source       string `json:"source"`
	Description  string `json:"description"`
	TotalSeason  int64  `json:"total_season,omitempty"`
	TotalEpisode int64  `json:"total_episode,omitempty"`
	Volume       int64  `json:"total_volume,omitempty"`
	Chapter      int64  `json:"total_chapter,omitempty"`
	Status       string `json:"status"`
	entity.MainPicture
}

func ParseStoryToResponseViewModel(story entity.Story) StoryResponseViewModel {
	return StoryResponseViewModel{
		ID:           story.ID,
		Name:         story.Name,
		Source:       story.Source,
		Description:  story.Description,
		TotalSeason:  story.TotalSeason,
		TotalEpisode: story.TotalEpisode,
		Status:       story.Status,
		MainPicture:  story.MainPicture,
	}
}
