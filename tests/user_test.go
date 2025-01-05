package tests

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/response"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func TestIndex(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	userService.Create(context.Background(), &request.UserRequest{FirstName: "lala", LastName: "move", Email: "lalamove@gmail.com"})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	userController.Index(w, req, httprouter.Params{})

	if w.Code != 200 {
		t.Errorf("Expected 200, but got %d", w.Code)
	}

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	if res.GetData() == nil {
		t.Errorf("The data shouldn't be nil, data: %v", res.GetData())
	}
}

func TestShow(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	r := &request.UserRequest{FirstName: faker.FirstName(), LastName: faker.LastName(), Email: faker.Email()}
	user := userService.Create(context.Background(), r)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%d", user.Id), nil)
	params := httprouter.Param{Key: "id", Value: strconv.Itoa(user.Id)}
	userController.Show(w, req, httprouter.Params{params})

	if w.Code != 200 {
		t.Errorf("Expected 200, but got %d", w.Code)
	}

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	data := res.GetData().(map[string]interface{})
	if e, _ := data["email"]; e != user.Email {
		t.Errorf("Expected %s, but got %s", user.Email, e)
	}
}
