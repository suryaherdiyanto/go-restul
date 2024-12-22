package model

import (
	"database/sql"
)

type User struct {
	Id        int
	FirstName string
	LastName  sql.NullString
	Password  string
	Email     string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
