package mysql

import (
	"database/sql"
)

type Mysql struct {
	db      *sql.DB
	StoryDB StoryDB
	UserDB  UserDB
}

func NewMysql(DB *sql.DB) *Mysql {
	return &Mysql{
		db:      DB,
		StoryDB: *newStoryDB(DB),
		UserDB:  *newUserDB(DB),
	}
}

func (m *Mysql) Close() error {
	return m.db.Close()
}
