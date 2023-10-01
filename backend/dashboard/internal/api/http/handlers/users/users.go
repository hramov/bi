package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal"
	"net/http"
)

type Handler struct {
	repo   Repository
	logger internal.Logger
}

type Repository interface {
}

func New(repo Repository, logger internal.Logger) *Handler {
	return &Handler{repo: repo, logger: logger}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/login", h.login)
	r.Post("/check-access", h.checkAccess)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) checkAccess(w http.ResponseWriter, r *http.Request) {
}
