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
}

func newBookmarksController(cfg *config.Config, service *service.Service) *BookmarksController {
	return &BookmarksController{
		cfg:              cfg,
		tokenController:  service.TokenService,
		bookmarksService: service.BookmarksService,
	}
}

func (c BookmarksController) FindByID(ctx context.Context, id string) (response.BookmarksResponseVieModel, error) {
	bookmarks, err := c.bookmarksService.FindBookmarksByID(ctx, id)
	if err != nil {
		return response.BookmarksResponseVieModel{}, err
	}

	return response.ParseBookMarksToResponseViewModel(bookmarks), nil

}

func (c BookmarksController) FindAllBookmarksByUser(ctx context.Context, userID int64) (retVal []response.BookmarksResponseVieModel, err error) {
	bookmarks, err := c.bookmarksService.FindAllByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarks {
		retVal = append(retVal, response.ParseBookMarksToResponseViewModel(bookmark))
	}

	return retVal, nil
}

func (c BookmarksController) CreateBookmarks(ctx context.Context, userID int64, storyID int64) error {
	return c.bookmarksService.CreateBookmarks(ctx, userID, storyID)
}

func (c BookmarksController) UpdateBookmarks(ctx context.Context, bookmarks request.BookmarksRequestViewModel) (entity.Bookmarks, error) {

	bookmark, err := entity.NewBookmarks(bookmarks)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return c.bookmarksService.UpdateBookmarks(ctx, bookmark)
}

func (c BookmarksController) DeleteBookmarks(ctx context.Context, bookmarksID string) error {
	return c.bookmarksService.DeleteBookmarks(ctx, bookmarksID)
}
