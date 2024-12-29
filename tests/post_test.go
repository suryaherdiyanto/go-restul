package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/go-restful/app/middleware"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/response"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
)

func TestCreatePost(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	postRequest := &request.PostRequest{
		Title:    faker.Sentence(),
		Category: faker.Word(),
		Content:  faker.Paragraph(),
	}
	postJson, _ := json.Marshal(postRequest)
	claims, err := token.ValidateToken(accessToken, os.Getenv("JWT_SECRET"))

	if err != nil {
		t.Errorf("Invalid token: %v", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/posts", strings.NewReader(string(postJson)))

	ctx := context.WithValue(r.Context(), middleware.UserKey, claims)
	postController.Store(w, r.WithContext(ctx), httprouter.Params{})

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}

}

func TestIndexPosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	claims, err := token.ValidateToken(accessToken, os.Getenv("JWT_SECRET"))

	if err != nil {
		t.Errorf("Invalid token: %v", err)
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/posts", nil)

	ctx := context.WithValue(r.Context(), middleware.UserKey, claims)
	postController.Index(w, r.WithContext(ctx), httprouter.Params{})

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestCorrectNumberOfPosts(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)

	claims, err := token.ValidateToken(accessToken, os.Getenv("JWT_SECRET"))

	if err != nil {
		t.Errorf("Invalid token: %v", err)
	}

	postService.Create(context.Background(), authUser.Id, &request.PostRequest{
		Title:    faker.Sentence(),
		Category: faker.Word(),
		Content:  faker.Paragraph(),
	})
	postService.Create(context.Background(), authUser.Id, &request.PostRequest{
		Title:    faker.Sentence(),
		Category: faker.Word(),
		Content:  faker.Paragraph(),
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/posts", nil)

	ctx := context.WithValue(r.Context(), middleware.UserKey, claims)
	postController.Index(w, r.WithContext(ctx), httprouter.Params{})

	res := &response.SuccessResponse{}
	response.ParseSuccessResponse(w.Body, res)

	if res.GetData() == nil {
		t.Errorf("Response data shouldn't be nil")
	}

	data := res.GetData().([]interface{})
	if len(data) < 2 {
		t.Errorf("Expected number of posts 2, but got %d", len(data))
	}
}
