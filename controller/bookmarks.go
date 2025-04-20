package controller

import (
	"context"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/response"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type BookmarksController struct {
	cfg              *config.Config
	tokenController  *service.TokenService
	bookmarksService *service.BookmarksService
	storyService     *service.StoryService
}

func newBookmarksController(cfg *config.Config, service *service.Service) *BookmarksController {
	return &BookmarksController{
		cfg:              cfg,
		tokenController:  service.TokenService,
		bookmarksService: service.BookmarksService,
		storyService:     service.StoryService,
	}
}

func (c BookmarksController) FindByID(ctx context.Context, id string) (response.BookmarksResponseVieModel, error) {
	bookmarks, err := c.bookmarksService.FindBookmarksByID(ctx, id)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	story, err := c.storyService.FindByID(bookmarks.StoryID)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	return response.ParseBookMarksToResponseViewModel(bookmarks, response.ParseStoryToResponseViewModel(story)), nil
}

func (c BookmarksController) FindAllBookmarksByUser(ctx context.Context, userID int64) (retVal []response.BookmarksResponseVieModel, err error) {
	bookmarks, err := c.bookmarksService.FindAllByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarks {
		story, err := c.storyService.FindByID(bookmark.StoryID)
		if err != nil {
			return nil, err
		}

		retVal = append(retVal, response.ParseBookMarksToResponseViewModel(bookmark, response.ParseStoryToResponseViewModel(story)))
	}

	return retVal, nil
}

func (c BookmarksController) CreateBookmarks(ctx context.Context, booksmark request.BookmarksRequestViewModel) error {
	bookmark, err := entity.NewBookmarks(booksmark)
	if err != nil {
		return err
	}

	return c.bookmarksService.CreateBookmarks(ctx, bookmark)
}

func (c BookmarksController) UpdateBookmarks(ctx context.Context, bookmarks request.BookmarksRequestViewModel) (response.BookmarksResponseVieModel, error) {

	bookmark, err := entity.NewBookmarks(bookmarks)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	updatedBookmarks, err := c.bookmarksService.UpdateBookmarks(ctx, bookmark)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	story, err := c.storyService.FindByID(updatedBookmarks.StoryID)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	return response.ParseBookMarksToResponseViewModel(updatedBookmarks, response.ParseStoryToResponseViewModel(story)), nil
}

func (c BookmarksController) DeleteBookmarks(ctx context.Context, bookmarksID string) error {
	return c.bookmarksService.DeleteBookmarks(ctx, bookmarksID)
}
