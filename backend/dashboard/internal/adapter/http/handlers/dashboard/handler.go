package dashboard_handler

import (
	"context"
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
	Get(ctx context.Context) ([]*dashboard_entity.Dashboard, error)
	GetByDashId(ctx context.Context, id string) (*dashboard_entity.Dashboard, error)
	Create(ctx context.Context, dto dashboards_dto_in.Dashboard) (*int, error)
	Update(ctx context.Context, dto dashboards_dto_in.Dashboard, id int) (*int, error)
	GetItemById(ctx context.Context, id int) (*dashboard_entity.Item, error)
	CreateItem(ctx context.Context, dto dashboards_dto_in.Item) (*int, error)
	UpdateItem(ctx context.Context, dto dashboards_dto_in.Item, id int) (*int, error)
	GetAvailableTypes(ctx context.Context) ([]*dashboard_entity.ItemType, error)
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
