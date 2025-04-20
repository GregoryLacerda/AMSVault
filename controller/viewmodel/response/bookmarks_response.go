package response

import (
	"time"

	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type BookmarksResponseVieModel struct {
	ID             string                 `json:"id"`
	UserID         int64                  `json:"user"`
	Story          StoryResponseViewModel `json:"story"`
	Status         string                 `json:"status"`
	CurrentSeason  int64                  `json:"current_season,omitempty"`
	CurrentEpisode int64                  `json:"current_episode,omitempty"`
	CurrentVolume  int64                  `json:"current_volume,omitempty"`
	CurrentChapter int64                  `json:"current_chapter,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"update_at"`
	DeletedAt      time.Time              `json:"deleted_at"`
}

func ParseBookMarksToResponseViewModel(bookmarks entity.Bookmarks, story StoryResponseViewModel) BookmarksResponseVieModel {
	return BookmarksResponseVieModel{
		ID:             bookmarks.ID,
		UserID:         bookmarks.UserID,
		Story:          story,
		Status:         bookmarks.Status,
		CurrentSeason:  bookmarks.CurrentSeason,
		CurrentEpisode: bookmarks.CurrentEpisode,
		CurrentVolume:  bookmarks.CurrentVolume,
		CurrentChapter: bookmarks.CurrentChapter,
		CreatedAt:      bookmarks.CreatedAt,
		UpdatedAt:      bookmarks.UpdatedAt,
		DeletedAt:      bookmarks.DeletedAt,
	}
}
