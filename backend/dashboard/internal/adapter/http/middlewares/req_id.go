package middlewares

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func ReqId(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New()
		now := time.Now()
		w.Header().Set("x-request-id", id.String())
		h.ServeHTTP(w, r.WithContext(context.WithValue(context.WithValue(r.Context(), "x-request-id", id), "start", now)))
	})
}
