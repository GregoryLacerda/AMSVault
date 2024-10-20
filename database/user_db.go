package database

import (
	"errors"
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (user *entity.User, err error) {

	err = u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (u *User) Delete(id string) error {
	return u.DB.Where("id = ?", id).Delete(&entity.User{}).Error
}

func (u *User) FindByID(id string) (user *entity.User, err error) {
	err = u.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (u *User) Update(user *entity.User) error {
	return u.DB.Save(user).Error
}
