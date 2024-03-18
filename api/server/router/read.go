package router

import (
	"amsvault/api/config"
	"amsvault/api/controller"

	"github.com/labstack/echo"
)

func RegisterReadRoutes(r *echo.Echo, cfg *config.Config) {
	const readRoute = "/read"

}

type ReadRoute struct {
	Ctrl *controller.Controller
	cfg  *config.Config
}
