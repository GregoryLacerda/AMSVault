package main

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/server"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	service := service.NewService(cfg)

	ctrl := controller.NewController(cfg, service)

	srv := server.New()

	srv.Start(cfg, ctrl, nil)

}
