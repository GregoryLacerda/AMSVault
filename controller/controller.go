package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type Controller struct {
	Token *TokenController
}

func NewController(cfg *config.Config, service *service.Service) *Controller {
	controller := new(Controller)

	controller.Token = newTokenController(cfg, service)

	return controller
}
