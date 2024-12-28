package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-restful/app/repository"
	"github.com/go-restful/app/request"
	"github.com/go-restful/app/resource"
	"github.com/go-restful/app/response"
	"github.com/go-restful/helper"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository repository.UserRepository
}

func NewAuthController(repo repository.UserRepository) *AuthController {
	return &AuthController{
		UserRepository: repo,
	}
}

func (c AuthController) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	userRequest, err := request.NewUserRequest(r.Body)

	helper.ErrorPanic(err)

	validate, ok := request.Validate(userRequest)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", validate.Map()))
		return
	}

	if _, ok = c.UserRepository.FindBy(ctx, "email", userRequest.Email); ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Validation Error!", &map[string][]interface{}{
			"email": {fmt.Sprintf("The email: %s, is already been taken", userRequest.Email)},
		}))
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)

	helper.ErrorPanic(err)
	userRequest.Password = string(hashPassword)

	user := c.UserRepository.Create(ctx, userRequest)

	response.JsonResponse(w, response.NewCreatedResponse("Register Successfully!", resource.NewUserResource(&user)))
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	userRequest, err := request.NewLoginRequest(r.Body)

	helper.ErrorPanic(err)

	user, ok := c.UserRepository.FindBy(ctx, "email", userRequest.Email)

	if !ok {
		response.JsonResponse(w, response.NewBadRequestResponse("Invalid email or password!", nil))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		response.JsonResponse(w, response.NewBadRequestResponse("Invalid email or password!", nil))
		return
	}

	authToken, err := token.GenerateToken(&user, "thesecrettoken", time.Hour*24)
	helper.ErrorPanic(err)

	refreshToken, err := token.GenerateToken(&user, "thesecrettoken", time.Hour*24*7)
	helper.ErrorPanic(err)

	cookie := &http.Cookie{
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

	response.JsonResponse(w, response.NewSuccessResponse(&map[string]interface{}{
		"user":  resource.NewUserResource(&user),
		"token": authToken,
	}))
}
