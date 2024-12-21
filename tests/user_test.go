package tests

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/response"
	"github.com/go-restful/app/router"
	"github.com/go-restful/app/service"
	_ "github.com/go-sql-driver/mysql"
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
	err = databaseUp(db, "gorestful_test")

	if err != nil {
		tb.Error(err)
	}

	userService = service.NewUserService(db)

	return func(tb testing.TB) {
		databaseDown(db, "gorestful_test")
	}
}

func TestIndex(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	userService.Create(context.Background(), &request.UserRequest{FirstName: "lala", LastName: "move", Email: "lalamove@gmail.com"})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	routes.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200, but got %d", w.Code)
	}

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	if res.GetData() == nil {
		t.Errorf("The data shouldn't be nil, data: %v", res.GetData())
	}
}
