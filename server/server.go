package server

import (
	"fmt"
	"log"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/server/router"
	"github.com.br/GregoryLacerda/AMSVault/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func New() *Server {
	return &Server{
		echo: echo.New(),
	}
}

func (s *Server) Start(cfg *config.Config, ctrl *controller.Controller, svc *service.Service) error {

	s.echo.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	router.Register(s.echo, cfg, ctrl)
	s.echo.Use(middleware.Logger())

	log.Printf("Starting server on port %s", cfg.WebServerPort)
	return s.echo.Start(fmt.Sprintf(":%s", cfg.WebServerPort))
}
