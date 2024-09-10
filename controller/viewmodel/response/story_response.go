package response

import "github.com.br/GregoryLacerda/AMSVault/entity"

type StoryResponseViewModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Season      int64  `json:"season,omitempty"`
	Episode     int64  `json:"episode,omitempty"`
	Volume      int64  `json:"volume,omitempty"`
	Chapter     int64  `json:"chapter,omitempty"`
	Status      string `json:"status"`
}

func ParseStoryToResponseViewModel(story entity.Story) StoryResponseViewModel {
	return StoryResponseViewModel{
		ID:          story.ID,
		Name:        story.Name,
		Kind:        story.Kind,
		Category:    story.Category,
		Description: story.Description,
		Season:      story.Season,
		Episode:     story.Episode,
		Status:      story.Status,
	}
}
