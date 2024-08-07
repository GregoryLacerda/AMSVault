package service

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/data"
)

type Service struct {
	cfg          *config.Config
	data         *data.Data
	TokenService *TokenService
	UserService  *UserService
	AnimeService *AnimeService
}

func NewService(cfg *config.Config, data *data.Data) *Service {
	service := new(Service)

	service.cfg = cfg
	service.data = data
	service.TokenService = newTokenService(cfg)
	service.UserService = newUserService(data)
	service.AnimeService = newAnimeService(data)

	return service
}
