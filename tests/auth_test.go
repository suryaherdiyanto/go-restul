package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-restful/app/request"
	"github.com/go-restful/app/response"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func TestSuccessRegister(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	userRequest := &request.UserRequest{
		FirstName:            "lala",
		LastName:             "move",
		Email:                "lalamove@example.com",
		Password:             "password123456",
		PasswordConfirmation: "password123456",
	}
	userJson, _ := json.Marshal(userRequest)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(userJson)))
	routes.ServeHTTP(w, req)

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	if w.Code != 201 {
		t.Errorf("Expected 201, but got %d", w.Code)
	}

	if res.GetData() == nil {
		t.Errorf("The data shouldn't be nil, data: %v", res.GetData())
	}

	data := res.GetData().(map[string]interface{})
	if email, _ := data["email"]; email != userRequest.Email {
		t.Errorf("Expected %s, but got %s", userRequest.Email, email)
	}

}

func TestFailedRegister(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	userRequest := &request.UserRequest{
		FirstName: "lala",
		LastName:  "move",
		Email:     "",
	}

	userJson, _ := json.Marshal(userRequest)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(userJson)))
	routes.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected 400, but got %d", w.Code)
	}
}

func TestFailedRegisterIfEmailAlreadyExist(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	userRequest := &request.UserRequest{
		FirstName:            "john",
		LastName:             "mayer",
		Email:                "johnmayer@gmail.com",
		Password:             "password123456",
		PasswordConfirmation: "password123456",
	}
	userService.Create(context.Background(), userRequest)

	userJson, _ := json.Marshal(userRequest)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", strings.NewReader(string(userJson)))
	routes.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Expected 400, but got %d", w.Code)
	}

}

func TestSuccessLogin(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	password, _ := bcrypt.GenerateFromPassword([]byte("password123456"), bcrypt.DefaultCost)
	userService.Create(context.Background(), &request.UserRequest{
		FirstName:            "manju",
		LastName:             "",
		Email:                "manju@gmail.com",
		Password:             string(password),
		PasswordConfirmation: "",
	})
	loginRequest := &request.LoginRequest{
		Email:    "manju@gmail.com",
		Password: "password123456",
	}
	loginJson, _ := json.Marshal(loginRequest)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(string(loginJson)))
	routes.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected 200, but got %d", w.Code)
		fmt.Printf("Response: %v\n", w.Body.String())
	}

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)
	data := res.GetData().(map[string]interface{})

	if _, ok := data["token"]; !ok {
		t.Errorf("Token key should be exist")
	}

}
