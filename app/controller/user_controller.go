package controller

import (
	"context"
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
	users := c.UserRepository.All(r.Context())

	res := response.NewSuccessResponse(resource.NewUsersResource(&users))
	response.JsonResponse(w, res)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	user, ok := c.UserRepository.FindById(r.Context(), id)

	if !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(resource.NewUserResource(&user))
	response.JsonResponse(w, res)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)
	r.Header.Set("Content-Type", "application/json")

	userRequest, err := request.NewUserUpdateRequest(r.Body)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(userRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	user := c.UserRepository.Update(r.Context(), id, userRequest)

	response.JsonResponse(w, response.NewSuccessResponse(resource.NewUserResource(&user)))

}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)
	r.Header.Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	c.UserRepository.Delete(ctx, id)

	response.JsonResponse(w, response.NewSuccessResponse(nil))

}
