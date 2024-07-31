package service

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type TokenService struct {
	cfg *config.Config
}

func newTokenService(cfg *config.Config) *TokenService {

	return &TokenService{
		cfg: cfg,
	}
}

func (t *TokenService) CreateToken(email, password string) entity.Token {

	jwt := t.cfg.TokenAuth
	jwtExpiration := t.cfg.JWTExpirationTime

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": email,
		"exp": jwtExpiration,
	})

	return entity.Token{
		Token:      token,
		Expiration: jwtExpiration,
	}
}
