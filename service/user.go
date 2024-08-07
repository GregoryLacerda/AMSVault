package service

import (
	"errors"
	"strings"

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

	return s.data.UserGormDB.Create(user).Error
}

func (s *UserService) FindByID(id string) (*entity.User, error) {
	user := new(entity.User)
	err := s.data.UserGormDB.Where("id = ?", id).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, err
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

func (s *UserService) Update(user *entity.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if _, err := s.FindByID(user.ID.String()); err != nil {
		return errors.New("user not found")
	}

	return s.data.UserGormDB.Save(user).Error
}

func (s *UserService) FindById(id string) (*entity.User, error) {
	user := new(entity.User)
	err := s.data.UserGormDB.Where("id = ?", id).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
