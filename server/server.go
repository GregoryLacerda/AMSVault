package server

import (
	"fmt"
	"log"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/server/router"
	"github.com.br/GregoryLacerda/AMSVault/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	srv *echo.Echo
}

func New() *Server {
	return &Server{
		srv: echo.New(),
	}
}

func (s *Server) Start(cfg *config.Config, ctrl *controller.Controller, svc *service.Service) error {

	router.Register(s.srv, cfg, ctrl)
	s.srv.Use(middleware.Logger())

	log.Printf("Starting server on port %s", cfg.WebServerPort)
	return s.srv.Start(fmt.Sprintf(":%s", cfg.WebServerPort))
}
