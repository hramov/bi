package dashboards

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/repository"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

type Model struct {
	Id          int       `json:"id"`
	DashId      string    `json:"dash_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`

	Items []*Item `json:"items"`
}

type Item struct {
	Id          int    `json:"id"`
	DashId      string `json:"dash_id"`
	ItemType    int    `json:"item_type"`
	Position    string `json:"position"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DataQueries any    `json:"data_queries"`
	Options     any    `json:"raw_options"`
	//CreatedAt   time.Time `json:"created_at"`
	//UpdatedAt   time.Time `json:"updated_at"`
	//DeletedAt   time.Time `json:"deleted_at"`
}

type Handler struct {
	repo   Repository
	logger internal.Logger
}

type Repository interface {
	Get() ([]*repository.Model, error)
	GetById(id int) (*repository.Model, error)
	GetItemById(id int) (*repository.Item, error)
	GetAvailableTypes() ([]*repository.TypeModel, error)
	GetByDashId(id string) (*repository.Model, error)
	Create(dto repository.Model) (*int, error)
	Update(dto repository.Model, id int) (*int, error)
	CreateItem(dto repository.Item) (*int, error)
	UpdateItem(dto repository.Item, id int) (*int, error)
}

func New(repo Repository, logger internal.Logger) *Handler {
	return &Handler{repo: repo, logger: logger}
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
	data, err := h.repo.Get()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var response []*Model

	for _, d := range data {
		response = append(response, &Model{
			Id:          d.Id,
			DashId:      d.DashId,
			Title:       d.Title,
			Description: d.Description.String,
			CreatedAt:   d.CreatedAt,
			UpdatedAt:   d.UpdatedAt.Time,
			DeletedAt:   d.DeletedAt.Time,
		})
	}

	utils.SendResponse(http.StatusOK, response, w)
}

func (h *Handler) getAvailableTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetAvailableTypes()
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
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendError(http.StatusInternalServerError, "no id found", w)
		return
	}

	data, err := h.repo.GetByDashId(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var items []*Item

	err = json.Unmarshal([]byte(data.Items), &items)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	result := &Model{
		Id:          data.Id,
		DashId:      data.DashId,
		Title:       data.Title,
		Description: data.Description.String,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt.Time,
		DeletedAt:   data.DeletedAt.Time,
		Items:       items,
	}

	utils.SendResponse(http.StatusOK, result, w)
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

	data, err := h.repo.GetItemById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	utils.SendResponse(http.StatusOK, data, w)
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.repo.CreateItem(repository.Item{
		DashId:      body.DashId,
		DataQueries: body.DataQueries,
		ItemType:    body.ItemType,
		Position: sql.NullString{
			String: body.Position,
		},
		Title: body.Title,
		Description: sql.NullString{
			String: body.Description,
		},
		Options: body.Options,
	})
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

	updatedId, err := h.repo.UpdateItem(body, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, updatedId, w)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	body, err := utils.GetBody[Model](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.repo.Create(repository.Model{
		Id:     body.Id,
		DashId: body.DashId,
		Title:  body.Title,
		Description: sql.NullString{
			String: body.Description,
			Valid:  true,
		},
	})
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot save data to database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
}
