package server

import (
	"amsvault/api/config"
	"fmt"

	"github.com/labstack/echo"
)

type Server struct {
	srv *echo.Echo
}

func Instance() *Server {
	return &Server{srv: echo.New()}
}

func (s *Server) Run(cfg *config.Config) error {

	fmt.Println("port: ", cfg.Internal.Port)

	if err := s.srv.Start(fmt.Sprintf(":%s", cfg.Internal.Port)); err != nil {
		fmt.Println("error starting server: ", err)
		return err
	}

	return nil
}
