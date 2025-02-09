package router

import (
	"net/http"
	"strconv"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel/request"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerStoryRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		story     = "story"
		storyByID = "story/:id"
	)
	r.Use(middleware.JWT([]byte(cfg.JWTSecret)))
	router := NewStoryRouters(cfg, ctrl)

	r.GET(storyByID, router.GetStoryByID)
	// r.GET(story, router.FindAllByUser)
	r.GET(story, router.GetStoryByName)
	// r.POST(story, router.CreateStory)
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, stories)
}

func (a *StoryRouters) GetStoryByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	story, err := a.Ctrl.StoryController.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, story)
}

func (a *StoryRouters) CreateStory(c echo.Context) error {

	// token := strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")

	story := request.StoryRequestViewModel{}
	c.Bind(story)

	if err := a.Ctrl.StoryController.CreateStory(story); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

// func (a *StoryRouters) FindAllByUser(c echo.Context) error {
// 	userID, err := strconv.ParseInt(c.QueryParam("user"), 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	stories, err := a.Ctrl.StoryController.FindAllByUser(userID)
// 	if err != nil {
// 		if err.Error() == "no stories found" {
// 			return c.JSON(http.StatusNotFound, err.Error())
// 		}
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, stories)
// }
