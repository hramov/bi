package user_handler

import (
	"github.com/go-chi/chi/v5"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Service interface {
}

type Handler struct {
	service Service
	logger  Logger
}

type Repository interface {
}

func New(service Service, logger Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/login", h.login)
	r.Post("/check-access", h.checkAccess)
}
