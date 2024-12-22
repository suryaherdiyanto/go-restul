package model

import "database/sql"

type Post struct {
	Id        int
	Title     string
	Category  string
	Content   string
	UserID    int
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
