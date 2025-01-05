package model

import (
	"database/sql"
)

type User struct {
	Id        int            `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Password  string         `db:"password"`
	Email     string         `db:"email"`
	CreatedAt sql.NullTime   `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}
