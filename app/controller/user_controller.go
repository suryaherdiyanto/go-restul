package controller

import (
	"encoding/json"
	"net/http"

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
