package mysql

import (
	"database/sql"

	"github.com.br/GregoryLacerda/AMSVault/entity"
)

type Mysql struct {
	db *sql.DB
}

func NewMysql(DB *sql.DB) *Mysql {
	return &Mysql{
		db: DB,
	}
}

func (m *Mysql) Close() error {
	return m.db.Close()
}

func (m *Mysql) Insert(user entity.User) error {
	_, err := m.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
