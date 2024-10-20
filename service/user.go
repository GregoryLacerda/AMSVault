package service

import (
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type UserService struct {
	data *data.Data
}

func newUserService(data *data.Data) *UserService {
	return &UserService{
		data: data,
	}
}

func (s *UserService) CreateUser(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	if _, err := s.FindByEmail(user.Email); err == nil {
		return errors.New("already exists a user with this email")
	}

	return s.data.UserDB.Create(user)
}

func (s *UserService) FindByID(id string) (*entity.User, error) {
	return s.data.UserDB.FindByID(id)
}

func (s *UserService) FindByEmail(email string) (user *entity.User, err error) {
	return s.data.UserDB.FindByEmail(email)
}

func (s *UserService) Delete(id string) error {
	return s.data.UserDB.Delete(id)
}

func (s *UserService) Update(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, err := s.FindByID(user.ID.String()); err != nil {
		return errors.New("user not found")
	}

	return s.data.UserDB.Update(user)
}

func (s *UserService) FindById(id string) (*entity.User, error) {
	return s.data.UserDB.FindByID(id)
}
