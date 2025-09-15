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

func registerStoryRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		story     = "story"
		storyByID = "story/:id"
	)
	r.Use(middleware.JWTMiddleware(cfg))
	router := NewStoryRouters(cfg, ctrl)

	r.GET(storyByID, router.GetStoryByID)
	r.GET(story, router.GetStoryByName)
	r.POST(story, router.CreateStory)
}

type StoryRouters struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewStoryRouters(cfg *config.Config, ctrl *controller.Controller) *StoryRouters {
	return &StoryRouters{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *StoryRouters) GetStoryByName(c echo.Context) error {
	name := c.QueryParam("name")

	stories, err := a.Ctrl.StoryController.FindByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, stories)
}

func (a *StoryRouters) GetStoryByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewInternalError("GetStoryByID", err))
	}

	story, err := a.Ctrl.StoryController.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, story)
}

func (a *StoryRouters) CreateStory(c echo.Context) error {
	story := request.StoryRequestViewModel{}
	if err := c.Bind(&story); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
	}

	if err := a.Ctrl.StoryController.CreateStory(story); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Story created successfully"})
}
