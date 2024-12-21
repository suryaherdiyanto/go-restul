package tests

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-restful/app/controller"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/response"
	"github.com/go-restful/app/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gorestful_test")
	err := databaseUp(db, "gorestful_test")

	if err != nil {
		t.Error(err)
	}

	userService := service.NewUserService(db)
	userService.Create(context.Background(), &request.UserRequest{FirstName: "lala", LastName: "move", Email: "lalamove@gmail.com"})
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	router.GET("/api/users", userController.Index)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	router.ServeHTTP(w, req)
	databaseDown(db, "gorestful_test")

	if w.Code != 200 {
		t.Errorf("Expected 200, but got %d", w.Code)
	}

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	if res.GetData() == nil {
		t.Errorf("The data shouldn't be nil, data: %v", res.GetData())
	}
}
