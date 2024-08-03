package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type Controller struct {
	TokenController *TokenController
	UserController  *UserController
}

func NewController(cfg *config.Config, service *service.Service) *Controller {
	controller := new(Controller)

	controller.TokenController = newTokenController(cfg, service)
	controller.UserController = newUserController(cfg, service)

	return controller
}
