package dashboards

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
	dashboards_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Service interface {
	Get() ([]*dashboard.Dashboard, error)
	GetByDashId(id string) (*dashboard.Dashboard, error)
	Create(dto dashboards_dto_in.Dashboard) (*int, error)
	Update(dto dashboards_dto_in.Dashboard, id int) (*int, error)
	GetItemById(id int) (*dashboard.Item, error)
	CreateItem(dto dashboards_dto_in.Item) (*int, error)
	UpdateItem(dto dashboards_dto_in.Item, id int) (*int, error)
	GetAvailableTypes() ([]*dashboard.ItemType, error)
}

type Handler struct {
	service Service
	logger  Logger
}

func New(service Service, logger dashboard.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/", h.get)
	r.Get("/types", h.getAvailableTypes)
	r.Get("/{id}", h.getByDashId)
	r.Get("/item/{id}", h.getItemById)
	r.Post("/item", h.createItem)
	r.Put("/item/{id}", h.updateItem)
	r.Post("/", h.create)
	r.Put("/", h.update)
	r.Delete("/", h.delete)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Get()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var response []*dashboards_dto_out.Dashboard

	for _, d := range data {
		response = append(response, &dashboards_dto_out.Dashboard{
			Id:          d.Id,
			DashId:      d.DashId,
			Title:       d.Title,
			Description: d.Description,
			CreatedAt:   d.CreatedAt,
			UpdatedAt:   d.UpdatedAt,
			DeletedAt:   d.DeletedAt,
		})
	}

	utils.SendResponse(http.StatusOK, response, w)
}

func (h *Handler) getByDashId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendError(http.StatusInternalServerError, "no id found", w)
		return
	}

	data, err := h.service.GetByDashId(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var items []*dashboards_dto_out.Item

	for _, v := range data.Items {
		items = append(items, &dashboards_dto_out.Item{
			Id:          v.Id,
			DashId:      v.DashId,
			ItemType:    v.ItemType,
			Position:    v.Position,
			Title:       v.Title,
			Description: v.Description,
			DataQueries: v.DataQueries,
			Options:     v.Options,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
		})
	}

	result := &dashboards_dto_out.Dashboard{
		Id:          data.Id,
		DashId:      data.DashId,
		Title:       data.Title,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
		Items:       items,
	}

	utils.SendResponse(http.StatusOK, result, w)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	dash, err := utils.GetBody[dashboards_dto_in.Dashboard](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.service.Create(dash)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot save data to database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
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

	dash, err := utils.GetBody[dashboards_dto_in.Dashboard](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	updatedId, err := h.service.Update(dash, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot save data to database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, updatedId, w)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.GetItemById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	utils.SendResponse(http.StatusOK, data, w)
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	item, err := utils.GetBody[dashboards_dto_in.Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.service.CreateItem(item)
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

	body, err := utils.GetBody[dashboards_dto_in.Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	updatedId, err := h.service.UpdateItem(body, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, updatedId, w)
}

func (h *Handler) getAvailableTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAvailableTypes()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}
	utils.SendResponse(http.StatusOK, data, w)
}
