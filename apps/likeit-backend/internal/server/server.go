// Package server provides the HTTP server for the "LikeIt" service.
package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const maxShutdownTimeout = 30

type Config struct {
	Port            string `env:"LIKEIT_SERVER_PORT" envDefault:"8080"`
	tokenSigningKey string `env:"LIKEIT_TOKEN_SIGNING_KEY" envDefault:"secret"`
}

type Server struct {
	config Config
	server *http.Server
	router chi.Router
}

func New(config Config) *Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5, "text/html", "text/css", "application/javascript"))

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	return &Server{
		config: config,
		server: server,
		router: router,
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(maxShutdownTimeout)*time.Second,
	)
	defer cancel()

	return s.server.Shutdown(ctx)
}
