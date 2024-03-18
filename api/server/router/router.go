package router

import (
	"amsvault/api/config"

	"github.com/labstack/echo"
)

func Register(e *echo.Echo, cfg *config.Config) {
	defaultGroup := e.Group("/api/v1")

	animeGroup := defaultGroup.Group("/anime")
	mangaGroup := defaultGroup.Group("/manga")
	seriesGroup := defaultGroup.Group("/series")

}

func registerDefaultApi(group *echo.Group, cfg *config.Config) {

}
