package service

import "github.com.br/GregoryLacerda/AMSVault/config"

type Service struct {
	cfg          *config.Config
	TokenService *TokenService
}

func NewService(cfg *config.Config) *Service {
	service := new(Service)

	service.cfg = cfg
	service.TokenService = newTokenService(cfg)

	return service
}
