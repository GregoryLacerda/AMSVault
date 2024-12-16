package entity

import (
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Story struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"user"`
	Name        string      `json:"name"`
	Source      string      `json:"source"`
	Description string      `json:"description"`
	Season      int64       `json:"season,omitempty"`
	Episode     int64       `json:"episode,omitempty"`
	Volume      int64       `json:"volume,omitempty"`
	Chapter     int64       `json:"chapter,omitempty"`
	Status      string      `json:"status"`
	MainPicture MainPicture `json:"main_picture"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"update_at"`
	DeletedAt   time.Time   `json:"deleted_at"`
}

type MainPicture struct {
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

func NewStory(req request.StoryRequestViewModel) (*Story, error) {

	story := &Story{
		UserID:      req.UserID,
		Name:        req.Name,
		Source:      req.Source,
		Description: req.Description,
		Season:      req.Season,
		Episode:     req.Episode,
		Volume:      req.Volume,
		Chapter:     req.Chapter,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Date(0001, 01, 01, 01, 01, 01, 01, time.UTC),
		DeletedAt:   time.Date(0001, 01, 01, 01, 01, 01, 01, time.UTC),
	}

	err := story.Validate()
	if err != nil {
		return nil, err
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
	if a.UserID == 0 {
		return errors.New(constants.ERROR_USER_REQUIRED)
	}

	return nil
}
