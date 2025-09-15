package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
)

type UserDB struct {
	DB *sql.DB
}

func newUserDB(DB *sql.DB) *UserDB {
	return &UserDB{
		DB: DB,
	}
}

func (m *UserDB) Insert(user entity.User) error {

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	if _, err := m.DB.Exec(query, user.Name, user.Email, user.Password); err != nil {
		return errors.NewDatabaseError("Insert", err)
	}
	return nil
}

func (m *UserDB) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := m.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return user, errors.NewDatabaseError("FindByEmail", fmt.Errorf("user not found with email %s", email))
		}

		return user, errors.NewDatabaseError("FindByEmail", err)
	}
	return user, nil
}

func (m *UserDB) FindByID(ID int64) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	err := m.DB.QueryRow(query, ID).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return user, errors.NewDatabaseError("FindByID", fmt.Errorf("user not found with id %d", ID))
		}
		return user, errors.NewDatabaseError("FindByID", err)
	}
	return user, nil
}

func (m *UserDB) Update(user entity.User) error {

	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	if _, err := m.DB.Exec(query, user.Name, user.Email, user.Password, user.ID); err != nil {
		return errors.NewDatabaseError("Update", err)
	}

	return nil
}

func (m *UserDB) Delete(ID int64) error {

	query := "DELETE FROM users WHERE id = ?"
	if _, err := m.DB.Exec(query, ID); err != nil {
		return errors.NewDatabaseError("Delete", err)
	}

	return nil
}
