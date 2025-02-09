package entity

import (
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Story struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Source      string      `json:"source"`
	Description string      `json:"description"`
	Season      int64       `json:"season,omitempty"`
	Episode     int64       `json:"episode,omitempty"`
	Volume      int64       `json:"volume,omitempty"`
	Chapter     int64       `json:"chapter,omitempty"`
	Status      string      `json:"status"`
	MainPicture MainPicture `json:"main_picture"`
}

type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

func NewStory(req request.StoryRequestViewModel) (Story, error) {

	story := Story{
		Name:        req.Name,
		Source:      req.Source,
		Description: req.Description,
		Season:      req.Season,
		Episode:     req.Episode,
		Volume:      req.Volume,
		Chapter:     req.Chapter,
		Status:      req.Status,
	}

	err := story.Validate()
	if err != nil {
		return Story{}, err
	}

	return story, nil
}

func (a *Story) Validate() error {

	if a.Name == "" {
		return errors.New(constants.ERROR_NAME_REQUIRED)
	}
	if a.Season < 0 {
		return errors.New(constants.ERROR_SEASON_INVALID)
	}
	if a.Episode < 0 {
		return errors.New(constants.ERROR_EPISODE_INVALID)
	}
	if a.Chapter < 0 {
		return errors.New(constants.ERROR_CHAPPTER_INVALID)
	}
	if a.Volume < 0 {
		return errors.New(constants.ERROR_VOLUME_INVALID)
	}
	if a.Status == "" {
		return errors.New(constants.ERROR_STATUS_REQUIRED)
	}

	return nil
}
