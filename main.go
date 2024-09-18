package main

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/integration"
	"github.com.br/GregoryLacerda/AMSVault/server"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	data, err := data.New(cfg)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	integrations := integration.NewIntegration(cfg)

	service := service.NewService(cfg, data, integrations)

	ctrl := controller.NewController(cfg, service)

	srv := server.New()

	srv.Start(cfg, ctrl, nil)

}
