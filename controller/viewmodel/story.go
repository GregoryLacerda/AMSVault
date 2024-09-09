package viewmodel

import "github.com.br/GregoryLacerda/AMSVault/entity"

type StoryRequestViewModel struct {
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
	User    string `json:"user"`
}

type StoryResponseViewModel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Season  int64  `json:"season"`
	Episode int64  `json:"episode"`
	Status  string `json:"status"`
}

func ParseStoryToResponseViewModel(story entity.Story) StoryResponseViewModel {
	return StoryResponseViewModel{
		ID:      story.ID,
		Name:    story.Name,
		Season:  story.Season,
		Episode: story.Episode,
		Status:  story.Status,
	}
}
