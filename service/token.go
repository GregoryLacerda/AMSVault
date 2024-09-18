package service

import (
	"fmt"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type TokenService struct {
	cfg  *config.Config
	data *data.Data
}

func newTokenService(cfg *config.Config, data *data.Data) *TokenService {

	return &TokenService{
		cfg:  cfg,
		data: data,
	}
}

func (t *TokenService) CreateToken(email, password string) entity.Token {

	userService := newUserService(t.data)

	user, err := userService.FindByEmail(email)
	if err != nil {
		return entity.Token{}
	}

	jwt := t.cfg.TokenAuth
	jwtExpiration := t.cfg.JWTExpirationTime

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID,
		"exp": jwtExpiration,
	})

	return entity.Token{
		Token:      token,
		Expiration: jwtExpiration,
	}
}

func (t *TokenService) GetUserIdFromToken(token string) string {
	jwt := t.cfg.TokenAuth

	claims, _ := jwt.Decode(token)

	userId, _ := claims.Get("sub")

	//TODO: remove this
	fmt.Println(userId)
	fmt.Println(userId.(string))

	return userId.(string)
}
