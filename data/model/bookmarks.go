package model

import (
	"time"

	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Bookmarks struct {
	ID        string    `bson:"_id"`
	UserID    int64     `bson:"user_id"`
	Story     Story     `bson:"story"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"update_at"`
	DeletedAt time.Time `bson:"deleted_at"`
}

func (b Bookmarks) ToEntity() entity.Bookmarks {
	return entity.Bookmarks{
		ID:     b.ID,
		UserID: b.UserID,
		Story:  b.Story.ToEntity(),
	}
}

func ToModelBookmarks(b entity.Bookmarks) Bookmarks {
	return Bookmarks{
		ID:        b.ID,
		UserID:    b.UserID,
		Story:     ToModelStory(b.Story),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}
