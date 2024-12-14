package controller

import (
	"context"
	"fmt"
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

	res := response.NewSuccessResponse(response.NewUsersDataResponse(&users))
	response.JsonResponse(w, res)
	fmt.Printf("%v \n", users)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != err {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	user, ok := c.UserRepository.FindById(ctx, id)

	if !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(response.NewUserDataResponse(&user))
	response.JsonResponse(w, res)
}
