package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

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
