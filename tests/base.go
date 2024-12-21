package tests

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func databaseUp(db *sql.DB, dbname string) error {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://../db/migrations", dbname, driver)

	if err != nil {
		return err
	}

	return m.Up()
}

func databaseDown(db *sql.DB, dbname string) error {
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://../db/migrations", dbname, driver)

	if err != nil {
		return err
	}

	return m.Drop()
}
