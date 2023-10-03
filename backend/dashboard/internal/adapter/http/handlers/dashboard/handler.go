package dashboard_handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
	dashboard_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Service interface {
	Get() ([]*dashboard_entity.Dashboard, error)
	GetByDashId(id string) (*dashboard_entity.Dashboard, error)
	Create(dto dashboards_dto_in.Dashboard) (*int, error)
	Update(dto dashboards_dto_in.Dashboard, id int) (*int, error)
	GetItemById(id int) (*dashboard_entity.Item, error)
	CreateItem(dto dashboards_dto_in.Item) (*int, error)
	UpdateItem(dto dashboards_dto_in.Item, id int) (*int, error)
	GetAvailableTypes() ([]*dashboard_entity.ItemType, error)
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
