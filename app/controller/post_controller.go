package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-restful/app/middleware"
	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/resource"
	"github.com/go-restful/app/response"
	"github.com/go-restful/helper"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostRepository repository.PostRepository
}

func NewPostController(postRepository repository.PostRepository) *PostController {
	return &PostController{postRepository}
}

func (c *PostController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	posts := c.PostRepository.FilterBy(ctx, "user_id", user.Id)

	res := response.NewSuccessResponse(resource.NewPostsResource(&posts))
	response.JsonResponse(w, res)
}

func (c *PostController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	post, ok := c.PostRepository.FindById(ctx, id)

	if !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(resource.NewPostResource(&post))
	response.JsonResponse(w, res)
}

func (c *PostController) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	postRequest, err := request.NewPostRequest(r.Body)
	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(postRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	post := c.PostRepository.Create(ctx, user.Id, postRequest)

	response.JsonResponse(w, response.NewCreatedResponse("Post Created!", resource.NewPostResource(&post)))
}

func (c *PostController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()
	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	postRequest, err := request.NewPostUpdateRequest(r.Body)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(postRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	post := c.PostRepository.Update(ctx, id, user.Id, postRequest)

	response.JsonResponse(w, response.NewSuccessResponse(resource.NewPostResource(&post)))

}

func (c *PostController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)

	if _, ok := c.PostRepository.FindById(r.Context(), id); !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	c.PostRepository.Delete(ctx, id)

	response.JsonResponse(w, response.NewSuccessResponse(nil))

}
