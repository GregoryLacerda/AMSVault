package main

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller"
	"github.com.br/GregoryLacerda/AMSVault/server"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctrl := controller.New()

	srv := server.New()

	srv.Start(cfg, ctrl, nil)

}
