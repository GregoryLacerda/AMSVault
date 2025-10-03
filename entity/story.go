package entity

import (
	"errors"
	"fmt"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Story struct {
	ID           int64
	MALID        int64
	Name         string
	Source       string
	Description  string
	TotalSeason  int64
	TotalEpisode int64
	TotalVolume  int64
	TotalChapter int64
	Status       string
	MainPicture  MainPicture
}

type MainPicture struct {
	Medium string
	Large  string
}

func NewStory(req request.StoryRequestViewModel) (Story, error) {

	story := Story{
		Name:         req.Name,
		MALID:        req.MALID,
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
