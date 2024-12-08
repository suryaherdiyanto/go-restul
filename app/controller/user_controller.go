package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	users := c.UserRepository.All()
	dec := json.NewEncoder(w)

	res := response.NewSuccessResponse(http.StatusText(http.StatusOK), response.NewUsersResponse(users))
	w.Header().Add("Content-Type", "application/json")
	dec.Encode(res)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")

	if err != err {
		response.HandleNotFound(w, enc, "Resource not found!")
		return
	}

	user, ok := c.UserRepository.FindById(id)

	if !ok {
		response.HandleNotFound(w, enc, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(http.StatusText(http.StatusOK), response.NewUserResponse(user))
	enc.Encode(res)
}
