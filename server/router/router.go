package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo, cfg *config.Config, ctrl *controller.Controller) {

	defaultGroup := e.Group("/")

	defaultAnimeGroup := defaultGroup.Group("anime")
	defaultMangaGroup := defaultGroup.Group("manga")
	defaultSerieGroup := defaultGroup.Group("serie")

	registerAnimeRouter(defaultAnimeGroup, cfg, ctrl)
	registerMangaRouter(defaultMangaGroup, cfg, ctrl)
	registerSerieRouter(defaultSerieGroup, cfg, ctrl)

	RegisterTokenRouter(defaultGroup, cfg, ctrl)
}
