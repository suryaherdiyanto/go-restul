package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/resource"
	"github.com/go-restful/app/response"
	"github.com/go-restful/helper"
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

	res := response.NewSuccessResponse(resource.NewUsersResource(&users))
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

	res := response.NewSuccessResponse(resource.NewUserResource(&user))
	response.JsonResponse(w, res)
}

func (c *UserController) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	userRequest := &request.UserRequest{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(userRequest)

	helper.ErrorPanic(err)

	user := c.UserRepository.Create(ctx, userRequest)

	response.JsonResponse(w, response.NewCreatedResponse("User Created!", resource.NewUserResource(&user)))
}
