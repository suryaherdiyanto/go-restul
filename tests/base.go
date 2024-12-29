package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB
var userService repository.UserRepository
var userController *controller.UserController
var authController *controller.AuthController

func setupTest(tb testing.TB) func(tb testing.TB) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gorestful_test")
	if err != nil {
		tb.Error(err)
	}

	err = databaseUp(db, "gorestful_test")

	if err != nil {
		tb.Error(err)
	}

	userService = service.NewUserService(db)
	userService.Create(context.Background(), &request.UserRequest{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  "password",
	})
	userController = controller.NewUserController(userService)
	authController = controller.NewAuthController(userService)

	return func(tb testing.TB) {
		databaseDown(db, "gorestful_test")
	}
}

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
