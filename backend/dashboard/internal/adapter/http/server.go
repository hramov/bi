package http

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/external_api"
	dashboard_handler "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/http/handlers/dashboard"
	datasource_handler "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/http/handlers/datasource"
	user_handler "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/http/handlers/user"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/http/middlewares"
	dashboard_repo "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/dashboard"
	data_source_repo "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/data_source"
	user_repo "github.com/hramov/gvc-bi/backend/dashboard/internal/adapter/postgresrepo/user"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/metrics"
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

// TODO separate composition login
func (s *Server) registerHandlers(r chi.Router) {
	userRepo := user_repo.NewRepository(s.db)
	u := user_handler.New(userRepo, s.logger)
	r.Route("/user", u.Register)

	dashRepo := dashboard_repo.NewRepository(s.db)
	dashService := dashboard.NewService(dashRepo, s.logger)
	d := dashboard_handler.New(dashService, s.logger)
	r.Route("/dashboards", d.Register)

	dsRepo := data_source_repo.NewRepository(s.db)
	externalApi := external_api.New()
	dsService := data_source.NewService(dsRepo, externalApi, s.logger)
	ds := datasource_handler.New(dsService, s.logger)
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

	r.Use(middlewares.ReqId)

	metrics.HandleMetrics(r)

	r.Handle("/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	r.Get("/swagger/*", httpSwagger.WrapHandler)

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
