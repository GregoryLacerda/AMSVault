package entity

import (
	"time"

	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalError("NewUser", err)
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ValidateLogin(email, password string) bool {
	return u.Email == email && u.ValidatePassword(password)
}

func (u *User) Validate() error {
	if u.Name != "" && u.Email != "" && u.Password != "" {
		return nil
	}
	return errors.NewValidationError("invalid user data")
}
