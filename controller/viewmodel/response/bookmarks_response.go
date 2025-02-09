package response

import (
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type BookmarksResponseVieModel struct {
	ID     string       `json:"id"`
	UserID int64        `json:"user"`
	Story  entity.Story `json:"story"`
}

func ParseBookMarksToResponseViewModel(bookmarks entity.Bookmarks) BookmarksResponseVieModel {
	return BookmarksResponseVieModel{
		ID:     bookmarks.ID,
		UserID: bookmarks.UserID,
		Story:  bookmarks.Story,
	}
}
