package mysql

import (
	"database/sql"
)

type Mysql struct {
	db      *sql.DB
	storyDB StoryDB
	userDb  UserDB
}

func NewMysql(DB *sql.DB) *Mysql {
	return &Mysql{
		db:      DB,
		storyDB: *newStoryDB(DB),
		userDb:  *newUserDB(DB),
	}
}

func (m *Mysql) Close() error {
	return m.db.Close()
}
