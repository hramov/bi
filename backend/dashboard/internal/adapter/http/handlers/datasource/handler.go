package datasource_handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
	data_source_entity "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Service interface {
	GetDrivers(ctx context.Context) ([]*data_source_entity.Driver, error)
	GetDriverById(ctx context.Context, id int) (*data_source_entity.Driver, error)
	Get(ctx context.Context) ([]*data_source_entity.Datasource, error)
	GetById(ctx context.Context, id int) (*data_source_entity.Datasource, error)
	Create(ctx context.Context, driver data_source_dto_in.Datasource) (*int, error)
	Update(ctx context.Context, driver data_source_dto_in.Datasource, id int) (*int, error)
	Delete(ctx context.Context, id int) (*int, error)
}

type Handler struct {
	service Service
	logger  Logger
}

func New(service Service, logger Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Register(r chi.Router) {
	r.Get("/driver", h.getDrivers)
	r.Get("/driver/{id}", h.getDriverById)
	r.Get("/", h.get)
	r.Get("/{id}", h.getById)
	r.Post("/", h.create)
	r.Put("/{id}", h.update)
	r.Delete("/{id}", h.delete)
}
