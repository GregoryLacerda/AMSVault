package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, cfg *config.Config, ctrl *controller.Controller) {

	defaultGroup := e.Group("/")

	registerLoginRouter(defaultGroup, cfg, ctrl)
	registerUserRouter(defaultGroup, cfg, ctrl)
	registerStoryRouter(defaultGroup, cfg, ctrl)
	registerBookmarksRouter(defaultGroup, cfg, ctrl)
}
