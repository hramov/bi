package datasource_handler

import (
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
	GetDrivers() ([]*data_source_entity.Driver, error)
	GetDriverById(id int) (*data_source_entity.Driver, error)
	Get() ([]*data_source_entity.Datasource, error)
	GetById(id int) (*data_source_entity.Datasource, error)
	Create(driver data_source_dto_in.Datasource) (*int, error)
	Update(driver data_source_dto_in.Datasource, id int) (*int, error)
	Delete(id int) (*int, error)
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
