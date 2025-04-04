package entity

import (
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Bookmarks struct {
	ID             string    `json:"id"`
	UserID         int64     `json:"user"`
	Story          Story     `json:"story"`
	CurrentSeason  int64     `json:"current_season,omitempty"`
	CurrentEpisode int64     `json:"current_episode,omitempty"`
	CurrentVolume  int64     `json:"current_volume,omitempty"`
	CurrentChapter int64     `json:"current_chapter,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"update_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

func NewBookmarks(req request.BookmarksRequestViewModel) (Bookmarks, error) {

	story, err := NewStory(req.Story)
	if err != nil {
		return Bookmarks{}, err
	}

	bookmark := Bookmarks{
		ID:     req.ID,
		UserID: req.UserID,
		Story:  story,
	}

	err = bookmark.Validate()
	if err != nil {
		return Bookmarks{}, err
	}

	return bookmark, nil
}

func (a *Bookmarks) Validate() error {

	if a.ID == "" {
		return errors.New(constants.ERROR_ID_REQUIRED)
	}
	if a.UserID <= 0 {
		return errors.New(constants.ERROR_USER_REQUIRED)
	}
	if a.Story.ID <= 0 {
		return errors.New(constants.ERROR_STORY_NOT_FOUND)
	}

	return nil
}
