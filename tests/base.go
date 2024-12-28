package tests

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/response"
	"github.com/go-restful/app/router"
	"github.com/go-restful/app/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var routes *httprouter.Router
var userService repository.UserRepository

func setupTest(tb testing.TB) func(tb testing.TB) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gorestful_test")
	if err != nil {
		tb.Error(err)
	}

	routes = router.NewRouter(db)
	routes.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		response.JsonResponse(w, response.NewInternalServerError("Something went wrong!", err))
	}
	err = databaseUp(db, "gorestful_test")

	if err != nil {
		tb.Error(err)
	}

	userService = service.NewUserService(db)

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
