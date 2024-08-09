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
		anime     = ""
		animeByID = "/:id"
	)
	r.Use(middleware.JWT([]byte(cfg.JWTSecret)))
	router := NewAnimeRouters(cfg, ctrl)

	r.GET(anime, router.GetAnime)
	r.GET(animeByID, router.GetAnimeByID)
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

func (a *AnimeRouters) GetAnime(c echo.Context) error {
	return nil
}

func (a *AnimeRouters) GetAnimeByID(c echo.Context) error {
	return nil
}

func (a *AnimeRouters) CreateAnime(c echo.Context) error {
	anime := new(viewmodel.AnimeRequestViewModel)
	c.Bind(anime)

	if err := a.Ctrl.AnimeController.CreateAnime(anime); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "")
}

func (a *AnimeRouters) UpdateAnime(c echo.Context) error {
	return nil
}

func (a *AnimeRouters) DeleteAnime(c echo.Context) error {
	return nil
}
