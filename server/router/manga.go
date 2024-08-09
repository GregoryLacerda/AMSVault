package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func registerMangaRouter(r *echo.Group, cfg *config.Config, ctrl *controller.Controller) {

	const (
		manga     = "/"
		mangaByID = "/:id"
	)

	r.Use(middleware.JWT([]byte(cfg.JWTSecret)))
	router := NewMangaRouters(cfg, ctrl)

	r.GET(manga, router.GetManga)
	r.GET(mangaByID, router.GetMangaByID)
	r.POST(manga, router.CreateManga)
	r.PUT(mangaByID, router.UpdateManga)
	r.DELETE(mangaByID, router.DeleteManga)
}

type MangaRouters struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}

func NewMangaRouters(cfg *config.Config, ctrl *controller.Controller) *MangaRouters {
	return &MangaRouters{
		Ctrl: ctrl,
		cfg:  cfg,
	}
}

func (a *MangaRouters) GetManga(c echo.Context) error {
	return nil
}

func (a *MangaRouters) GetMangaByID(c echo.Context) error {
	return nil
}

func (a *MangaRouters) CreateManga(c echo.Context) error {
	return nil
}

func (a *MangaRouters) UpdateManga(c echo.Context) error {
	return nil
}

func (a *MangaRouters) DeleteManga(c echo.Context) error {
	return nil
}
