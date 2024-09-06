package router

import (
	"net/http"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerAnimeRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		anime       = ""
		animeByID   = "/:id"
		animeByUser = "/:user"
	)
	r.Use(middleware.JWT([]byte(cfg.JWTSecret)))
	router := NewAnimeRouters(cfg, ctrl)

	r.GET(animeByID, router.GetAnimeByID)
	r.GET(animeByUser, router.FindAllByUser)
	r.POST(anime, router.CreateAnime)
	r.PUT(animeByID, router.UpdateAnime)
	r.DELETE(animeByID, router.DeleteAnime)
}

type AnimeRouters struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewAnimeRouters(cfg *config.Config, ctrl *controller.Controller) *AnimeRouters {
	return &AnimeRouters{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *AnimeRouters) GetAnimeByID(c echo.Context) error {
	id := c.Param("id")
	anime, err := a.Ctrl.AnimeController.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, anime)
}

func (a *AnimeRouters) CreateAnime(c echo.Context) error {
	anime := new(viewmodel.AnimeRequestViewModel)
	c.Bind(anime)

	if err := a.Ctrl.AnimeController.CreateAnime(anime); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (a *AnimeRouters) FindAllByUser(c echo.Context) error {
	user := c.Param("user")

	animes, err := a.Ctrl.AnimeController.FindAllByUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, animes)
}

func (a *AnimeRouters) UpdateAnime(c echo.Context) error {
	return nil
}

func (a *AnimeRouters) DeleteAnime(c echo.Context) error {
	return nil
}
