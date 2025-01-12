package controller

import (
	"net/http"
	"strconv"

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
	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	posts := c.PostRepository.FilterBy(r.Context(), "user_id", user.Id)

	res := response.NewSuccessResponse(resource.NewPostsResource(&posts))
	response.JsonResponse(w, res)
}

func (c *PostController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	post, ok := c.PostRepository.FindById(r.Context(), id)

	if !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	res := response.NewSuccessResponse(resource.NewPostResource(&post))
	response.JsonResponse(w, res)
}

func (c *PostController) Store(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	postRequest, err := request.NewPostRequest(r.Body)
	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(postRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	post := c.PostRepository.Create(r.Context(), user.Id, postRequest)

	response.JsonResponse(w, response.NewCreatedResponse("Post Created!", resource.NewPostResource(&post)))
}

func (c *PostController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)

	user := r.Context().Value(middleware.UserKey).(*token.UserClaims)

	postRequest, err := request.NewPostUpdateRequest(r.Body)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(postRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	post := c.PostRepository.Update(r.Context(), id, user.Id, postRequest)

	response.JsonResponse(w, response.NewSuccessResponse(resource.NewPostResource(&post)))

}

func (c *PostController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	helper.ErrorPanic(err)

	if _, ok := c.PostRepository.FindById(r.Context(), id); !ok {
		response.HandleNotFound(w, "Resource not found!")
		return
	}

	c.PostRepository.Delete(r.Context(), id)

	response.JsonResponse(w, response.NewSuccessResponse(nil))

}
