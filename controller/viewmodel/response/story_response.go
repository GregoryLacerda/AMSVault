package response

import "github.com.br/GregoryLacerda/AMSVault/entity"

type StoryResponseViewModel struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	UserID      int64  `json:"user"`
	Source      string `json:"source"`
	Description string `json:"description"`
	Season      int64  `json:"season,omitempty"`
	Episode     int64  `json:"episode,omitempty"`
	Volume      int64  `json:"volume,omitempty"`
	Chapter     int64  `json:"chapter,omitempty"`
	Status      string `json:"status"`
	entity.MainPicture
}

func ParseStoryToResponseViewModel(story entity.Story) StoryResponseViewModel {
	return StoryResponseViewModel{
		ID:          story.ID,
		Name:        story.Name,
		UserID:      story.UserID,
		Source:      story.Source,
		Description: story.Description,
		Season:      story.Season,
		Episode:     story.Episode,
		Status:      story.Status,
		MainPicture: story.MainPicture,
	}
}
