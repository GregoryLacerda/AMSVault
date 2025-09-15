package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type TokenController struct {
	cfg          *config.Config
	TokenService *service.TokenService
}

func newTokenController(cfg *config.Config, service *service.Service) *TokenController {
	return &TokenController{
		cfg:          cfg,
		TokenService: service.TokenService,
	}
}

func (t *TokenController) CreateToken(email, password string) (viewmodel.LoginResponseViewModel, error) {

	token, err := t.TokenService.CreateToken(email, password)
	if err != nil {
		return viewmodel.LoginResponseViewModel{}, err
	}

	return viewmodel.MapLoginResponseToViewModel(token), nil
}
