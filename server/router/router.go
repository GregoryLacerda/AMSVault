package router

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo, cfg *config.Config, ctrl *controller.Controller) {

	defaultGroup := e.Group("/")

	defaultStoryGroup := defaultGroup.Group("story")

	registerStoryRouter(defaultStoryGroup, cfg, ctrl)

	RegisterTokenRouter(defaultGroup, cfg, ctrl)

	RegisterUserRouter(defaultGroup, cfg, ctrl)
}
