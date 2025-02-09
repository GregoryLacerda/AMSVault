package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo, cfg *config.Config, ctrl *controller.Controller) {

	defaultGroup := e.Group("/")

	registerStoryRouter(defaultGroup, cfg, ctrl)

	registerTokenRouter(defaultGroup, cfg, ctrl)

	registerUserRouter(defaultGroup, cfg, ctrl)

	registerBookmarksRouter(defaultGroup, cfg, ctrl)
}
