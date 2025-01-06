package tests

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/model"
	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/service"
	"github.com/go-restful/token"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB
var userService repository.UserRepository
var postService repository.PostRepository
var userController *controller.UserController
var authController *controller.AuthController
var postController *controller.PostController
var accessToken string
var authUser model.User

func setupTest(tb testing.TB) func(tb testing.TB) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_TEST_URL"))
	if err != nil {
		tb.Error(err)
	}

	err = databaseUp(db, os.Getenv("DATABASE_TEST_NAME"))

	if err != nil {
		tb.Error(err)
	}

	userService = service.NewUserService(db)
	postService = service.NewPostService(db)
	authUser = userService.Create(context.Background(), &request.UserRequest{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  "password",
	})
	accessToken, _ = token.GenerateToken(&authUser, os.Getenv("JWT_SECRET"), time.Hour)
	userController = controller.NewUserController(userService)
	authController = controller.NewAuthController(userService)
	postController = controller.NewPostController(postService)

	return func(tb testing.TB) {
		databaseDown(db, os.Getenv("DATABASE_TEST_NAME"))
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
