package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-restful/app/model"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
)

func TestNoAuthrizationHeader(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	handler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}
	authCheckHandler := CheckAuth(handler)

	authCheckHandler(w, r, httprouter.Params{})

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected: %d, but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestWrongToken(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	r.Header.Add("Authorization", "Bearer thewrongtoken")
	handler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}
	authCheckHandler := CheckAuth(handler)

	authCheckHandler(w, r, httprouter.Params{})

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected: %d, but got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestSuccessValidateToken(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	var claims *token.UserClaims

	jwt, _ := token.GenerateToken(&model.User{Id: 1, FirstName: "john", Email: "johndoe@gmail.com"}, os.Getenv("JWT_SECRET"), time.Hour)
	r.Header.Add("Authorization", "Bearer "+jwt)
	handler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		claims = r.Context().Value(UserKey).(*token.UserClaims)
	}
	authCheckHandler := CheckAuth(handler)

	authCheckHandler(w, r, httprouter.Params{})

	if claims.Email != "johndoe@gmail.com" {
		t.Errorf("Expected johndoe@gmail.com, but got %v", claims)
	}
}
