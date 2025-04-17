package service

import (
	"context"
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/integration"
)

type BookmarksService struct {
	data         *data.Data
	Integrations *integration.Integrations
}

func newBookmarksService(data *data.Data, integrations *integration.Integrations) *BookmarksService {
	return &BookmarksService{
		data:         data,
		Integrations: integrations,
	}
}

func (s BookmarksService) FindBookmarksByID(ctx context.Context, bookmarksID string) (retVal entity.Bookmarks, err error) {

	retVal, err = s.data.Mongo.FindOne(ctx, bookmarksID)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	if retVal.ID == "" {
		return entity.Bookmarks{}, errors.New("bookmarks not found")
	}

	return retVal, nil
}

func (s BookmarksService) FindAllByUser(ctx context.Context, userID int64) ([]entity.Bookmarks, error) {
	bookmarks, err := s.data.Mongo.FindAllByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (s BookmarksService) CreateBookmarks(ctx context.Context, userID int64, storyID int64) error {

	story, err := s.Integrations.MALIntegration.GetStoryByID(storyID)
	if err != nil {
		return err
	}

	return s.data.Mongo.Insert(ctx, userID, story)
}

func (s BookmarksService) UpdateBookmarks(ctx context.Context, bookmarks entity.Bookmarks) (entity.Bookmarks, error) {

	if bookmarks.ID == "" {
		return entity.Bookmarks{}, errors.New("empty story id")
	}

	updatedBookmarks, err := s.data.Mongo.UpdateOne(ctx, &bookmarks)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return updatedBookmarks, nil
}

func (s BookmarksService) DeleteBookmarks(ctx context.Context, bookmarksID string) error {
	return s.data.Mongo.DeleteOne(ctx, bookmarksID)
}
