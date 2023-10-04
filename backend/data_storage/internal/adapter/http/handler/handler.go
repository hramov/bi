package handler

import (
	"github.com/go-chi/chi/v5"
	data_source_repo "github.com/hramov/gvc-bi/backend/datastorage/internal/adapter/postgresrepo/data_source"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type QueryOptions struct {
	Source string `json:"source"`
	Query  string `json:"query"`
	Params []any  `json:"params"`
}

type CheckResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Handler struct {
	repo   *data_source_repo.RepositoryImpl
	logger Logger
}

func New(repo *data_source_repo.RepositoryImpl, logger Logger) *Handler {
	return &Handler{
		repo:   repo,
		logger: logger,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/check", h.checkConnection)
	r.Post("/perform", h.performQuery)
	r.Get("/recall", h.recallDataSources)
}
