package controller

import (
	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/controller/viewmodel"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/service"
)

type UserController struct {
	cfg         *config.Config
	UserService *service.UserService
}

func newUserController(cfg *config.Config, service *service.Service) *UserController {
	return &UserController{
		cfg:         cfg,
		UserService: service.UserService,
	}
}

func (u *UserController) CreateUser(name string, email string, password string) error {
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return err
	}

	return u.UserService.CreateUser(user)
}

func (u *UserController) FindByEmail(email string) (viewmodel.UserResponseViewModel, error) {

	user, err := u.UserService.FindByEmail(email)
	if err != nil {
		return viewmodel.UserResponseViewModel{}, err
	}

	return viewmodel.MapUserResponseToViewModel(user), nil
}

func (u *UserController) Delete(id string) error {
	return u.UserService.Delete(id)
}

func (u *UserController) Update(user *entity.User) error {
	return u.UserService.Update(user)
}

func (u *UserController) FindById(id string) (viewmodel.UserResponseViewModel, error) {
	user, err := u.UserService.FindById(id)
	if err != nil {
		return viewmodel.UserResponseViewModel{}, err
	}

	return viewmodel.MapUserResponseToViewModel(user), nil
}
