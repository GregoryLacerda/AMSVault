package service

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/integration"
)

type Service struct {
	cfg          *config.Config
	data         *data.Data
	TokenService *TokenService
	UserService  *UserService
	StoryService *StoryService
}

func NewService(cfg *config.Config, data *data.Data, Integrations *integration.Integrations) *Service {
	service := new(Service)

	service.cfg = cfg
	service.data = data
	service.TokenService = newTokenService(cfg, data)
	service.UserService = newUserService(data)
	service.StoryService = newStoryService(data, Integrations)

	return service
}
