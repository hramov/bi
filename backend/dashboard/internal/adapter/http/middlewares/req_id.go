package middlewares

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func ReqId(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		w.Header().Set("x-request-id", id.String())
		h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "x-request-id", id)))
	})
}
