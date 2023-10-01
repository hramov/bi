package http

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/gateway/internal"
	"net/http"
)

type Server struct {
	port   int
	logger internal.Logger
}

func New(port int, logger internal.Logger) *Server {
	return &Server{port: port, logger: logger}
}

func (s *Server) registerHandlers(r chi.Router) {
}

func (s *Server) Start(ctx context.Context) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Route("/data_storage", s.registerHandlers)

	go func() {
		s.logger.Info("starting gateway http server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			s.logger.Error(fmt.Sprintf("cannot start server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	s.logger.Info(fmt.Sprintf("starting graceful shutdown for gateway http server"))
	err := s.StopServer()
	if err != nil {
		s.logger.Error(fmt.Sprintf("cannot stop gateway http server"))
	}
}

func (s *Server) StopServer() error {
	return nil
}
