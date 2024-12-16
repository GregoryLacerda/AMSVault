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

	return s.data.Mysql.UserDB.Insert(*user)
}

func (s *UserService) FindByID(id int64) (entity.User, error) {
	return s.data.Mysql.UserDB.FindByID(id)
}

func (s *UserService) FindByEmail(email string) (user entity.User, err error) {
	return s.data.Mysql.UserDB.FindByEmail(email)
}

func (s *UserService) Delete(ID int64) error {
	return s.data.Mysql.UserDB.Delete(ID)
}

func (s *UserService) Update(user entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, err := s.FindByID(user.ID); err != nil {
		return errors.New("user not found")
	}

	return s.data.Mysql.UserDB.Update(user)
}

func (s *UserService) FindById(ID int64) (entity.User, error) {
	return s.data.Mysql.UserDB.FindByID(ID)
}
