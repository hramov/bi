package users

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/repository"
	"net/http"
)

type Handler struct {
	Repository repository.UsersRepository
}

func New(db *sql.DB) *Handler {
	repo := repository.UsersRepository{
		Db: db,
	}
	return &Handler{Repository: repo}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/login", h.login)
	r.Post("/check-access", h.checkAccess)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) checkAccess(w http.ResponseWriter, r *http.Request) {
}
