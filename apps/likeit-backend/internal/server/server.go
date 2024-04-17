// Package server provides the HTTP server for the "LikeIt" service.
package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Port            string `env:"PORT" envDefault:"8080"`
	ReadTimeout     int    `env:"LIKEIT_SERVER_READ_TIMEOUT" envDefault:"5"`
	WriteTimeout    int    `env:"LIKEIT_SERVER_WRITE_TIMEOUT" envDefault:"5"`
	ShutdownTimeout int    `env:"LIKEIT_SERVER_SHUTDOWN_TIMEOUT" envDefault:"5"`
}

type Server struct {
	config Config
	server *http.Server
	router chi.Router
}

func New(ctx context.Context, config Config) *Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}
	server.BaseContext = func(net.Listener) context.Context {
		return ctx
	}

	return &Server{
		config: config,
		server: server,
		router: router,
	}
}

func (s *Server) Start() error {
	log.Println("Starting likeit HTTP server... ", s.config.Port)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(s.config.ShutdownTimeout)*time.Second,
	)
	defer cancel()

	return s.server.Shutdown(ctx)
}
