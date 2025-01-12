package middleware

import (
	"context"
	"net/http"
	"time"
)

func DBMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		h.ServeHTTP(w, r.WithContext(ctx))
	})

}
