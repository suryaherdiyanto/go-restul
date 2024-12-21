package model

import "database/sql"

type Post struct {
	Id        int          `json:"id"`
	Title     string       `json:"title"`
	Category  string       `json:"category"`
	Content   string       `json:"content"`
	UserID    int          `json:"user_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
