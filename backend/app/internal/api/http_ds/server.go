package http_ds

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/internal/api/http_ds/handler_ds"
	"github.com/hramov/gvc-bi/backend/internal/repository"
	"log"
	"net/http"
)

type Server struct {
	port int
	db   *sql.DB
}

func New(port int, db *sql.DB) *Server {
	return &Server{port: port, db: db}
}

func (s *Server) registerHandlers(r chi.Router) {
	dsRepo := repository.DatasourceRepository{Db: s.db}
	h := handler_ds.New(dsRepo)
	r.Route("/ds", h.Register)
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
		log.Println("starting datasource server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			log.Println(fmt.Sprintf("cannot start datasource server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	log.Println(fmt.Sprintf("starting graceful shutdown for datasource server"))
	err := s.StopServer()
	if err != nil {
		log.Println(fmt.Sprintf("cannot stop datasource server"))
	}
}

func (s *Server) StopServer() error {
	return nil
}
