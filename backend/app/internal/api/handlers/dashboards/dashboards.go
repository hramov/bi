package dashboards

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/internal/repository"
	"net/http"
)

type Handler struct {
	Repository repository.DashboardsRepository
}

func New(db *sql.DB) *Handler {
	repo := repository.DashboardsRepository{
		Db: db,
	}
	return &Handler{Repository: repo}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", h.get)
	r.Get("/types", h.getAvailableTypes)
	r.Get("/{id}", h.getById)
	r.Post("/", h.create)
	r.Put("/", h.update)
	r.Delete("/", h.delete)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	data, err := h.Repository.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rawData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(rawData)
}

func (h *Handler) getAvailableTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.Repository.GetAvailableTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rawData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	w.Write(rawData)
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	data, err := h.Repository.GetByDashId(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	rawData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(rawData)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
}
