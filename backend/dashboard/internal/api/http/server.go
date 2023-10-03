package http

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http/handlers/dashboards"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http/handlers/datasource"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http/handlers/users"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/repository"
	"net/http"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

type Server struct {
	port   int
	db     *sql.DB
	logger Logger
}

func New(port int, db *sql.DB, logger Logger) *Server {
	return &Server{port: port, db: db, logger: logger}
}

func (s *Server) registerHandlers(r chi.Router) {
	userRepo := &repository.UsersRepository{Db: s.db}
	u := users.New(userRepo, s.logger)
	r.Route("/users", u.Register)

	dashRepo := dashboard.NewRepository(s.db)
	dashService := dashboard.NewService(dashRepo, s.logger)
	d := dashboards.New(dashService, s.logger)
	r.Route("/dashboards", d.Register)

	dsRepo := &repository.DatasourceRepository{Db: s.db}
	ds := datasource.New(dsRepo)
	r.Route("/datasource", ds.Register)
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

	r.Route("/api", s.registerHandlers)

	go func() {
		s.logger.Info("starting server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			s.logger.Error(fmt.Sprintf("cannot start server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	s.logger.Info(fmt.Sprintf("starting graceful shutdown for server"))
	err := s.StopServer()
	if err != nil {
		s.logger.Error(fmt.Sprintf("cannot stop server"))
	}
}

func (s *Server) StopServer() error {
	s.db.Close()
	return nil
}
