package dashboards

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/internal/repository"
	"github.com/hramov/gvc-bi/backend/pkg/utils"
	"net/http"
	"strconv"
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
	r.Get("/item/{id}", h.getItemById)
	r.Post("/item", h.createItem)
	r.Put("/item/{id}", h.updateItem)
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

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	data, err := h.Repository.GetItemById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	utils.SendResponse(http.StatusOK, data, w)
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[repository.Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.Repository.CreateItem(body)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot save data to database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	body, err := utils.GetBody[repository.Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	updatedId, err := h.Repository.UpdateItem(body, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, updatedId, w)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
}
