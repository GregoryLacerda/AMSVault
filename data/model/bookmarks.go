package model

import (
	"time"

	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Bookmarks struct {
	ID             string    `bson:"_id"`
	UserID         int64     `bson:"user_id"`
	StoryID        int64     `bson:"story_id"`
	Status         string    `bson:"status"`
	CurrentSeason  int64     `bson:"current_season"`
	CurrentEpisode int64     `bson:"current_episode"`
	CurrentVolume  int64     `bson:"current_volume"`
	CurrentChapter int64     `bson:"current_chapter"`
	CreatedAt      time.Time `bson:"created_at"`
	UpdatedAt      time.Time `bson:"update_at"`
	DeletedAt      time.Time `bson:"deleted_at"`
}

func (b Bookmarks) ToEntity() entity.Bookmarks {
	return entity.Bookmarks{
		ID:             b.ID,
		UserID:         b.UserID,
		StoryID:        b.StoryID,
		Status:         b.Status,
		CurrentSeason:  b.CurrentSeason,
		CurrentEpisode: b.CurrentEpisode,
		CurrentVolume:  b.CurrentVolume,
		CurrentChapter: b.CurrentChapter,
		CreatedAt:      b.CreatedAt,
		UpdatedAt:      b.UpdatedAt,
		DeletedAt:      b.DeletedAt,
	}
}

func ToModelBookmarks(b entity.Bookmarks) Bookmarks {
	return Bookmarks{
		ID:             b.ID,
		UserID:         b.UserID,
		StoryID:        b.StoryID,
		Status:         b.Status,
		CurrentSeason:  b.CurrentSeason,
		CurrentEpisode: b.CurrentEpisode,
		CurrentVolume:  b.CurrentVolume,
		CurrentChapter: b.CurrentChapter,
		CreatedAt:      b.CreatedAt,
		UpdatedAt:      b.UpdatedAt,
		DeletedAt:      b.DeletedAt,
	}
}
