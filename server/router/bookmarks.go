package router

import (
	"net/http"
	"strconv"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
	"github.com.br/GregoryLacerda/AMSVault/server/middleware"
	"github.com/labstack/echo/v4"
)

func registerBookmarksRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		bookmarks       = "bookmarks"
		bookmarksID     = "bookmarks/:id"
		bookmarksUserID = "bookmarks/user/:user_id"
	)

	r.Use(middleware.JWTMiddleware(cfg))
	router := NewBookmarksRouter(cfg, *ctrl)

	r.GET(bookmarksID, router.FindBookmarksByID)
	r.GET(bookmarksUserID, router.FindAllBookmarksByUser)
	r.POST(bookmarks, router.CreateBookmarks)
	r.PUT(bookmarksID, router.UpdateBookmarks)
	r.DELETE(bookmarksID, router.DeleteBookmarks)
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

	id := c.Param("id")

	bookmarks, err := p.ctrl.BookmarksController.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func (p *Bookmarks) FindAllBookmarksByUser(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	bookmarks, err := p.ctrl.BookmarksController.FindAllBookmarksByUser(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func (p Bookmarks) CreateBookmarks(c echo.Context) error {

	bookmarksRequest := request.BookmarksRequestViewModel{}
	if err := c.Bind(&bookmarksRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
	}

	err := p.ctrl.BookmarksController.CreateBookmarks(c.Request().Context(), bookmarksRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Bookmarks created successfully"})
}

func (p Bookmarks) UpdateBookmarks(c echo.Context) error {

	bookmarksRequest := request.BookmarksRequestViewModel{}
	if err := c.Bind(&bookmarksRequest); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
	}

	if bookmarksRequest.ID == "" {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError("empty bookmarks ID"))
	}

	bookmarks, err := p.ctrl.BookmarksController.UpdateBookmarks(c.Request().Context(), bookmarksRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func (p Bookmarks) DeleteBookmarks(c echo.Context) error {

	bokmarksID := c.QueryParam("id")

	err := p.ctrl.BookmarksController.DeleteBookmarks(c.Request().Context(), bokmarksID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Bookmarks deleted successfully"})
}
