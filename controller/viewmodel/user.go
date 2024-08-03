package viewmodel

import "github.com.br/GregoryLacerda/AMSVault/entity"

type UserRequestViewModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseViewModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MapUserResponseToViewModel(user *entity.User) UserResponseViewModel {
	return UserResponseViewModel{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}
