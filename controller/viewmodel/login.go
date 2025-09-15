package viewmodel

import "github.com.br/GregoryLacerda/AMSVault/entity"

type LoginRequestViewModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseViewModel struct {
	Token      string `json:"acces_token"`
	Expiration int    `json:"expiration"`
}

func MapLoginResponseToViewModel(token entity.Token) LoginResponseViewModel {
	return LoginResponseViewModel{
		Token:      token.Token,
		Expiration: token.Expiration,
	}
}
