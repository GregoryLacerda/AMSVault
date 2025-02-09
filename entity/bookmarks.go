package entity

import "time"

type Bookmarks struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user"`
	Story     Story     `json:"story"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
