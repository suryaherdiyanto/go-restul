package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
