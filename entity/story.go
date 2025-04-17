package entity

import (
	"errors"
	"fmt"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Story struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	Source       string      `json:"source"`
	Description  string      `json:"description"`
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

func NewStory(req request.StoryRequestViewModel) (Story, error) {

	story := Story{
		Name:         req.Name,
		Source:       req.Source,
		Description:  req.Description,
		TotalSeason:  req.TotalSeason,
		TotalEpisode: req.TotalEpisode,
		TotalVolume:  req.TotalVolume,
		TotalChapter: req.TotalChapter,
		Status:       req.Status,
	}

	err := story.Validate()
	if err != nil {
		return Story{}, err
	}

	return story, nil
}

func (a *Story) Validate() error {

	if a.Name == "" {
		fmt.Println(a.Name)
		return errors.New(constants.ERROR_NAME_REQUIRED)
	}
	if a.TotalSeason < 0 {
		return errors.New(constants.ERROR_SEASON_INVALID)
	}
	if a.TotalEpisode < 0 {
		return errors.New(constants.ERROR_EPISODE_INVALID)
	}
	if a.TotalChapter < 0 {
		return errors.New(constants.ERROR_CHAPPTER_INVALID)
	}
	if a.TotalVolume < 0 {
		return errors.New(constants.ERROR_VOLUME_INVALID)
	}
	if a.Status == "" {
		return errors.New(constants.ERROR_STATUS_REQUIRED)
	}

	return nil
}
