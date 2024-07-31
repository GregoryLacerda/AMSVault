package viewmodel

import "github.com.br/GregoryLacerda/AMSVault/entity"

type TokenRequestViewModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponseViewModel struct {
	Token      string `json:"acces_token"`
	Expiration int    `json:"expiration"`
}

func MapTokenResponseToViewModel(token entity.Token) TokenResponseViewModel {
	return TokenResponseViewModel{
		Token:      token.Token,
		Expiration: token.Expiration,
	}
}
