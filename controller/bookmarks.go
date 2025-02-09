package controller

import (
	"context"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/response"
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
