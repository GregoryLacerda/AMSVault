package integration

import "github.com.br/GregoryLacerda/AMSVault/config"

type Integrations struct {
	MALIntegration *MALIntegration
}

func NewIntegration(cfg *config.Config) *Integrations {
	return &Integrations{
		MALIntegration: newMALIntegration(cfg),
	}
}
