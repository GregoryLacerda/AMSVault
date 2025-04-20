package request

import (
	"errors"
)

type BookmarksRequestViewModel struct {
	ID             string `json:"id"`
	UserID         int64  `json:"user_id"`
	StoryID        int64  `json:"story_id"`
	Status         string `json:"status"`
	CurrentSeason  int64  `json:"current_season"`
	CurrentEpisode int64  `json:"current_episode"`
	CurrentVolume  int64  `json:"current_volume"`
	CurrentChapter int64  `json:"current_chapter"`
}

func (p BookmarksRequestViewModel) Validate() error {
	if p.UserID == 0 {
		return errors.New("empty bookmarks UserID")
	}

	if p.StoryID == 0 {
		return errors.New("empty bookmarks storyID")
	}

	return nil
}
