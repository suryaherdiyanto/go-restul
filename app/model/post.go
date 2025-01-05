package model

import "database/sql"

type Post struct {
	Id        int          `db:"id"`
	Title     string       `db:"title"`
	Category  string       `db:"category"`
	Content   string       `db:"content"`
	UserID    int          `db:"user_id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
