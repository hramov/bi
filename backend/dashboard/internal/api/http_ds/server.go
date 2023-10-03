package http_ds

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/api/http_ds/handler_ds"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/connections"
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
	dsRepo := repository.DatasourceRepository{Db: s.db}
	h := handler_ds.New(dsRepo, s.logger)
	r.Route("/ds", h.Register)
}

func (s *Server) Start(ctx context.Context) error {
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
		s.logger.Info("starting datasource server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			s.logger.Error(fmt.Sprintf("cannot start datasource server: %v", err))
			return
		}
	}()

	dsRepo := repository.DatasourceRepository{Db: s.db}
	ds, err := dsRepo.Get()
	if err != nil {
		return err
	}

	var rc []connections.RawConnection

	for _, v := range ds {
		rc = append(rc, connections.RawConnection{
			SourceId: v.Id,
			DriverId: v.DriverId,
			Dsn:      v.Dsn,
		})
	}

	errs := connections.Connect(rc)

	errStr := ""

	for _, e := range errs {
		errStr += e.Error() + ":"
	}

	if len(errStr) > 0 {
		s.logger.Error(fmt.Sprintf("cannot connect to some data sources: %v", errStr))
	}

	<-ctx.Done()
	s.logger.Info(fmt.Sprintf("starting graceful shutdown for datasource server"))
	err = s.StopServer()
	if err != nil {
		s.logger.Error(fmt.Sprintf("cannot stop datasource server"))
		return err
	}

	return nil
}

func (s *Server) StopServer() error {
	return nil
}
