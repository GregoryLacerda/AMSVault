package router

import (
	"context"
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerBookmarksRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		bookmarks       = "bookmarks"
		bookmarksID     = "/:id"
		bookmarksUserID = "/:userID"
	)

	r.Use(middleware.JWT([]byte(cfg.JWTSecret)))
	router := NewBookmarksRouter(cfg, *ctrl)

	r.GET(bookmarks+bookmarksID, router.FindBookmarksByID)
}

type Bookmarks struct {
	cfg  *config.Config
	ctrl *controller.Controller
}

func NewBookmarksRouter(cfg *config.Config, ctrl controller.Controller) *Bookmarks {
	return &Bookmarks{
		cfg:  cfg,
		ctrl: &ctrl,
	}
}

func (p *Bookmarks) FindBookmarksByID(c echo.Context) error {

	id := c.QueryParam("id")

	bookmarks, err := p.ctrl.BookmarksController.FindByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, bookmarks)
}
