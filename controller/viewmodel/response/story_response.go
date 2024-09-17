package response

import "github.com.br/GregoryLacerda/AMSVault/entity"

type StoryResponseViewModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	User        string `json:"user"`
	Category    string `json:"category"`
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
		User:        story.User,
		Source:      story.Source,
		Category:    story.Category,
		Description: story.Description,
		Season:      story.Season,
		Episode:     story.Episode,
		Status:      story.Status,
		MainPicture: story.MainPicture,
	}
}
