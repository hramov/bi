package http

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/hramov/gvc-bi/backend/internal/api/handlers/dashboards"
	"github.com/hramov/gvc-bi/backend/internal/api/handlers/users"
	"log"
	"net/http"
)

type Server struct {
	db   *sql.DB
	port int
}

func New(port int, db *sql.DB) *Server {
	return &Server{port: port, db: db}
}

func (s *Server) registerHandlers(r chi.Router) {
	u := users.New(s.db)
	r.Route("/users", u.Register)

	d := dashboards.New(s.db)
	r.Route("/dashboards", d.Register)
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

	r.Route("/api", s.registerHandlers)

	go func() {
		log.Println("starting server")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r); err != nil {
			log.Println(fmt.Sprintf("cannot start server: %v", err))
			return
		}
	}()

	<-ctx.Done()
	log.Println(fmt.Sprintf("starting graceful shutdown for server"))
	err := s.StopServer()
	if err != nil {
		log.Println(fmt.Sprintf("cannot stop server"))
	}
}

func (s *Server) StopServer() error {
	s.db.Close()
	return nil
}
