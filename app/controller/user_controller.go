package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/response"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepository,
	}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	users := c.UserRepository.All(ctx)
	dec := json.NewEncoder(w)

	res := response.NewSuccessResponse(http.StatusText(http.StatusOK), response.NewUsersResponse(users))
	w.Header().Add("Content-Type", "application/json")
	dec.Encode(res)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(ps.ByName("id"))
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")

	if err != err {
		response.HandleNotFound(w, enc, "Resource not found!")
		return
	}

	user, ok := c.UserRepository.FindById(ctx, id)

	if !ok {
		response.HandleNotFound(w, enc, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(http.StatusText(http.StatusOK), response.NewUserResponse(user))
	enc.Encode(res)
}
