package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	data_source_entity "github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/entity"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Service interface {
	Get(ctx context.Context) ([]*data_source_entity.Datasource, error)
	GetByCode(ctx context.Context, code string) (*data_source_entity.Datasource, error)
}

type Handler struct {
	service Service
	logger  Logger
}

func New(service Service, logger Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/check", h.checkConnection)
	r.Post("/perform", h.performQuery)
	r.Get("/recall", h.recallDataSources)
}
