package service

import (
	"errors"
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/data"
	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type UserService struct {
	cfg  *config.Config
	data *data.Data
}

func newUserService(cfg *config.Config, data *data.Data) *UserService {
	return &UserService{
		cfg:  cfg,
		data: data,
	}
}

func (s *UserService) CreateUser(user *entity.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("invalid user data")
	}
	if _, err := s.FindByEmail(user.Email); err == nil {
		return errors.New("user already exists")
	}

	return s.data.UserGormDB.Create(user).Error
}

func (s *UserService) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	err := s.data.UserGormDB.Where("email = ?", email).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, err
}

func (s *UserService) Delete(id string) error {
	return s.data.UserGormDB.Where("id = ?", id).Delete(&entity.User{}).Error
}
