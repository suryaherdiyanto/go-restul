package middleware

import (
	"net/http"
	"strings"

	"github.com/go-restful/app/response"
	"github.com/go-restful/token"
	"github.com/julienschmidt/httprouter"
)

func CheckAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response.JsonResponse(w, response.NewUnAuthorizedResponse("Authorization header missing"))
			return
		}

		tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", 1)
		_, err := token.ValidateToken(tokenHeader, "thesecrettoken")

		if err != nil {
			response.JsonResponse(w, response.NewUnAuthorizedResponse(err.Error()))
			return
		}

		next(w, r, ps)
	}
}
