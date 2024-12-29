package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-restful/app/response"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
)

type key string

const UserKey key = "auth_user"

func CheckAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response.JsonResponse(w, response.NewUnAuthorizedResponse("Authorization header missing"))
			return
		}

		tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", 1)
		claims, err := token.ValidateToken(tokenHeader, "thesecrettoken")

		if err != nil {
			response.JsonResponse(w, response.NewUnAuthorizedResponse(err.Error()))
			return
		}
		ctx := context.WithValue(r.Context(), UserKey, claims)

		next(w, r.WithContext(ctx), ps)
	}
}
