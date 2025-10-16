package entity

import (
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
)

type Bookmarks struct {
	ID             string
	UserID         int64
	StoryID        int64
	Status         string
	CurrentSeason  int64
	CurrentEpisode int64
	CurrentVolume  int64
	CurrentChapter int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func NewBookmarks(req request.BookmarksRequestViewModel) (Bookmarks, error) {

	bookmark := Bookmarks{
		ID:             req.ID,
		UserID:         req.UserID,
		StoryID:        req.StoryID,
		CurrentSeason:  req.CurrentSeason,
		CurrentEpisode: req.CurrentEpisode,
		CurrentVolume:  req.CurrentVolume,
		CurrentChapter: req.CurrentChapter,
		Status:         req.Status,
	}

	err := bookmark.Validate()
	if err != nil {
		return Bookmarks{}, err
	}

	return bookmark, nil
}

func (a *Bookmarks) Validate() error {

	if a.UserID <= 0 {
		return errors.New(constants.ERROR_USER_REQUIRED)
	}
	if a.StoryID <= 0 {
		return errors.New(constants.ERROR_STORY_NOT_FOUND)
	}

	return nil
}
