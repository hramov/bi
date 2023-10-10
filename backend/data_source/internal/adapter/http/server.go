package http

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/datasource/internal/adapter/http/handler"
	"github.com/hramov/gvc-bi/backend/datasource/internal/adapter/http/middlewares"
	data_source_repo "github.com/hramov/gvc-bi/backend/datasource/internal/adapter/postgresrepo/data_source"
	"github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source"
	"github.com/hramov/gvc-bi/backend/datasource/internal/domain/data_source/connections"
	"github.com/hramov/gvc-bi/backend/datasource/pkg/metrics"
	httpSwagger "github.com/swaggo/http-swagger"
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
	repo := data_source_repo.NewRepository(s.db)
	service := data_source.NewService(repo, s.logger)
	h := handler.New(service, s.logger)
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

	r.Use(middlewares.ReqId)

	metrics.HandleMetrics(r)

	r.Handle("/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api", s.registerHandlers)

	go func() {
		s.logger.Info("starting datasource server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			s.logger.Error(fmt.Sprintf("cannot start datasource server: %v", err))
			return
		}
	}()

	dsRepo := data_source_repo.NewRepository(s.db)
	ds, err := dsRepo.Get(ctx)
	if err != nil {
		return err
	}

	var rc []connections.RawConnection

	for _, v := range ds {
		rc = append(rc, connections.RawConnection{
			SourceId:   v.Id,
			DriverId:   v.DriverId,
			PluginName: v.PluginName,
			Dsn:        v.Dsn,
		})
	}

	errs := connections.Connect(ctx, rc)

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
