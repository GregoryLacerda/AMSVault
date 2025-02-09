package service

import (
	"context"
	"errors"
	"strconv"

	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type BookmarksService struct {
	data *data.Data
}

func newBookmarksService(data *data.Data) *BookmarksService {
	return &BookmarksService{
		data: data,
	}
}

func (p BookmarksService) FindBookmarksByID(ctx context.Context, bookmarksID string) (retVal entity.Bookmarks, err error) {

	retVal, err = p.data.Mongo.FindOne(ctx, bookmarksID)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	if retVal.ID == "" {
		return entity.Bookmarks{}, errors.New("bookmarks not found")
	}

	return retVal, nil
}

func (s *StoryService) FindAllByUser(ctx context.Context, userID int64) ([]entity.Bookmarks, error) {

	userIDStr := strconv.FormatInt(userID, 10)

	bookmarks, err := s.data.Mongo.FindAllByField(ctx, "user", userIDStr)
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}
